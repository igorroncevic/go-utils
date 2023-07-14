package queue

import "fmt"

type Queue[T any] struct {
	elems []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		elems: []T{},
	}
}

// Dequeue puts the item at the back of the queue.
func (q *Queue[T]) Queue(val T) {
	q.elems = append([]T{val}, q.elems...)
}

// Dequeue returns the item at the front of the queue and removes it from the queue.
func (q *Queue[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	element := (q.elems)[q.Len()-1]
	q.elems = (q.elems)[:(q.Len() - 1)]

	return &element, nil
}

// DequeueAll returns all of the items from the queue and empties it.
func (q *Queue[T]) DequeueAll() []T {
	elems := make([]T, q.Len())

	for i := 0; i < len(elems); i++ {
		if elem, err := q.Dequeue(); err != nil {
			elems[i] = *elem
		}
	}

	return elems
}

// Peek returns the item at the front of the queue without removing it.
func (q *Queue[T]) Peek() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	return &q.elems[q.Len()-1], nil
}

func (q *Queue[T]) String() string {
	stringified := ""

	for i := q.Len(); i >= 0; i-- {
		stringified = fmt.Sprintf("%+v %+v", q.elems[i], stringified)
	}

	return stringified
}

func (q *Queue[T]) Len() int {
	return len(q.elems)
}

// IsEmpty: check if queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Clear() {
	q.elems = []T{}
}

// Each calls 'fn' on every item in the queue, starting with the least
// recently pushed element.
func (q *Queue[T]) Each(fn func(t *T)) {
	for i := 0; i < q.Len(); i++ {
		fn(&q.elems[i])
	}
}

// Copy returns a copy of this queue.
func (q *Queue[T]) Copy() *Queue[T] {
	elems := make([]T, q.Len())
	copy(elems, q.elems)

	return &Queue[T]{elems}
}
