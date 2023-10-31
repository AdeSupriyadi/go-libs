package main

import (
	"fmt"

	"github.com/AdeSupriyadi/libs/crypto/aes"
)

func main() {
	privateKey := []byte("just private key")
	plainText := "plain text"
	encrypted, err := aes.Encrypt(privateKey, privateKey, []byte(plainText))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Plain Text: %s", plainText)
	fmt.Printf("Encrypted: %s", encrypted)
}
