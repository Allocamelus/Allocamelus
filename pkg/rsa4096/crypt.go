package rsa4096

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"errors"
)

// ErrNilPublicKey Nil Public Key
var ErrNilPublicKey = errors.New("rsa4096/crypt: Error Nil Public Key")

// ErrNilPrivateKey Nil Private Key
var ErrNilPrivateKey = errors.New("rsa4096/crypt: Error Nil Private Key")

// ErrNilRsa4096 Nil Rsa4096
var ErrNilRsa4096 = errors.New("rsa4096/crypt: Error Nil Rsa4096")

var label = []byte("OAEP-SHA512")

// Encrypt Data
func (r *Rsa4096) Encrypt(data []byte) ([]byte, error) {
	if r != nil {
		if r.PublicKey == nil {
			if r.PrivateKey == nil {
				return nil, ErrNilPublicKey
			}
			r.PublicKey = &r.PrivateKey.PublicKey
		}
		return rsa.EncryptOAEP(sha512.New(), rand.Reader, r.PublicKey, data, label)
	}
	return nil, ErrNilRsa4096
}

// Decrypt Data
func (r *Rsa4096) Decrypt(data []byte) ([]byte, error) {
	if r != nil {
		if r.PrivateKey != nil {
			return rsa.DecryptOAEP(sha512.New(), rand.Reader, r.PrivateKey, data, label)
		}
		return nil, ErrNilPrivateKey
	}
	return nil, ErrNilRsa4096
}
