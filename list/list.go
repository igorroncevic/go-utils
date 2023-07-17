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

func (l *List[T]) Push(val T) {
	// If no head, create it
	if l.Head == nil {
		l.Head = &Node[T]{
			Value: val,
		}

		return
	}

	// NOP
	if l.isEqualFunc(l.Head.Value, val) {
		return
	}

	// If element is smaller than head, it should be the new head
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
		// If there is the exact same element, just update the value
		if l.isEqualFunc(curr.Value, val) {
			curr.Value = val
			return true
		}

		// Sorted list requires us to find first node where `node.Value > val`
		if l.isLessFunc(curr.Value, val) {
			// We are at the last node, value is still not smaller -> val be in the last node
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
		// therefore, previous element should still be smaller and we should put val in between them.
		if curr.Prev != nil {
			curr = curr.Prev
		}

		// 1. curr -> next
		// 2. insert newNode
		// 	- 2.a. curr.Next ----> newNode,
		// 	- 2.b. newNode.Next -> next
		// 	- 2.c. curr <-------- newNode.Prev,
		// 	- 2.d. newNode <----- next.Prev

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

// Remove removes the node 'n' from the list.
func (l *List[T]) Remove(val T) {
	if l.Head == nil {
		return
	}

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
		if !l.isEqualFunc(curr.Value, val) {
			return false
		}

		// 1. prev -> curr -> next
		// 2. remove newNode
		// 	- 2.a. curr.Prev.Next -> curr.Next == prev -> next
		// 	- 2.b. curr.Next.Prev -> curr.Prev ======= prev <- next

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

func (l *List[T]) Get(val T) *T {
	var (
		found bool
		node  *T
	)

	l.Head.Each(func(curr *Node[T]) bool {
		if !l.isEqualFunc(curr.Value, val) {
			return false
		}

		node = &curr.Value
		found = true
		return found
	})

	return node
}

func (l *List[T]) Size() int {
	var size int

	l.Head.Each(func(n *Node[T]) bool {
		size++
		return false
	})

	return size
}

func (l *List[T]) Clear() {
	l.Head = nil
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
