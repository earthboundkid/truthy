package pointers

// GetFirst returns the first non-nil pointer it is passed.
func GetFirst[T any](ptrs ...*T) *T {
	for _, ptr := range ptrs {
		if ptr != nil {
			return ptr
		}
	}
	return nil
}
