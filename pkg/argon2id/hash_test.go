package argon2id

import (
	"testing"
)

func TestHashSalt(t *testing.T) {
	hash := HashSalt(testPassword, []byte(testPass.Salt), cost)
	if err := hash.Compare(testPassword); err != nil {
		t.Error("Failed to generate correct hash")
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hash(testPassword, cost)
	}
}
