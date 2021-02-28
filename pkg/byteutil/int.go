package byteutil

import (
	"encoding/binary"
	"runtime/debug"

	"k8s.io/klog/v2"
)

// Itob returns a variable length byte slice of an int
func Itob(i int) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(i))
	return buf[:n]
}

// Btoi returns an int from a byte slice
func Btoi(b []byte) int {
	i, n := binary.Varint(b)
	if n <= 0 {
		klog.Error("Error in Btoi: n of ", n, "\n"+string(debug.Stack()))
	}
	return int(i)
}
