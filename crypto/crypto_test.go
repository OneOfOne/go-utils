package crypto

import (
	"bytes"
	"crypto/des"
	"io/ioutil"
	"testing"
)

var (
	iv  = GenerateAesIV()
	key = GenerateKey(192, []byte(`I'm Jack's Wasted Life`))
	msg = []byte(`I'm Jack's Complete Lack Of Surprise.`)
)

func TestAes(t *testing.T) {
	b := &bytes.Buffer{}
	enc, err := NewAesWriter(b, iv, key)
	if err != nil {
		t.Error(err)
		return
	}
	enc.Write(msg)
	emsg := b.Bytes()
	dec, err := NewAesReader(b, iv, key)
	if err != nil {
		t.Error(err)
		return
	}

	dmsg, _ := ioutil.ReadAll(dec)
	t.Logf("\niv  (%2d): %x\nkey (%2d): %x\nenc (%d): %x\ndec (%d): %s\n", len(iv), iv, len(key), key, len(emsg), emsg, len(dmsg), dmsg)
}

func BenchmarkAes(b *testing.B) {
	buf := &bytes.Buffer{}
	enc, _ := NewAesWriter(buf, iv, key)
	dec, _ := NewAesReader(buf, iv, key)
	for i := 0; i < b.N; i++ {
		enc.Write(msg)
		ioutil.ReadAll(dec)
	}
}

func BenchmarkTripleDes(b *testing.B) {
	buf := &bytes.Buffer{}
	enc, _ := NewWriter(buf, iv[:8], key, des.NewTripleDESCipher)
	dec, _ := NewReader(buf, iv[:8], key, des.NewTripleDESCipher)
	for i := 0; i < b.N; i++ {
		enc.Write(msg)
		ioutil.ReadAll(dec)
	}
}
