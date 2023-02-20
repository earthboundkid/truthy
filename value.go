package truthy

import (
	"reflect"
)

// ValueAny returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
//
// Note that the usual Go type system caveats apply around a nil pointer value not being a nil interface value.
//
// In benchmark testing,
// ValueAny is approximately 25 times slower than Value,
// and 50 times slower than native Go comparisons.
func ValueAny[T any](v T) bool {
	switch m := any(v).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&v)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}

// ValueSlice returns true if the length of the slice is greater than 0.
// Note that it does not distinguish nil slices from empty slices.
func ValueSlice[T any, S ~[]T](v S) bool {
	return len(v) > 0
}

// ValueMap returns true if the length of the map is greater than 0.
// Note that it does not distinguish nil maps from empty maps.
func ValueMap[K comparable, V any, M ~map[K]V](v M) bool {
	return len(v) > 0
}

// Value returns the truthy value of comparable types.
// Values are truthy if they are not equal to the zero value for the type.
func Value[T comparable](v T) bool {
	return v != *new(T)
}
