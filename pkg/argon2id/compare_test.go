package argon2id

import (
	"testing"
)

var cost = Cost{
	Time:    1,
	Memory:  32 * 1024,
	Threads: 1,
	KeyLen:  32,
}
var (
	testPassword   = "password1"
	testSalt       = "theSaltIs16Bytes"
	testPass       = HashSalt(testPassword, []byte(testSalt), cost)
	testPassString = "$argon2id$v=19$m=32768,t=1,p=1$dGhlU2FsdElzMTZCeXRlcw$lCCvfkidCQNkCb74+V6Jr+eajhfL0yITFeq8Z6hjSDc"
)

func TestCompare(t *testing.T) {
	if testPass.Compare(testPassword) == ErrMismatchedHashAndPassword {
		t.Error("Failed to compare hash")
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testPass.Compare(testPassword)
	}
}
func BenchmarkCompareF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testPass.Compare("notPassword")
	}
}
