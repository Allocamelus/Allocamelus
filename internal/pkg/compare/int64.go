package compare

import "crypto/subtle"

func equalInt64(x, y int64) bool {
	if subtle.ConstantTimeEq(int32(x), int32(y)) == 0 ||
		subtle.ConstantTimeEq(int32(x>>32), int32(y>>32)) == 0 {
		return false
	}
	return true
}
