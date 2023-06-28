package slices

import "fmt"

var (
	ErrNotFound = fmt.Errorf("element not found")
)

// Find attempts to find the specified element in the slice.
// If the element is found, it will be returned, otherwise ErrNotFound is returned.
//
//	words := []string{"somebody", "that i", "used to know"}
//	found, err := Find(words, "somebody") // returns '"somebody", nil'
func Find[T comparable](slice []T, toFind T) (*T, error) {
	var (
		index = -1
		found *T
	)

	for i, elem := range slice {
		if elem == toFind {
			index = i
			break
		}
	}

	if index != -1 {
		found = &slice[index]
	}

	if found == nil {
		return nil, ErrNotFound
	}

	return found, nil
}
