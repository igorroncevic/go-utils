package slices

func Map[T any](slice []T, mapper func(T) T) []T {
	res := make([]T, len(slice))

	for i, elem := range slice {
		res[i] = mapper(elem)
	}

	return res
}
