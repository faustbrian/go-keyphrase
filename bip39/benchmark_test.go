package bip39_test

import (
	"context"
	"testing"

	"github.com/faustbrian/go-keyphrase/v2/bip39"
)

// BenchmarkFromEntropy times English encoding of a public all-zero 32-byte
// fixture, not random generation; lazy list initialization may precede warm use.
func BenchmarkFromEntropy(b *testing.B) {
	entropy := make([]byte, 32)
	for b.Loop() {
		if _, err := bip39.FromEntropy(entropy, bip39.English); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkParseAndNormalize times explicit Japanese parsing of a public
// 12-word fixture, including normalization and checksum validation.
func BenchmarkParseAndNormalize(b *testing.B) {
	phrase := "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あおぞら"
	for b.Loop() {
		if _, err := bip39.ParseLanguage(phrase, bip39.Japanese); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSeedDerivation times normalization, PBKDF2 and output clearing for
// a public 12-word English fixture; mnemonic construction is outside b.Loop.
func BenchmarkSeedDerivation(b *testing.B) {
	mnemonic, err := bip39.FromEntropy(make([]byte, 16), bip39.English)
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	for b.Loop() {
		seed, seedErr := bip39.Seed(context.Background(), mnemonic, "benchmark")
		if seedErr != nil {
			b.Fatal(seedErr)
		}
		clear(seed)
	}
}
