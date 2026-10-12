package keyphrase

import (
	"fmt"
	"io"
	"log/slog"
)

// Secret is a caller-owned byte slice whose fmt representation is always
// redacted. Convert it explicitly to []byte or string only at a reviewed
// integration boundary. Assignment or []byte conversion shares its mutable
// backing array; redaction is not encryption and does not cover arbitrary
// binary encoders or wrappers that explicitly extract the bytes.
type Secret []byte

// Clear overwrites the slice's visible length as a best-effort measure, not
// unused capacity or independently copied buffers. Go and the operating system
// may retain compiler, runtime, string, or page copies.
// Aliases of this same slice storage observe the clearing; independent copies
// and immutable strings do not.
func (s Secret) Clear() {
	clear(s)
}

// Format prevents accidental disclosure through logging and debugging verbs.
func (Secret) Format(state fmt.State, _ rune) {
	_, _ = io.WriteString(state, "<redacted>")
}

// LogValue prevents disclosure through the standard structured logger.
func (Secret) LogValue() slog.Value {
	return slog.StringValue("<redacted>")
}

// MarshalText prevents disclosure through standard text and JSON encoders.
func (Secret) MarshalText() ([]byte, error) {
	return []byte("<redacted>"), nil
}
