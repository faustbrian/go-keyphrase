package keyphrase_test

import (
	"context"
	"fmt"

	"github.com/faustbrian/go-keyphrase/v2/passphrase"
	"github.com/faustbrian/go-keyphrase/v2/wordlist/eff"
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
