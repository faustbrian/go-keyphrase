package keyphrase_test

import (
	"context"
	"testing"

	keyphrase "github.com/faustbrian/go-keyphrase"
)

func BenchmarkSelectorIndex(b *testing.B) {
	selector := keyphrase.DefaultSelector()
	ctx := context.Background()
	for b.Loop() {
		if _, err := selector.Index(ctx, 7776); err != nil {
			b.Fatal(err)
		}
	}
}
