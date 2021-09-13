package truthy

func All[T any](vs ...T) bool {
	for _, v := range vs {
		if !Value(v) {
			return false
		}
	}
	return true
}

func Any[T any](vs ...T) bool {
	for _, v := range vs {
		if Value(v) {
			return true
		}
	}
	return false
}

func First[T any](vs ...T) (t T) {
	for _, v := range vs {
		if Value(v) {
			return v
		}
	}
	return t
}

func Filter[T any](vs *[]T) {
	filtered := (*vs)[:0]
	for _, v := range *vs {
		if Value(v) {
			filtered = append(filtered, v)
		}
	}
	*vs = filtered
}

func FilterFunc[T, U any](vs *[]T, keep func(T) U) {
	filtered := (*vs)[:0]
	for _, v := range *vs {
		if Value(keep(v)) {
			filtered = append(filtered, v)
		}
	}
	*vs = filtered
}
