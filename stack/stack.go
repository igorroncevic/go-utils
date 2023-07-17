package stack

import "fmt"

type Stack[T any] struct {
	elems []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		elems: []T{},
	}
}

// IsEmpty: check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(val T) {
	s.elems = append(s.elems, val)
}

// Remove and return top element of stack.
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	index := s.Size() - 1
	element := (s.elems)[index]
	s.elems = (s.elems)[:index]

	return &element, nil
}

// Peek returns the stack's top element but does not remove it.
func (s *Stack[T]) Peek() (*T, error) {
	if s.Size() == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	return &s.elems[s.Size()-1], nil
}

func (s *Stack[T]) String() string {
	stringified := ""

	for _, elem := range s.elems {
		stringified = fmt.Sprintf("%+v %+v", elem, stringified)
	}

	return stringified
}

func (s *Stack[T]) Size() int {
	return len(s.elems)
}

// Copy returns a copy of this stack.
func (s *Stack[T]) Copy() *Stack[T] {
	elems := make([]T, s.Size())
	copy(elems, s.elems)

	return &Stack[T]{elems}
}
