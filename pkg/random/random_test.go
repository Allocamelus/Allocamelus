package random

import "testing"

const n = 32

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(n)
	}
}

func BenchmarkFastBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastBytes(n)
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(n)
	}
}

func BenchmarkFastInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastInt(n)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(n)
	}
}

func BenchmarkStringBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBase64(n)
	}
}

func BenchmarkStringBase58(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBase58(n)
	}
}

func BenchmarkFastString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastString(n)
	}
}
