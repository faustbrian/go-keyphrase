package password_test

import (
	"context"
	"testing"

	keyphrase "github.com/faustbrian/go-keyphrase/v2"
	"github.com/faustbrian/go-keyphrase/v2/keyphrasetest"
	"github.com/faustbrian/go-keyphrase/v2/password"
)

// BenchmarkGenerateConstrained times full preparation, counting, sampling and
// unranking for 24 symbols/four classes, including clearing each output.
// Policy/generator setup is excluded; this is not a prepared-policy cache test.
func BenchmarkGenerateConstrained(b *testing.B) {
	policy := password.Policy{
		Length:   24,
		Alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%",
		Required: []password.Class{
			{Name: "lower", Characters: "abcdefghijklmnopqrstuvwxyz"},
			{Name: "upper", Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			{Name: "digits", Characters: "0123456789"},
			{Name: "symbols", Characters: "!@#$%"},
		},
	}
	generator := password.DefaultGenerator()
	ctx := context.Background()
	b.ReportAllocs()
	for b.Loop() {
		secret, err := generator.Generate(ctx, policy)
		if err != nil {
			b.Fatal(err)
		}
		clear(secret)
	}
}

// BenchmarkAnalyzeLargePolicy times exact counting for 1024 symbols and an
// unconstrained 62-character alphabet; it does not read randomness.
func BenchmarkAnalyzeLargePolicy(b *testing.B) {
	policy := password.Policy{Length: 1024, Alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}
	for b.Loop() {
		if _, err := password.Analyze(policy); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRejectOversizedPolicy times state-cell admission rejection of a
// 1024-length/10-class policy; it does not measure successful generation.
func BenchmarkRejectOversizedPolicy(b *testing.B) {
	policy := password.Policy{
		Length:   1024,
		Alphabet: "ab",
		Required: make([]password.Class, 10),
	}
	for b.Loop() {
		if _, err := password.Analyze(policy); err == nil {
			b.Fatal("oversized policy was accepted")
		}
	}
}

// BenchmarkGenerateSourceFailure times preparation plus failure at an exhausted
// deterministic source. Generator setup is excluded; no secret is produced.
func BenchmarkGenerateSourceFailure(b *testing.B) {
	selector, err := keyphrase.NewSelector(keyphrasetest.NewSource(nil))
	if err != nil {
		b.Fatal(err)
	}
	generator, err := password.NewGenerator(selector)
	if err != nil {
		b.Fatal(err)
	}
	policy := password.Policy{Length: 24, Alphabet: "abcdefghijklmnopqrstuvwxyz"}
	b.ReportAllocs()
	for b.Loop() {
		if _, generationErr := generator.Generate(context.Background(), policy); generationErr == nil {
			b.Fatal("failing source generated a password")
		}
	}
}
