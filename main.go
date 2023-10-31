package main

import (
	"fmt"

	"github.com/AdeSupriyadi/go-libs/crypto/aes"
)

func main() {
	privateKey := []byte("ZHdy2k72FxQjz9yN")
	plainText := "d3v3l0pm3n7"
	encrypted, err := aes.Encrypt(privateKey, privateKey, []byte(plainText))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Plain Text: %s", plainText)
	fmt.Printf("Encrypted: %s", encrypted)
}
