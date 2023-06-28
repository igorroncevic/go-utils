package slices

// Every checks whether all elements of the slice satisfy the requirement specified in the filtrator func.
//
//	/* returns true, since none of the elements is negative */
//	allPositive := []int{1, 2, 3}
//	isEvery := Every(allPositive, func(elem int){ return elem > 0 }) // returns true
//
//	/* returns false, since one of the elements is negative */
//	allPositiveButOne := []int{-10, 6, 14}
//	wontBeEvery := Every(allPositiveButOne, func(elem int){ return elem > 0 })
func Every[T any](slice []T, filtrator func(T) bool) bool {
	for _, elem := range slice {
		if !filtrator(elem) {
			return false
		}
	}

	return true
}
