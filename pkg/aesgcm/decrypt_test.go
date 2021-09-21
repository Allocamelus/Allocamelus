package aesgcm_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/allocamelus/allocamelus/pkg/aesgcm"
)

var (
	key               = []byte("NoASecureThirtyTwoByteKeyWithPad")
	plainText         = []byte("ThisIsNotEncrypted")
	encryptedBytes    = []byte{190, 173, 58, 208, 163, 213, 16, 65, 90, 50, 109, 172, 204, 109, 79, 142, 247, 86, 109, 188, 196, 247, 9, 166, 201, 97, 236, 30, 220, 79, 206, 246, 164, 94, 21, 74, 57, 146, 241, 28, 216, 30, 41, 187, 79, 124}
	encryptedBytesB64 = base64.RawStdEncoding.EncodeToString(encryptedBytes)
)

func TestDecrypt(t *testing.T) {
	// Decrypt encryptText
	decryptedText, err := aesgcm.Decrypt(key, encryptedBytes)
	if err != nil {
		t.Error("Failed aesgcm Decrypt: ", err)
	}
	// Check decryptedText
	if !bytes.Equal(decryptedText, plainText) {
		t.Error("Failed aesgcm Decrypt decryptedText != plainText")
	}
}

func TestDecryptBase64(t *testing.T) {
	// Decrypt encryptText
	decryptedText, err := aesgcm.DecryptBase64([]byte(key), encryptedBytesB64)
	if err != nil {
		t.Error("Failed aesgcm DecryptBase64: ", err)
	}
	// Check decryptedText
	if !bytes.Equal(decryptedText, plainText) {
		t.Error("Failed aesgcm DecryptBase64 decryptedText != plainText")
	}
}
