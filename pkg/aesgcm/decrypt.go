package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/allocamelus/allocamelus/pkg/logger"
)

// Decrypt Using AES GCM
func Decrypt(key, ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	logger.Fatal(err)

	gcm, err := cipher.NewGCM(c)
	logger.Fatal(err)

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("utils/aesgcm: input to small")
	}

	nonce := ciphertext[:nonceSize]
	data := ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, data, nil)
}

// DecryptBase64 Decrypt base64 encoded cipher
func DecryptBase64(key []byte, cipherString string) ([]byte, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(cipherString)
	if err != nil {
		return nil, err
	}
	return Decrypt(key, ciphertext)
}
