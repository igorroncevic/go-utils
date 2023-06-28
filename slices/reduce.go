package slices

func Reduce[T, M any](slice []T, reducer func(M, T) M, initValue M) M {
	acc := initValue

	for _, elem := range slice {
		acc = reducer(acc, elem)
	}

	return acc
}
