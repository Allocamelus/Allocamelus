package random_test

import (
	"testing"

	"github.com/allocamelus/allocamelus/pkg/random"
)

const n = 32

func TestBytes(t *testing.T) {
	b := random.Bytes(n)
	if len(b) != n {
		t.Error("Failed random Bytes")
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Bytes(n)
	}
}

func TestFastBytes(t *testing.T) {
	b := random.FastBytes(n)
	if len(b) != n {
		t.Error("Failed random FastBytes")
	}
}

func BenchmarkFastBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.FastBytes(n)
	}
}

func TestInt(_ *testing.T) {
	random.Int(n)
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Int(n)
	}
}

func TestFastInt(_ *testing.T) {
	random.FastInt(n)
}

func BenchmarkFastInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.FastInt(n)
	}
}

func TestString(t *testing.T) {
	s := random.String(n)
	if len(s) != n {
		t.Error("Failed random TestString")
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.String(n)
	}
}

func TestStringBase64(t *testing.T) {
	s := random.StringBase64(n)
	if len(s) < n {
		t.Error("Failed random StringBase64")
	}
}

func BenchmarkStringBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.StringBase64(n)
	}
}

func TestStringBase58(t *testing.T) {
	s := random.StringBase58(n)
	if len(s) < n {
		t.Error("Failed random StringBase58")
	}

}
func BenchmarkStringBase58(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.StringBase58(n)
	}
}

func TestFastString(t *testing.T) {
	s := random.FastString(n)
	if len(s) != n {
		t.Error("Failed random FastString")
	}
}

func BenchmarkFastString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.FastString(n)
	}
}
