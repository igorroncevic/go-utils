package slices

func Filter[T any](slice []T, filtrator func(T) bool) []T {
	var res []T

	for _, elem := range slice {
		if filtrator(elem) {
			res = append(res, elem)
		}
	}

	return res
}
