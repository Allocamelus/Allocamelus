package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
)

// Encrypt Using AES GCM
func Encrypt(key, plainData []byte) []byte {
	c, err := aes.NewCipher(key)
	logger.Fatal(err)

	gcm, err := cipher.NewGCM(c)
	logger.Fatal(err)

	nonce := random.Bytes(int64(gcm.NonceSize()))

	return gcm.Seal(nonce, nonce, plainData, nil)
}

// EncryptBase64 Encrypt return base64
func EncryptBase64(key, plainData []byte) string {
	return base64.RawStdEncoding.EncodeToString(Encrypt(key, plainData))
}
