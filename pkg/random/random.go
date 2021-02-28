package random

import (
	random "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math"
	"math/big"
	"math/rand"
	"strings"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/mr-tron/base58"
	"k8s.io/klog/v2"
)

func init() {
	var b [16]byte
	_, err := random.Read(b[:])
	if err != nil {
		klog.Fatal("Can't generate secure random bytes:", err)
	}
	rand.Seed(int64(binary.BigEndian.Uint64(b[:])))
}

// Bytes returns a secure random byte array
func Bytes(n int64) []byte {
	b := make([]byte, n)
	_, err := random.Read(b)
	logger.Fatal(err)
	return b
}

// FastBytes returns a fast random byte array
func FastBytes(n int64) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	logger.Fatal(err)
	return b
}

// Int returns a secure random int between 0 and max
// if max <= 0 them max is 64
func Int(max int64) int64 {
	if max <= 0 {
		max = 64
	}
	n, err := random.Int(random.Reader, big.NewInt(max))
	logger.Fatal(err)
	return n.Int64()
}

// FastInt returns a fast random int between 0 and max
// if max <= 0 them max is 64
func FastInt(max int64) int64 {
	if max <= 0 {
		max = 64
	}
	return rand.Int63n(max)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// String returns a secure random string
func String(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, Int(int64(math.MaxInt64)), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = Int(int64(math.MaxInt64)), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// StringBase64 returns a secure random Url Safe string of n*4/3 min length
//
// up to 4x faster than String(n) with only 3 allocations
func StringBase64(n int) string {
	return base64.RawURLEncoding.EncodeToString(Bytes(int64(n)))
}

// StringBase58 returns a secure random base58 string
//
// up to 2.5x faster than String(n) with only 3 allocations
func StringBase58(n int) string {
	return base58.Encode(Bytes(int64(n)))
}

// FastString returns a fast random string
func FastString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
