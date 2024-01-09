package bools

// Ternary returns ifVal if cond is true,
// otherwise it returns elseVal.
func Ternary[T any](cond bool, ifVal, elseVal T) T {
	if cond {
		return ifVal
	}
	return elseVal
}
