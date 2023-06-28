package slices

// Reduce executes a reducer func on each element of the slice, where the return value is only one element.
//
//	numbers := []int{1, 2, 3}
//	result := Reduce(numbers, func(acc, elem int) { return acc + elem }, 0) // returns 6
func Reduce[T, M any](slice []T, reducer func(M, T) M, initValue M) M {
	acc := initValue

	for _, elem := range slice {
		acc = reducer(acc, elem)
	}

	return acc
}
