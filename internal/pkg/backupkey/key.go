package backupkey

import (
	"bytes"
	"strings"

	"github.com/allocamelus/allocamelus/pkg/random"
	"github.com/mr-tron/base58"
)

const (
	splitNum  = 6
	splitRune = '-'
)

// Create backup key
//
//	returns 32 byte key & human base58 string
func Create() ([]byte, string) {
	key := random.Bytes(32)
	return key, encode(key)
}

// Decode backupkey string
//
//	return key []byte & error
func Decode(s string) ([]byte, error) {
	encodedKey := strings.ReplaceAll(s, string(splitRune), "")
	return base58.Decode(encodedKey)
}

// Adds - every splitNumber
func encode(key []byte) string {
	var bytesBuffer bytes.Buffer

	keyStr := base58.Encode(key)
	before := splitNum - 1
	last := len(keyStr) - 1

	for i, rune := range keyStr {
		bytesBuffer.WriteRune(rune)
		if (i+4)%splitNum == before && i != last {
			bytesBuffer.WriteRune(splitRune)
		}
	}

	return bytesBuffer.String()
}
