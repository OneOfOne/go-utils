package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
)

type CryptoInitilizer func(key []byte) (cipher.Block, error)

// ZeroSlice sets all the bytes in a slice to 0, should be used on keys and IVs.
func ZeroSlice(p []byte) {
	for i := range p {
		p[i] = 0
	}
}

// GenerateIV generates an IV for the given blocksize
func GenerateIV(blockSize int) (p []byte) {
	p = make([]byte, blockSize)
	if _, err := rand.Read(p); err != nil {
		panic(err)
	}
	return
}

// GenerateAesIV generates an IV with AES's BlockSize
func GenerateAesIV() []byte {
	return GenerateIV(aes.BlockSize)
}

// GenerateKey generates a hashed key based on the number of bits
func GenerateKey(bits int, key []byte) []byte {
	if bits%64 != 0 || bits < 128 {
		panic("bits % 64 != 0 || bits < 128")
	}
	n := bits / 8
	p := make([]byte, n+(n%32))

	for i := 0; i < n; i += 32 { // future proof in case we need a key larger than 256 bits
		tmp := append(p, key...)
		h := sha256.Sum256(tmp)
		copy(p[i:], h[:])
		ZeroSlice(tmp)
	}
	return p[:n:n]
}
