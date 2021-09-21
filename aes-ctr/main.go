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
	iv := make([]byte, c.BlockSize())
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	msg := []byte("important message")
	dst := make([]byte, len(msg))

	cipher.NewCTR(c, iv).XORKeyStream(dst, msg)

	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("dst: %s\n", dst)

	cipher.NewCTR(c, iv).XORKeyStream(msg, dst)

	fmt.Printf("dec msg: %s\n", msg)
	fmt.Printf("dec dst: %s\n", dst)
}
