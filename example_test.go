package keyphrase_test

import (
	"context"
	"fmt"

	"github.com/faustbrian/go-keyphrase/passphrase"
	"github.com/faustbrian/go-keyphrase/wordlist/eff"
)

func Example_passphrase() {
	list, err := eff.Large()
	if err != nil {
		panic(err)
	}
	secret, err := passphrase.DefaultGenerator().Generate(context.Background(), passphrase.Policy{
		WordList:  list,
		Words:     6,
		Separator: " ",
	})
	if err != nil {
		panic(err)
	}
	defer clear(secret)

	fmt.Println(len(secret) > 0)
	// Output: true
}
