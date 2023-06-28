package slices

func Every[T any](slice []T, filtrator func(T) bool) bool {
	for _, elem := range slice {
		if !filtrator(elem) {
			return false
		}
	}

	return true
}
