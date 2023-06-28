package slices

import "fmt"

var (
	ErrNotFound = fmt.Errorf("element not found")
)

func Find[T comparable](slice []T, toFind T) (*T, error) {
	var found *T

	for _, elem := range slice {
		if elem == toFind {
			found = &elem
			break
		}
	}

	if found == nil {
		return nil, ErrNotFound
	}

	return found, nil
}
