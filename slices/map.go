package slices

// Map mutates every element in the slice using the mapper func and returns the mutated slice.
//
//	numbers := []int{1, 2, 3}
//	mappedNumbers := Map(numbers, func(elem int) { return a + 1 }) // returns [2, 3, 4]
func Map[T any](slice []T, mapper func(T) T) []T {
	res := make([]T, len(slice))

	for i, elem := range slice {
		res[i] = mapper(elem)
	}

	return res
}
