package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("MatejKarolcik123")
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	p := []byte("this is a very important and secret message")
	enc := gcm.Seal(p[:0], nonce, p, nil)
	fmt.Printf("ciphertext: %s\n", enc)
	fmt.Printf("dst: %s\n", p)
	dec, err := gcm.Open(enc[:0], nonce, enc, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decrypted plaintext: %s\n", dec)
	fmt.Printf("dst: %s\n", enc)
}
