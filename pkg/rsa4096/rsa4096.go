package rsa4096

import (
	"crypto/rand"
	"crypto/rsa"
)

const (
	pemPublic  = "RSA PUBLIC KEY"
	pemPrivate = "RSA PRIVATE KEY"
)

// Rsa4096 Key wrapper
type Rsa4096 struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// GenerateNew Rsa 4096 Keypair
func GenerateNew() (*Rsa4096, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}
	return FromPrivateKey(privateKey), nil
}

// FromPrivateKey Rsa4096 struct from rsa.PrivateKey
func FromPrivateKey(key *rsa.PrivateKey) *Rsa4096 {
	return &Rsa4096{
		PrivateKey: key,
		PublicKey:  &key.PublicKey,
	}
}
