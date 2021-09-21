package aesgcm_test

import (
	"bytes"
	"testing"

	"github.com/allocamelus/allocamelus/pkg/aesgcm"
)

func TestEncrypt(t *testing.T) {
	// Encrypt plainText
	encrypted := aesgcm.Encrypt([]byte(key), []byte(plainText))
	// Decrypt to check
	decryptedText, err := aesgcm.Decrypt([]byte(key), encrypted)
	if err != nil {
		t.Error("Failed aesgcm Encrypt: ", err)
	}
	// Check decryptedText
	if !bytes.Equal(decryptedText, plainText) {
		t.Error("Failed aesgcm Encrypt decryptedText != plainText")
	}
}

func TestEncryptBase64(t *testing.T) {
	// Encrypt plainText
	encrypted := aesgcm.EncryptBase64([]byte(key), []byte(plainText))
	// Decrypt to check
	decryptedText, err := aesgcm.DecryptBase64([]byte(key), encrypted)
	if err != nil {
		t.Error("Failed aesgcm EncryptBase64: ", err)
	}
	// Check decryptedText
	if !bytes.Equal(decryptedText, plainText) {
		t.Error("Failed aesgcm EncryptBase64 decryptedText != plainText")
	}
}
