// Package wordlist validates and exposes immutable word lists with explicit
// provenance metadata. Validation checks entries and checksum consistency, not
// publisher authority or dictionary suitability for a threat model. Raw words
// and order are preserved; normalization checks do not rewrite stored entries.
package wordlist

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

const (
	maxWords    = 1 << 16
	maxWordSize = 1 << 10
)

// ErrorCode identifies a word-list validation failure.
type ErrorCode string

const (
	// CodeEmpty reports a list without entries.
	CodeEmpty ErrorCode = "empty"
	// CodeDuplicate reports an exactly repeated entry.
	CodeDuplicate ErrorCode = "duplicate"
	// CodeNormalizationCollision reports entries with equal NFKD forms.
	CodeNormalizationCollision ErrorCode = "normalization_collision"
	// CodeNormalizationRequired reports an entry that is not already NFKD.
	CodeNormalizationRequired ErrorCode = "normalization_required"
	// CodeMalformed reports an invalid UTF-8, empty, whitespace, or control entry.
	CodeMalformed ErrorCode = "malformed"
	// CodeMetadata reports missing or inconsistent provenance metadata.
	CodeMetadata ErrorCode = "metadata"
	// CodeChecksum reports an embedded-content checksum mismatch.
	CodeChecksum ErrorCode = "checksum"
	// CodePrefixCollision reports entries with a prohibited common prefix.
	CodePrefixCollision ErrorCode = "prefix_collision"
	// CodeOversized reports a list above resource limits.
	CodeOversized ErrorCode = "oversized"
	// CodeInvalidOption reports an invalid validation option.
	CodeInvalidOption ErrorCode = "invalid_option"
)

// Error intentionally excludes words and other potentially sensitive input.
type Error struct {
	// Code identifies the category; package errors use the constants above.
	// Caller-supplied values are rendered verbatim.
	Code ErrorCode
}

// Error formats the validation category without including list entries.
func (e *Error) Error() string {
	return fmt.Sprintf("wordlist: validation failed (%s)", e.Code)
}

// Metadata pins provenance and integrity for a list.
type Metadata struct {
	// ID is the nonempty caller-selected identity, not an authenticated
	// registry entry.
	ID string
	// Language is a nonempty descriptive language label, not validated against
	// a registry.
	Language string
	// Source records a nonempty source location; New does not fetch it.
	Source string
	// Version records a nonempty source version or revision without verifying
	// it.
	Version string
	// License is a nonempty provenance label, not a license compliance verdict.
	License string
	// ExpectedWords must equal the supplied entry count.
	ExpectedWords int
	// SHA256 is the hex SHA-256 of raw entries with a newline after each;
	// matching is case-insensitive and preserves entry order.
	SHA256 string
	// SourceSHA256 is the hex SHA-256 claimed for original source bytes. New
	// checks its shape, not those bytes; trusted provenance is separate.
	SourceSHA256 string
}

// List is an immutable validated word list, safe for concurrent read-only use.
// It must come from New or a validated embedded loader; its zero value does not
// satisfy list-construction guarantees. Methods require a nonnil receiver.
type List struct {
	metadata Metadata
	words    []string
	indices  map[string]int
}

type config struct {
	uniquePrefix int
	requireNFKD  bool
}

// WithNFKD requires every entry to already use BIP-39-compatible NFKD.
func WithNFKD() Option {
	return func(configuration *config) error {
		configuration.requireNFKD = true
		return nil
	}
}

// Option configures list validation.
type Option func(*config) error

// WithUniquePrefix rejects entries sharing the first length Unicode code
// points after NFC display normalization, not raw UTF-8 byte prefixes. Shorter
// entries use their whole normalized form. It checks uniqueness; Index does not
// accept abbreviations.
func WithUniquePrefix(length int) Option {
	return func(configuration *config) error {
		if length <= 0 {
			return &Error{Code: CodeInvalidOption}
		}
		configuration.uniquePrefix = length
		return nil
	}
}

// New validates words and returns an immutable copy of the slice. It admits
// 1..65536 entries of at most 1024 UTF-8 bytes each, rejecting whitespace,
// controls, exact duplicates and NFKD collisions without rewriting raw entries.
// Metadata strings are copied by value; SourceSHA256 is only shape-checked.
// Callers must not modify the input slice during validation.
func New(metadata Metadata, words []string, options ...Option) (*List, error) {
	if len(words) == 0 {
		return nil, &Error{Code: CodeEmpty}
	}
	if exceedsWordLimit(len(words)) {
		return nil, &Error{Code: CodeOversized}
	}
	if !validMetadata(metadata) || metadata.ExpectedWords != len(words) {
		return nil, &Error{Code: CodeMetadata}
	}

	configuration := config{}
	for _, option := range options {
		if option == nil {
			return nil, &Error{Code: CodeInvalidOption}
		}
		if err := option(&configuration); err != nil {
			return nil, err
		}
	}

	copyOfWords := append([]string(nil), words...)
	indices := make(map[string]int, len(copyOfWords))
	normalized := make(map[string]struct{}, len(copyOfWords))
	prefixes := make(map[string]struct{}, len(copyOfWords))
	for index, word := range copyOfWords {
		if !validWord(word) {
			return nil, &Error{Code: CodeMalformed}
		}
		if configuration.requireNFKD && !norm.NFKD.IsNormalString(word) {
			return nil, &Error{Code: CodeNormalizationRequired}
		}
		if _, exists := indices[word]; exists {
			return nil, &Error{Code: CodeDuplicate}
		}

		normalizedWord := norm.NFKD.String(word)
		if _, exists := normalized[normalizedWord]; exists {
			return nil, &Error{Code: CodeNormalizationCollision}
		}
		indices[word] = index
		normalized[normalizedWord] = struct{}{}

		if configuration.uniquePrefix > 0 {
			prefix := runePrefix(normalizedWord, configuration.uniquePrefix)
			if _, exists := prefixes[prefix]; exists {
				return nil, &Error{Code: CodePrefixCollision}
			}
			prefixes[prefix] = struct{}{}
		}
	}

	if !strings.EqualFold(metadata.SHA256, Checksum(copyOfWords)) {
		return nil, &Error{Code: CodeChecksum}
	}

	return &List{metadata: metadata, words: copyOfWords, indices: indices}, nil
}

// Checksum returns the SHA-256 checksum of newline-terminated UTF-8 entries.
func Checksum(words []string) string {
	hash := sha256.New()
	for _, word := range words {
		_, _ = hash.Write([]byte(word))
		_, _ = hash.Write([]byte{'\n'})
	}

	return hex.EncodeToString(hash.Sum(nil))
}

// Len returns the number of entries.
func (l *List) Len() int {
	return len(l.words)
}

// Word returns the entry at index.
func (l *List) Word(index int) (string, bool) {
	if index < 0 || index >= len(l.words) {
		return "", false
	}

	return l.words[index], true
}

// Index returns the exact raw entry index without normalization or case folding.
func (l *List) Index(word string) (int, bool) {
	index, exists := l.indices[word]
	return index, exists
}

// Words returns a copy of all entries.
func (l *List) Words() []string {
	return append([]string(nil), l.words...)
}

// Metadata returns the pinned list metadata.
func (l *List) Metadata() Metadata {
	return l.metadata
}

func validMetadata(metadata Metadata) bool {
	if metadata.ID == "" {
		return false
	}
	if metadata.Language == "" {
		return false
	}
	if metadata.Source == "" {
		return false
	}
	if metadata.Version == "" {
		return false
	}
	if metadata.License == "" {
		return false
	}
	if metadata.ExpectedWords <= 0 {
		return false
	}
	if !validChecksum(metadata.SHA256) {
		return false
	}
	return validChecksum(metadata.SourceSHA256)
}

func validChecksum(value string) bool {
	decoded, err := hex.DecodeString(value)
	if err != nil {
		return false
	}
	return len(decoded) == sha256.Size
}

func exceedsWordLimit(count int) bool {
	return count > maxWords
}

func validWord(word string) bool {
	if word == "" || len(word) > maxWordSize || !utf8.ValidString(word) {
		return false
	}
	for _, character := range word {
		if unicode.IsSpace(character) || unicode.IsControl(character) {
			return false
		}
	}

	return true
}

// runePrefix checks display prefixes in NFC, preserving the raw entry/checksum
// elsewhere. Normalization must respect Unicode starters before slicing runes.
func runePrefix(value string, length int) string {
	runes := []rune(norm.NFC.String(value))
	return string(runes[:min(length, len(runes))])
}
