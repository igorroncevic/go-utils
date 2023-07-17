package list_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/igorroncevic/go-utils/list"
)

func lessFn(a, b int) bool {
	return a < b
}

func equalFn(a, b int) bool {
	return a == b
}

func TestList(t *testing.T) {
	linkedList := list.New[int](lessFn, equalFn)

	// Add elements
	linkedList.Push(1)
	linkedList.Push(2)
	linkedList.Push(5)

	assert.Equal(t, 3, linkedList.Size(), "unexpected list size")
	assert.EqualValues(t, []int{1, 2, 5}, linkedList.ToSlice(), "unexpected slice values")

	// Remove the head
	linkedList.Remove(1)

	assert.Nil(t, linkedList.Get(1), "element was not removed")
	assert.Equal(t, 2, linkedList.Size(), "unexpected list size after remove")
	assert.EqualValues(t, []int{2, 5}, linkedList.ToSlice(), "unexpected slice values after remove")

	// Add more elements
	linkedList.Push(10)
	linkedList.Push(3)
	linkedList.Push(6)
	linkedList.Push(4)

	assert.Equal(t, 6, linkedList.Size(), "unexpected list size after remove and another add")
	assert.EqualValues(t, []int{2, 3, 4, 5, 6, 10}, linkedList.ToSlice(), "unexpected slice values after remove and another add")
	assert.True(t, linkedList.Contains(10), "element is not in the list")

	// Remove the middle element
	linkedList.Remove(4)

	assert.Nil(t, linkedList.Get(4), "element was not removed")
	assert.Equal(t, 5, linkedList.Size(), "unexpected list size after remove")
	assert.EqualValues(t, []int{2, 3, 5, 6, 10}, linkedList.ToSlice(), "unexpected slice values after remove")

	// Clear the list
	linkedList.Clear()

	assert.Equal(t, 0, linkedList.Size(), "unexpected list size after clear")
}
