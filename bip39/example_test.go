package bip39_test

import (
	"fmt"

	"github.com/faustbrian/go-keyphrase/bip39"
)

func ExampleFromEntropy() {
	mnemonic, err := bip39.FromEntropy(make([]byte, 16), bip39.English)
	if err != nil {
		panic(err)
	}

	fmt.Println(mnemonic.String())
	// Output: abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about
}
