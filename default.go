package truthy

// First returns the first value in vs which is non-zero.
func First[T comparable](vs ...T) (t T) {
	for _, v := range vs {
		if v != *new(T) {
			return v
		}
	}
	return
}

// FirstAny returns the first value in vs which is truthy.
func FirstAny[T any](vs ...T) (t T) {
	for _, v := range vs {
		if ValueAny(v) {
			return v
		}
	}
	return
}

// SetDefault sets p to the first non-zero value in defaults
// if it is not already non-zero.
func SetDefault[T comparable](p *T, defaults ...T) {
	if *p == *new(T) {
		*p = First(defaults...)
	}
}

// SetDefaultAny sets p to the first truthy value in defaults
// if it is not already truthy.
func SetDefaultAny[T any](p *T, defaults ...T) {
	if !ValueAny(*p) {
		*p = FirstAny(defaults...)
	}
}
