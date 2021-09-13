package truthy

// Cond returns the _if value if cond is truthy,
// otherwise it returns the _else value.
//
// Note that Cond cannot do lazy evaluation of its arguments.
func Cond[T, U any](cond T, _if, _else U) U {
	if Value(cond) {
		return _if
	}
	return _else
}
