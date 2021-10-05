package truthy

// Cond returns ifVal if cond is truthy,
// otherwise it returns elseVal.
//
// Note that Cond cannot do lazy evaluation of its arguments.
func Cond[T, U any](cond T, ifVal, elseVal U) U {
	if Value(cond) {
		return ifVal
	}
	return elseVal
}
