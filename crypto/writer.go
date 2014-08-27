package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

type writer struct {
	enc cipher.Stream
	w   io.Writer
}

// NewReader returns a new Reader with the specific crypto Algroithm
// Note that it uses CFB mode for streams.
func NewWriter(w io.Writer, iv, key []byte, init CryptoInitilizer) (io.Writer, error) {
	c, err := init(key)
	if err != nil {
		return nil, err
	}
	wr := &writer{
		enc: cipher.NewCFBEncrypter(c, iv),
		w:   w,
	}
	return wr, nil
}

func (w *writer) Write(p []byte) (n int, err error) {
	out := make([]byte, len(p))
	w.enc.XORKeyStream(out, p)
	return w.w.Write(out)
}

// NewAesWriter alias for NewWriter(w, iv, key, aes.NewCipher)
func NewAesWriter(w io.Writer, iv, key []byte) (io.Writer, error) {
	return NewWriter(w, iv, key, aes.NewCipher)
}
