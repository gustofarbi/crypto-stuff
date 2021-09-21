package main

import (
	"crypto-stuff/crypto"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	nonce1 := make([]byte, 24)
	if _, err := io.ReadFull(rand.Reader, nonce1); err != nil {
		panic(err)
	}
	nonce2 := make([]byte, 24)
	if _, err := io.ReadFull(rand.Reader, nonce2); err != nil {
		panic(err)
	}
	firstKey := sha256.Sum256([]byte("MatejKarolcik123"))
	secondKey := sha256.Sum256([]byte("foobarblablabla"))
	msg := []byte("this is important")
	fmt.Println("original msg: ", string(msg))
	enc := crypto.Enc(msg, nonce1, firstKey)
	fmt.Println("firstEnc: ", string(enc))
	enc = crypto.Enc(msg, nonce2, secondKey)
	fmt.Println("secondEnc: ", string(enc))
	fmt.Println("firstDec: ", string(crypto.Dec(enc, nonce1, firstKey)))
	fmt.Println("secondDec: ", string(crypto.Dec(enc, nonce2, secondKey)))
}
