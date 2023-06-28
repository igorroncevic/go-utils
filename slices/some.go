package slices

func Some[T any](slice []T, filtrator func(T) bool) bool {
	for _, elem := range slice {
		if filtrator(elem) {
			return true
		}
	}

	return false
}
