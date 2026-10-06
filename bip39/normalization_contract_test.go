package bip39_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/faustbrian/go-keyphrase/v2/bip39"
	reference "github.com/tyler-smith/go-bip39"
)

func TestSeedStarterNormalizationUsesLiteralNFKDInput(t *testing.T) {
	t.Parallel()
	const phrase = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
	const normalized = "i\U000113c2\u0316\u0300"
	mnemonic, err := bip39.ParseLanguage(phrase, bip39.English)
	if err != nil {
		t.Fatalf("ParseLanguage() error = %v", err)
	}
	// The reference owns PBKDF2; its expected salt is literal, not calculated
	// through the same normalization dependency as the implementation.
	want := reference.NewSeed(phrase, normalized)
	defer clear(want)
	for _, input := range []string{"i\U000113c2\u0300\u0316", normalized} {
		got, err := bip39.Seed(context.Background(), mnemonic, input)
		if err != nil {
			t.Fatalf("Seed() error = %v", err)
		}
		equal := bytes.Equal(got, want)
		clear(got)
		if !equal {
			t.Fatal("seed differs from literal normalized reference inputs")
		}
	}
}
