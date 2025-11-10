package compare_test

import (
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/compare"
)

var i8list = []int8{
	1,
	2,
	3,
	127,
}
var i16list = []int16{
	1,
	2,
	3,
	127,
	1984,
	32767,
}
var i32list = []int{
	1,
	2,
	3,
	127,
	1984,
	3099479,
	2025613,
	3351760,
	1468285,
	766529,
	1494273,
}
var i64list = []int64{
	1,
	2,
	3,
	127,
	1984,
	3099479,
	2025613,
	3351760,
	1468285,
	766529,
	1494273,
	34359771135,
	72057628431253503,
}

func TestEqualInt(t *testing.T) {
	for _, i := range i8list {
		if !compare.EqualInt(i, i) {
			t.Error("Failed compare int8")
		}
	}
	for _, i := range i16list {
		if !compare.EqualInt(i, i) {
			t.Error("Failed compare int16")
		}
	}
	for _, i := range i32list {
		if !compare.EqualInt(i, i) {
			t.Error("Failed compare int32")
		}
	}
	for _, i := range i64list {
		if !compare.EqualInt(i, i) {
			t.Error("Failed compare int64")
		}
	}
}
