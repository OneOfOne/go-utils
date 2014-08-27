package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

type reader struct {
	dec cipher.Stream
	r   io.Reader
}

// NewReader returns a new Reader with the specific crypto Algroithm
// Note that it uses CFB mode for streams.
func NewReader(r io.Reader, iv, key []byte, init CryptoInitilizer) (io.Reader, error) {
	c, err := init(key)
	if err != nil {
		return nil, err
	}
	rd := &reader{
		dec: cipher.NewCFBDecrypter(c, iv),
		r:   r,
	}
	return rd, nil
}

func (r *reader) Read(p []byte) (n int, err error) {
	in := make([]byte, len(p))
	if n, err = r.r.Read(in); err != nil {
		return
	}
	r.dec.XORKeyStream(p, in)
	return
}

// NewAesReader an alias for NewReader(r, iv, key, aes.NewCipher)
func NewAesReader(r io.Reader, iv, key []byte) (io.Reader, error) {
	return NewReader(r, iv, key, aes.NewCipher)
}
