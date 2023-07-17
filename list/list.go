package list

import "github.com/igorroncevic/go-utils/util"

// List is an implementation of single linked-list. No duplicates allowed.
type List[T any] struct {
	Head *Node[T]

	isLessFunc  util.LessFn[T]
	isEqualFunc util.EqualsFn[T]
}

type Node[T any] struct {
	Value      T
	Next, Prev *Node[T]
}

func New[T any](isLessFunc util.LessFn[T], isEqualFunc util.EqualsFn[T]) *List[T] {
	return &List[T]{
		isLessFunc:  isLessFunc,
		isEqualFunc: isEqualFunc,
	}
}

// Push adds a value to the list in an ascending order.
func (l *List[T]) Push(val T) {
	// If there's no head, create it
	if l.Head == nil {
		l.Head = &Node[T]{
			Value: val,
		}

		return
	}

	// head >= val - If element is smaller than head, it should be the new head
	if !l.isLessFunc(l.Head.Value, val) {
		oldHead := l.Head
		l.Head = &Node[T]{
			Value: val,
			Next:  oldHead,
		}

		l.Head.Next.Prev = l.Head

		if oldHead.Next != nil {
			l.Head.Next.Next = oldHead.Next
		}

		return
	}

	l.Head.Each(func(curr *Node[T]) bool {
		// curr < val - we skip over this node
		// Sorted list requires us to find first node where `node.Value > val`
		if l.isLessFunc(curr.Value, val) {
			// Unless, we are at the last node and value is still not smaller -> 'val' should be in the last node
			if curr.Next == nil {
				newNode := Node[T]{
					Value: val,
					Next:  curr.Next,
					Prev:  curr,
				}
				curr.Next = &newNode
				return true
			}

			return false
		}

		// Important: Go back one step, because we've found a first element that is larger than val
		// therefore, previous element should still be smaller and we should put 'val' in between them.
		if curr.Prev != nil {
			curr = curr.Prev
		}

		// 1. curr <-> next
		// 2. insert newNode
		// 	- 2.a. curr.Next ----> newNode,
		// 	- 2.b. newNode.Next -> next
		// 	- 2.c. curr <-------- newNode.Prev,
		// 	- 2.d. newNode <----- next.Prev
		// result: curr <-> newNode <-> next

		// 2.b, 2.c.
		newNode := Node[T]{
			Value: val,
			Next:  curr.Next,
			Prev:  curr,
		}

		// 2.d
		if curr.Next != nil {
			curr.Next.Prev = &newNode
		}

		// 2.a.
		curr.Next = &newNode

		return true
	})
}

// Reverse returns a new reversed list.
func (l *List[T]) Reverse() *List[T] {
	listCopy := l.Copy()

	var (
		tempPrev, tempNext *Node[T]
		curr               = listCopy.Head
	)

	for curr != nil {
		// Temporary store old values
		tempNext = curr.Next
		tempPrev = curr.Prev

		// Reverse current with temporary old values
		curr.Next = tempPrev
		curr.Prev = tempNext

		// Move forward
		tempPrev = curr
		curr = tempNext
	}

	listCopy.Head = tempPrev

	return listCopy
}

// Remove removes the node with value 'val' from the list.
func (l *List[T]) Remove(val T) {
	// If no head, there's nothing to remove
	if l.Head == nil {
		return
	}

	// If there's head, check if we're deleting its value
	if l.Head != nil {
		if l.isEqualFunc(l.Head.Value, val) {
			if l.Head.Next == nil {
				l.Clear()
				return
			}

			next := l.Head.Next
			l.Head = next
			l.Head.Prev = nil

			return
		}
	}

	l.Head.Each(func(curr *Node[T]) bool {
		// If this is not the values we're looking for, move on
		if !l.isEqualFunc(curr.Value, val) {
			return false
		}

		// 1. prev <-> curr <-> next
		// 2. remove newNode
		// 	- 2.a. prev -> next ===== curr.Prev.Next -> curr.Next
		// 	- 2.b. prev <- next ===== curr.Next.Prev -> curr.Prev

		// 2.a.
		if curr.Prev != nil {
			curr.Prev.Next = curr.Next
		}

		// 2.b.
		if curr.Next != nil {
			curr.Next.Prev = curr.Prev
		}

		return true
	})
}

// Contains returns whether the value is contained in the list.
func (l *List[T]) Contains(val T) bool {
	var found bool

	l.Head.Each(func(curr *Node[T]) bool {
		if !l.isEqualFunc(curr.Value, val) {
			return false
		}

		found = true
		return found
	})

	return found
}

// Get returns a node with the specified value.
func (l *List[T]) Get(val T) *Node[T] {
	var (
		found bool
		node  *Node[T]
	)

	l.Head.Each(func(curr *Node[T]) bool {
		if !l.isEqualFunc(curr.Value, val) {
			return false
		}

		node = curr
		found = true
		return found
	})

	return node
}

// Size returns the number of elements in the list.
func (l *List[T]) Size() int {
	var size int

	l.Head.Each(func(n *Node[T]) bool {
		size++
		return false
	})

	return size
}

// Clear clears the list completely.
func (l *List[T]) Clear() {
	l.Head = nil
}

// Copy creates a new list with same values as the original.
func (l *List[T]) Copy() *List[T] {
	copy := &List[T]{
		isLessFunc:  l.isLessFunc,
		isEqualFunc: l.isEqualFunc,
	}

	// Recreate the list from a slice to avoid pointer issues
	sliced := l.ToSlice()

	for _, elem := range sliced {
		copy.Push(elem)
	}

	return copy
}

// ToSlice returns a slice representation of the list.
func (l *List[T]) ToSlice() []T {
	sliced := []T{}

	l.Head.Each(func(n *Node[T]) bool {
		sliced = append(sliced, n.Value)
		return false
	})

	return sliced
}

// EachNode calls 'fn' on every node from this node onward in the list.
func (n *Node[T]) Each(fn func(n *Node[T]) bool) {
	node := n
	for node != nil {
		if shouldStop := fn(node); shouldStop {
			return
		}
		node = node.Next
	}
}
