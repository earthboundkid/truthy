package bools

// GetFirst returns the first value in vs which is non-zero.
func GetFirst[T comparable](vs ...T) (t T) {
	for _, v := range vs {
		if Comparable(v) {
			return v
		}
	}
	return
}

// GetFirstAny returns the first value in vs which is truthy.
func GetFirstAny[T any](vs ...T) (t T) {
	for _, v := range vs {
		if Any(v) {
			return v
		}
	}
	return
}

// SetFirst sets p to the first non-zero value in defaults
// if it is not already non-zero.
func SetFirst[T comparable](p *T, defaults ...T) {
	if !Pointer(p) {
		*p = GetFirst(defaults...)
	}
}

// SetFirstAny sets p to the first truthy value in defaults
// if it is not already truthy.
func SetFirstAny[T any](p *T, defaults ...T) {
	if !Any(*p) {
		*p = GetFirstAny(defaults...)
	}
}
