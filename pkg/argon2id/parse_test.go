package argon2id

import (
	"testing"
)

func TestParse(t *testing.T) {
	hash, _, err := Parse(testPassString)
	if err != nil {
		t.Error("Failed to parse encoded hash")
	}
	if err := hash.Compare(testPassword); err != nil {
		t.Error("Failed to generate correct hash")
	}
}
