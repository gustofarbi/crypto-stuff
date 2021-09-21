package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type Mode string

type Encryptor interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

var (
	CtrMode Mode = "ctr"
	GcmMode Mode = "gcm"
)

func NewEcryptor(c cipher.Block, m Mode) Encryptor {
	if m == CtrMode {
		return newCtr(c)
	} else if m == GcmMode {
		return newGcm(c)
	} else {
		return nil
	}
}

func readRandomBytes(b []byte) {
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Fatalf("cannot read random bytes: %+v\n", err)
	}
}