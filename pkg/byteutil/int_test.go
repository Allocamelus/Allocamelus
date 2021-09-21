package byteutil_test

import (
	"bytes"
	"testing"

	"github.com/allocamelus/allocamelus/pkg/byteutil"
)

var ibList = map[int][]byte{
	1:       {2},
	2:       {4},
	3:       {6},
	1984:    {128, 31},
	3099479: {174, 173, 250, 2},
	2025613: {154, 162, 247, 1},
	3351760: {160, 147, 153, 3},
	1468285: {250, 157, 179, 1},
	766529:  {130, 201, 93},
	1494273: {130, 180, 182, 1},
}

func TestItob(t *testing.T) {
	for i, b := range ibList {
		itob := byteutil.Itob(i)
		if !bytes.Equal(b, itob) {
			t.Error("Failed byteutil Itob")
		}
	}
}

func TestBtoi(t *testing.T) {
	for i, b := range ibList {
		btoi := byteutil.Btoi(b)
		if btoi != i {
			t.Error("Failed byteutil Btoi")
		}
	}
}
