package util

// Itoabs returns the absolute value of x.
func Itoabs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Ftoabs returns the absolute value of x.
func Ftoabs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
