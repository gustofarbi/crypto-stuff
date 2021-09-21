package crypto

import (
	"golang.org/x/crypto/salsa20"
)

func Enc(msg, nonce []byte, key [32]byte) []byte {
	res := make([]byte, len(msg))
	salsa20.XORKeyStream(res, msg, nonce, &key)
	return res
}

func Dec(src, nonce []byte, key [32]byte) []byte {
	res := make([]byte, len(src))
	salsa20.XORKeyStream(res, src, nonce, &key)
	return res
}
