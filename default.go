package truthy

func First[T any](vs ...T) (t T) {
	for _, v := range vs {
		if Value(v) {
			return v
		}
	}
	return
}

func SetDefault[T any](p *T, defaults ...T) {
	if !Value(*p) {
		*p = First(defaults...)
	}
}
