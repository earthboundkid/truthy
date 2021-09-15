package truthy

// And returns true if both a and b are truthy.
func And[T, U any](a T, b U) bool {
	return Value(a) && Value(b)
}

// Or returns false iff neither a nor b is truthy.
func Or[T, U any](a T, b U) bool {
	return Value(a) || Value(b)
}

// Xor returns true iff a or b but not both is truthy.
func Xor[T, U any](a T, b U) bool {
	valA := Value(a)
	valB := Value(b)
	return (valA || valB) && valA != valB
}

// Nor returns true iff neither a nor b is truthy.
func Nor[T, U any](a T, b U) bool {
	return !(Value(a) || Value(b))
}

// Nand returns false iff both a and b are truthy.
func Nand[T, U any](a T, b U) bool {
	return !Value(a) || !Value(b)
}

// Xnor returns true iff both a and b or neither a nor b are truthy.
func Xnor[T, U any](a T, b U) bool {
	valA := Value(a)
	valB := Value(b)
	return (valA && valB) || (!valA && !valB)
}
