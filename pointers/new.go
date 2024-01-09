package pointers

// New allocates a new variable of a given value and returns a pointer to it.
func New[T any](value T) *T {
	return &value
}
