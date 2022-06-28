package truthy

// Coalesce returns *p if p is not nil, otherwise v.
func Coalesce[T any](p *T, v T) T {
	if p != nil {
		return *p
	}

	return v
}
