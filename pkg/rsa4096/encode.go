package rsa4096

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

// Encoder for Rsa4096
type Encoder struct {
	*Rsa4096
}

// Encode return Encoder
func (r *Rsa4096) Encode() *Encoder {
	return &Encoder{r}
}

// Base64PrivateKey return base64 encoded Private Key
func (e *Encoder) Base64PrivateKey() string {
	return base64.RawStdEncoding.EncodeToString(e.MarshalPrivateKey())
}

// Base64PublicKey return base64 encoded Public Key
func (e *Encoder) Base64PublicKey() string {
	return base64.RawStdEncoding.EncodeToString(e.MarshalPublicKey())
}

// MarshalPrivateKey return byte encoded Private Key
func (e *Encoder) MarshalPrivateKey() []byte {
	return x509.MarshalPKCS1PrivateKey(e.PrivateKey)
}

// MarshalPublicKey return byte encoded Public Key
func (e *Encoder) MarshalPublicKey() []byte {
	return x509.MarshalPKCS1PublicKey(e.PublicKey)
}

// PemPrivateKey return Pem encoded Private Key
func (e *Encoder) PemPrivateKey() []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  pemPrivate,
			Bytes: e.MarshalPrivateKey(),
		},
	)
}

// PemPublicKey return Pem encoded Public Key
func (e *Encoder) PemPublicKey() []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  pemPublic,
			Bytes: e.MarshalPublicKey(),
		},
	)
}
