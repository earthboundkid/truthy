package truthy

// Cond returns ifVal if cond is true,
// otherwise it returns elseVal.
func Cond[T any](cond bool, ifVal, elseVal T) T {
	if cond {
		return ifVal
	}
	return elseVal
}
