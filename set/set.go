package set

import "fmt"

type Set[T comparable] struct {
	values map[T]bool
}

func New[T comparable]() *Set[T] {
	values := make(map[T]bool)
	return &Set[T]{values}
}

func (s *Set[T]) Add(val T) error {
	if _, exists := s.values[val]; exists {
		return fmt.Errorf("value '%v' already exists", val)
	}

	s.values[val] = true

	return nil
}

func (s *Set[T]) Clear() {
	s.values = make(map[T]bool)
}

func (s *Set[T]) Each(fn func(val T)) {
	for key, exists := range s.values {
		if exists {
			fn(key)
		}
	}
}

func (s *Set[T]) Contains(val T) bool {
	return s.values[val]
}

func (s *Set[T]) Remove(val T) error {
	if _, exists := s.values[val]; !exists {
		return fmt.Errorf("value '%v' does not exist", val)
	}

	delete(s.values, val)

	return nil
}

func (s *Set[T]) Size() int {
	return len(s.values)
}
