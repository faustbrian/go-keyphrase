package keyphrase

import (
	"fmt"
	"io"
)

// ErrorCode identifies a generation failure without exposing secret material.
type ErrorCode string

const (
	// CodeInvalidBound reports a nonpositive selection bound.
	CodeInvalidBound ErrorCode = "invalid_bound"
	// CodeInvalidSource reports a nil randomness source.
	CodeInvalidSource ErrorCode = "invalid_source"
	// CodeInvalidOption reports an invalid selector option.
	CodeInvalidOption ErrorCode = "invalid_option"
	// CodeOversized reports a request above a resource limit.
	CodeOversized ErrorCode = "oversized"
	// CodeSource reports a randomness-source failure.
	CodeSource ErrorCode = "source_failure"
	// CodeShortRead reports a source that made invalid progress.
	CodeShortRead ErrorCode = "short_read"
	// CodeAttemptsExceeded reports too many rejected samples.
	CodeAttemptsExceeded ErrorCode = "attempts_exceeded"
	// CodeCanceled reports context cancellation.
	CodeCanceled ErrorCode = "canceled"
)

// Error is a typed, secret-safe generation error.
//
// Error deliberately omits the wrapped error text. Callers that need the
// underlying cause may inspect it with errors.Is or errors.As, but should not
// log arbitrary source errors without reviewing their disclosure behavior.
type Error struct {
	// Code identifies the category; package errors use the constants above.
	// Caller-supplied codes are rendered verbatim.
	Code ErrorCode
	// Cause retains the underlying error for errors.Is/As. Default formatting
	// omits its text; logging an extracted cause may disclose source data.
	Cause error
}

// Error formats only the operation and Code, never the wrapped Cause.
func (e *Error) Error() string {
	return fmt.Sprintf("keyphrase: generation failed (%s)", e.Code)
}

// Unwrap returns Cause without sanitizing it for logging or encoding.
func (e *Error) Unwrap() error {
	return e.Cause
}

// Format prevents wrapped source diagnostics from appearing in debug output.
func (e *Error) Format(state fmt.State, _ rune) {
	_, _ = io.WriteString(state, e.Error())
}

// MarshalText omits wrapped source diagnostics from encoded output.
func (e *Error) MarshalText() ([]byte, error) {
	return []byte(e.Error()), nil
}
