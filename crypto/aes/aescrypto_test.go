package aes

import (
	"testing"
)

func TestAesCrypto_Encrypt(t *testing.T) {
	key := []byte("secretKeyhere123")
	iv := key
	plaintext := "123456"

	t.Logf("Plain text : %s\n", plaintext)
	ciphertext, err := Encrypt(key, iv, []byte(plaintext))
	if err != nil {
		t.Error(err)
	}
	//t.Logf("Encrypted  : %0x\n", ciphertext)
	t.Logf("Encrypted  : %s\n", ciphertext)

}

func TestAesCrypto_Decrypt(t *testing.T) {
	key := []byte("secretKeyhere123")
	iv := key
	ciphertext := "375586C352377F0338F1F6368F492051"

	t.Logf("Encrypted text : %s\n", ciphertext)
	plaintextb, err := Decrypt(key, iv, ciphertext)
	if err != nil {
		t.Error(err)
	}
	plaintext := string(plaintextb)
	t.Logf("Encrypted  : %s\n", plaintext)
}
