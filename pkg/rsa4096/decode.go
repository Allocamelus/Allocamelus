package rsa4096

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// ErrInvalidPem Invalid Pem Block Type
var ErrInvalidPem = errors.New("rsa4096/decode: Error Invalid Pem Block Type")

// Decoder for Rsa4096
type Decoder struct{}

// Decode return Decoder
func Decode() *Decoder {
	return &Decoder{}
}

// Base64PrivateKey return base64 Decoded Private Key
func (e *Decoder) Base64PrivateKey(s string) (*Rsa4096, error) {
	b, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return e.UnmarshalPrivateKey(b)
}

// Base64PublicKey return base64 Decoded Public Key
func (e *Decoder) Base64PublicKey(s string) (*Rsa4096, error) {
	b, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return e.UnMarshalPublicKey(b)
}

// UnmarshalPrivateKey return byte Decoded Private Key
func (e *Decoder) UnmarshalPrivateKey(b []byte) (*Rsa4096, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return FromPrivateKey(privateKey), nil
}

// UnMarshalPublicKey return byte Decoded Public Key
func (e *Decoder) UnMarshalPublicKey(b []byte) (*Rsa4096, error) {
	rsaKey, err := x509.ParsePKCS1PublicKey(b)
	if err != nil {
		return nil, err
	}
	return &Rsa4096{PublicKey: rsaKey}, nil
}

// PemPrivateKey return Pem Decoded Private Key
func (e *Decoder) PemPrivateKey(b []byte) (*Rsa4096, error) {
	block, _ := pem.Decode(b)

	if block.Type != pemPrivate {
		return nil, ErrInvalidPem
	}

	return e.UnmarshalPrivateKey(b)
}

// PemPublicKey return Pem Decoded Public Key
func (e *Decoder) PemPublicKey(b []byte) (*Rsa4096, error) {
	block, _ := pem.Decode(b)

	if block.Type != pemPublic {
		return nil, ErrInvalidPem
	}

	return e.UnMarshalPublicKey(b)
}
