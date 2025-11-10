package compare

func EqualInt[T int | int8 | int16 | int32 | int64](x, y T) bool {
	return equalInt64(int64(x), int64(y))
}
