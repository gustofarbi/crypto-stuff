package crypto

import (
	"crypto/cipher"
)

type ctr struct {
	iv      []byte
	backend cipher.Block
}

func newCtr(c cipher.Block) *ctr {
	return &ctr{
		iv:      make([]byte, c.BlockSize()),
		backend: c,
	}
}

func (c *ctr) Encrypt(src []byte) []byte {
	readRandomBytes(c.iv)
	return append(c.doXor(src, c.iv), c.iv...)
}

func (c *ctr) Decrypt(src []byte) []byte {
	iv := src[len(src)-c.backend.BlockSize():]
	src = src[:len(src)-c.backend.BlockSize()]
	return c.doXor(src, iv)
}

func (c *ctr) doXor(src []byte, iv []byte) []byte {
	dst := make([]byte, len(src))
	cipher.NewCTR(c.backend, iv).XORKeyStream(dst, src)
	return dst
}
