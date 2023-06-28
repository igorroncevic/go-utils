package slices

// Some checks whether any of elements of the slice satisfy the requirement specified in the filtrator func.
//
//	/* returns true, since 1 is positive */
//	onlyOnePositive := []int{-100, -99, 1}
//	isSome := Some(onlyOnePositive, func(elem int){ return elem > 0 }) // returns true
//
//	/* returns false, since none of the elements is positive */
//	allNegative := []int{-10, -84, -36}
//	wontBeAny := Some(allNegative, func(elem int){ return elem > 0 }) // returns false
func Some[T any](slice []T, filtrator func(T) bool) bool {
	for _, elem := range slice {
		if filtrator(elem) {
			return true
		}
	}

	return false
}
