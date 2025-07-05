package pointer

func Deal[T any](p T) *T {
	return &p
}

func DealWithConvert[E any, T any](target *E, fn func(wrapTarget E) *T) *T {
	if target == nil {
		return nil
	}
	return fn(*target)
}
