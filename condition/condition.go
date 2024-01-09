package condition

// Evaluate returns ifVal if cond is true,
// otherwise it returns elseVal.
func Evaluate[T any](cond bool, ifVal, elseVal T) T {
	if cond {
		return ifVal
	}
	return elseVal
}
