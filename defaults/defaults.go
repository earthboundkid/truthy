package defaults

import "github.com/carlmjohnson/truthy/is"

// GetFirst returns the first value in vs which is non-zero.
func GetFirst[T comparable](vs ...T) (t T) {
	for _, v := range vs {
		if is.Truthy(v) {
			return v
		}
	}
	return
}

// GetFirstAny returns the first value in vs which is truthy.
func GetFirstAny[T any](vs ...T) (t T) {
	for _, v := range vs {
		if is.TruthyAny(v) {
			return v
		}
	}
	return
}

// SetFirst sets p to the first non-zero value in defaults
// if it is not already non-zero.
func SetFirst[T comparable](p *T, defaults ...T) {
	if !is.TruthyPointer(p) {
		*p = GetFirst(defaults...)
	}
}

// SetFirstAny sets p to the first truthy value in defaults
// if it is not already truthy.
func SetFirstAny[T any](p *T, defaults ...T) {
	if !is.TruthyAny(*p) {
		*p = GetFirstAny(defaults...)
	}
}
