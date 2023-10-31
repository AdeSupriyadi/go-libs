package main

import (
	"fmt"

	"github.com/AdeSupriyadi/go-libs/crypto/aes"
)

func main() {
	privateKey := []byte("just private key")
	plainText := "palin text"
	encrypted, err := aes.Encrypt(privateKey, privateKey, []byte(plainText))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Plain Text: %s", plainText)
	fmt.Printf("Encrypted: %s", encrypted)
}
