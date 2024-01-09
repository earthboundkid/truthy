package bools

import (
	"reflect"
)

// Comparable returns the truthy value of comparable types.
// Values are truthy if they are not equal to the zero value for the type.
// Use the Pointer method to evaluate underlying pointer values, otherwise
// this will return true if the pointer is not nil.
func Comparable[T comparable](v T) bool {
	return v != *new(T)
}

// Any returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
//
// Note that the usual Go type system caveats apply around a nil pointer value not being a nil interface value.
//
// In benchmark testing,
// Any is approximately 25 times slower than Comparable,
// and 50 times slower than native Go comparisons.
func Any[T any](v T) bool {
	switch m := any(v).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&v)
}

// Map returns true if the length of the map is greater than 0.
// Note that it does not distinguish nil maps from empty maps.
func Map[K comparable, V any, M ~map[K]V](v M) bool {
	return len(v) > 0
}

// Pointer returns the truthy value of dereferenced pointers for comparable types.
// Values are truthy if they are not equal to the zero value for the dereferenced type.
// Note that it will evaluate to true for double pointers.
func Pointer[T comparable](v *T) bool {
	if v == nil {
		return false
	}
	return Comparable(*v)
}

// Slice returns true if the length of the slice is greater than 0.
// Note that it does not distinguish nil slices from empty slices.
func Slice[T any, S ~[]T](v S) bool {
	return len(v) > 0
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}
