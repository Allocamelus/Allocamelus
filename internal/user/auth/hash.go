package auth

import (
	"encoding/base64"

	"golang.org/x/crypto/blake2b"
)

func HashKey(key string) (keyHash []byte, err error) {
	keyB, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return
	}
	keyBHash := blake2b.Sum512(keyB)
	keyHash = keyBHash[:]
	return
}

func HashKeyB64(key string) (keyHash string, err error) {
	keyBHash, err := HashKey(key)
	if err != nil {
		return
	}
	keyHash = base64.RawStdEncoding.EncodeToString(keyBHash[:])
	return
}
