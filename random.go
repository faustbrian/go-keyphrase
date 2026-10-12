package keyphrase

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"math/bits"
)

const defaultMaxAttempts = 128
const maxSampleBits = 1 << 20

// Source supplies cryptographic random bytes and participates in cancellation.
// Implementations must fill as much of destination as possible and must return
// promptly after ctx is canceled. Randomness quality is a trusted caller
// assumption, not verified by this interface or by statistical test helpers.
type Source interface {
	// ReadContext reads into caller-provided space. Selector accepts positive
	// partial progress with nil error, but rejects any nonnil error, even with
	// a full count. Zero or invalid counts fail without a retry loop.
	ReadContext(ctx context.Context, destination []byte) (int, error)
}

type cryptoSource struct{}

func (cryptoSource) ReadContext(ctx context.Context, destination []byte) (int, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	return rand.Read(destination)
}

// Selector performs bounded rejection sampling from a Source. Uniform and
// independent source bytes are required for uniform results. Selector retains
// source without wrapping it in a lock; concurrency safety depends on source.
// The zero value is invalid. Contexts must be nonnil when an operation checks
// them; the library cannot interrupt a Source call that ignores cancellation.
type Selector struct {
	source      Source
	maxAttempts int
}

// Option configures a Selector.
type Option func(*Selector) error

// WithMaxAttempts bounds candidate draws per Index or BigInt operation. It
// requires a positive count; it does not bound partial-read calls or time spent
// inside the source. The default is 128 candidate draws.
func WithMaxAttempts(maxAttempts int) Option {
	return func(selector *Selector) error {
		if maxAttempts <= 0 {
			return &Error{Code: CodeInvalidOption}
		}

		selector.maxAttempts = maxAttempts
		return nil
	}
}

// NewSelector constructs a selector retaining source. Options run in order,
// are not retained, and may return their own errors. Nil source/option values
// are rejected; custom source quality and sharing remain caller-owned.
func NewSelector(source Source, options ...Option) (*Selector, error) {
	if source == nil {
		return nil, &Error{Code: CodeInvalidSource}
	}

	selector := &Selector{source: source, maxAttempts: defaultMaxAttempts}
	for _, option := range options {
		if option == nil {
			return nil, &Error{Code: CodeInvalidOption}
		}
		if err := option(selector); err != nil {
			return nil, err
		}
	}

	return selector, nil
}

// DefaultSelector returns a selector backed by crypto/rand. It checks context
// before random reads, but cannot interrupt an in-progress crypto/rand read.
func DefaultSelector() *Selector {
	return &Selector{source: cryptoSource{}, maxAttempts: defaultMaxAttempts}
}

// Index rejection-samples a value in [0, upper), uniformly if source bytes are
// uniform and independent. upper must be nonzero. A bound of one returns zero
// without a read after checking ctx. Cancellation is checked before reads,
// not again after a successful final read.
func (s *Selector) Index(ctx context.Context, upper uint64) (uint64, error) {
	if s == nil || s.source == nil {
		return 0, &Error{Code: CodeInvalidSource}
	}
	if upper == 0 {
		return 0, &Error{Code: CodeInvalidBound}
	}
	if err := ctx.Err(); err != nil {
		return 0, &Error{Code: CodeCanceled, Cause: err}
	}
	if upper == 1 {
		return 0, nil
	}

	bitCount := bits.Len64(upper - 1)
	byteCount := (bitCount + 7) / 8
	mask := uint64(1)<<bitCount - 1
	if bitCount == 64 {
		mask = ^uint64(0)
	}

	var buffer [8]byte
	for range s.maxAttempts {
		if err := s.readFull(ctx, buffer[8-byteCount:]); err != nil {
			return 0, err
		}

		// Masking chooses a power-of-two domain; rejecting its excess tail avoids
		// modulo bias without assuming upper divides the byte domain.
		candidate := binary.BigEndian.Uint64(buffer[:]) & mask
		if candidate < upper {
			return candidate, nil
		}
	}

	return 0, &Error{Code: CodeAttemptsExceeded}
}

// BigInt returns a uniformly distributed arbitrary-precision value in
// [0, upper), assuming uniform independent source bytes. upper must be positive
// and must not be modified during the call. At most 2^20 bits are sampled; the
// returned integer never aliases upper. A bound of one needs no source read.
func (s *Selector) BigInt(ctx context.Context, upper *big.Int) (*big.Int, error) {
	if s == nil || s.source == nil {
		return nil, &Error{Code: CodeInvalidSource}
	}
	if upper == nil || upper.Sign() <= 0 {
		return nil, &Error{Code: CodeInvalidBound}
	}
	if err := ctx.Err(); err != nil {
		return nil, &Error{Code: CodeCanceled, Cause: err}
	}
	if upper.Cmp(big.NewInt(1)) == 0 {
		return new(big.Int), nil
	}

	bitCount := new(big.Int).Sub(new(big.Int).Set(upper), big.NewInt(1)).BitLen()
	if bitCount > maxSampleBits {
		return nil, &Error{Code: CodeOversized}
	}
	byteCount := (bitCount + 7) / 8
	excessBits := byteCount*8 - bitCount
	buffer := make([]byte, byteCount)

	for range s.maxAttempts {
		clear(buffer)
		if err := s.readFull(ctx, buffer); err != nil {
			clear(buffer)
			return nil, err
		}
		buffer[0] &= byte(0xff >> excessBits)

		candidate := new(big.Int).SetBytes(buffer)
		if candidate.Cmp(upper) < 0 {
			clear(buffer)
			return candidate, nil
		}
	}

	clear(buffer)
	return nil, &Error{Code: CodeAttemptsExceeded}
}

// Fill writes cryptographic random bytes into destination. On failure it
// clears destination before returning an error. It allows at most 128 KiB;
// a valid selector with an empty destination succeeds without reading or
// checking ctx. Callers must not observe or mutate destination during a call;
// it can contain partial bytes until an error clears it.
func (s *Selector) Fill(ctx context.Context, destination []byte) error {
	if s == nil || s.source == nil {
		clear(destination)
		return &Error{Code: CodeInvalidSource}
	}
	if len(destination) > maxSampleBits/8 {
		clear(destination)
		return &Error{Code: CodeOversized}
	}
	if err := s.readFull(ctx, destination); err != nil {
		clear(destination)
		return err
	}

	return nil
}

func (s *Selector) readFull(ctx context.Context, destination []byte) error {
	offset := 0
	for offset < len(destination) {
		if err := ctx.Err(); err != nil {
			return &Error{Code: CodeCanceled, Cause: err}
		}

		count, err := s.source.ReadContext(ctx, destination[offset:])
		if err != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				return &Error{Code: CodeCanceled, Cause: ctxErr}
			}
			return &Error{Code: CodeSource, Cause: err}
		}
		if count < 1 || count > len(destination)-offset {
			return &Error{Code: CodeShortRead, Cause: io.ErrNoProgress}
		}
		offset += count
	}

	return nil
}
