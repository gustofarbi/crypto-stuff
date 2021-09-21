package crypto

import (
	"crypto/cipher"
	"log"
)

type gcm struct {
	nonce   []byte
	backend cipher.AEAD
}

func newGcm(c cipher.Block) *gcm {
	g, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatalf("cannot create gcm: %+v\n", err)
	}
	return &gcm{
		nonce:   make([]byte, g.NonceSize()),
		backend: g,
	}
}

func (g *gcm) Encrypt(src []byte) []byte {
	readRandomBytes(g.nonce)
	return append(g.backend.Seal(src[:0], g.nonce, src, nil), g.nonce...)
}

func (g *gcm) Decrypt(src []byte) []byte {
	nonce := src[len(src)-g.backend.NonceSize():]
	src = src[:len(src)-g.backend.NonceSize()]

	plain, err := g.backend.Open(src[:0], nonce, src, nil)
	if err != nil {
		log.Fatalf("cannot decrypt ciphertext: %+v\n", err)
	}
	return plain
}
