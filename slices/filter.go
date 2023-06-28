package slices

// Filter filters out elements that do not satisfy the requirement specified in the filtrator func and returns the filtrated slice.
//
//	allPositive := []int{1, 2, 3}
//	result := Filter(allPositive, func(elem int){ return elem > 0 }) // returns [1, 2, 3]
//
//	allPositiveButOne := []int{-10, 6, 14}
//	result := Filter(allPositiveButOne, func(elem int){ return elem > 0 }) // returns [6, 14]
func Filter[T any](slice []T, filtrator func(T) bool) []T {
	var res []T

	for _, elem := range slice {
		if filtrator(elem) {
			res = append(res, elem)
		}
	}

	return res
}
