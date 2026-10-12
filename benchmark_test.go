package keyphrase_test

import (
	"context"
	"testing"

	keyphrase "github.com/faustbrian/go-keyphrase/v2"
)

// BenchmarkSelectorIndex times default crypto/rand rejection sampling below 7776.
// Selector/context setup is outside b.Loop; no entropy-quality claim is made.
func BenchmarkSelectorIndex(b *testing.B) {
	selector := keyphrase.DefaultSelector()
	ctx := context.Background()
	for b.Loop() {
		if _, err := selector.Index(ctx, 7776); err != nil {
			b.Fatal(err)
		}
	}
}
