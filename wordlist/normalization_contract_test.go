package wordlist_test

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/faustbrian/go-keyphrase/v2/wordlist"
)

func TestUniquePrefixRespectsUnicodeStarter(t *testing.T) {
	t.Parallel()
	words := []string{"i\U000113c2\u0300\u0316", "ixfixture"}
	metadata := metadataFor(words)
	for _, length := range []int{0, 1, 2} {
		t.Run(strconv.Itoa(length), func(t *testing.T) {
			var options []wordlist.Option
			if length != 0 {
				options = append(options, wordlist.WithUniquePrefix(length))
			}
			list, err := wordlist.New(metadata, words, options...)
			if length == 1 {
				var failure *wordlist.Error
				if list != nil || !errors.As(err, &failure) || failure.Code != wordlist.CodePrefixCollision {
					t.Fatalf("New() = %v, %v; want nil list and prefix collision", list != nil, err)
				}
				for _, word := range words {
					if strings.Contains(err.Error(), word) {
						t.Fatal("prefix error disclosed list input")
					}
				}
				return
			}
			if err != nil || list == nil {
				t.Fatalf("New() error = %v", err)
			}
			if !reflect.DeepEqual(list.Words(), words) || list.Metadata() != metadata {
				t.Fatal("successful validation changed raw words or metadata")
			}
			for index, word := range words {
				got, ok := list.Word(index)
				if !ok || got != word {
					t.Fatal("Word() changed a raw entry")
				}
				gotIndex, found := list.Index(word)
				if !found || gotIndex != index {
					t.Fatal("Index() lost a raw entry")
				}
			}
		})
	}
}

func TestStarterNormalizationRequirementPreservesRawIdentity(t *testing.T) {
	t.Parallel()
	raw := "i\U000113c2\u0300\u0316"
	normalized := "i\U000113c2\u0316\u0300"
	for _, test := range []struct {
		word     string
		accepted bool
	}{{raw, false}, {normalized, true}} {
		words := []string{test.word}
		metadata := metadataFor(words)
		list, err := wordlist.New(metadata, words, wordlist.WithNFKD())
		if !test.accepted {
			var failure *wordlist.Error
			if list != nil || !errors.As(err, &failure) || failure.Code != wordlist.CodeNormalizationRequired {
				t.Fatalf("New() error = %v; want normalization required", err)
			}
			continue
		}
		if err != nil || list == nil {
			t.Fatalf("New() error = %v", err)
		}
		if !reflect.DeepEqual(list.Words(), words) || list.Metadata() != metadata {
			t.Fatal("NFKD admission changed raw words or metadata")
		}
		got, ok := list.Word(0)
		index, found := list.Index(normalized)
		if !ok || got != normalized || !found || index != 0 {
			t.Fatal("normalized raw entry lost identity")
		}
	}
}
