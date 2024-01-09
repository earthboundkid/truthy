package bools

// Int returns 1 if v is true, otherwise 0.
// Inspired by https://github.com/golang/go/issues/64825
func Int[T bool](v T) int {
	if v {
		return 1
	}
	return 0
}
