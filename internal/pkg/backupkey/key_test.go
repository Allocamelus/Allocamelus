package backupkey_test

import (
	"crypto/subtle"
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/backupkey"
)

func TestCreate(t *testing.T) {
	key, encodedKey := backupkey.Create()
	if len(key) != 32 {
		t.Error("Failed to generate 32 byte key")
	}
	if len(encodedKey) != 51 {
		t.Error("Failed to encode key")
	}
}
func TestDecode(t *testing.T) {
	key, encodedKey := backupkey.Create()
	decodedKey, err := backupkey.Decode(encodedKey)
	if err != nil {
		t.Error("Failed to decode key", err)
	}
	if subtle.ConstantTimeCompare(key, decodedKey) != 1 {
		t.Error("Failed to decode matching key")
	}
}
