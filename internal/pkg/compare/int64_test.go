package compare_test

import (
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/compare"
)

var i64list = []int64{
	1,
	2,
	3,
	1984,
	3099479,
	2025613,
	3351760,
	1468285,
	766529,
	1494273,
}

func TestEqualInt64(t *testing.T) {
	for _, i := range i64list {
		if !compare.EqualInt64(i, i) {
			t.Error("Failed aesgcm Decrypt decryptedText != plainText")
		}
	}
}
