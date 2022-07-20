package truthy

// Coalesce returns *p if p is not nil, otherwise v.
func Coalesce[T any](p *T, v T) T {
	if p != nil {
		return *p
	}

	return v
}

// Deref returns *p if p is not nil, otherwise the zero value of T.
func Deref[T any](p *T) T {
	return Coalesce(p, *new(T))
}
