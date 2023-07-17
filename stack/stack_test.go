package stack_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/igorroncevic/go-utils/stack"
)

func TestStack(t *testing.T) {
	st := stack.New[int]()

	// Start pushing
	st.Push(0)
	val1, err1 := st.Peek()
	assert.NoError(t, err1)
	assert.Equal(t, 0, *val1)

	st.Push(42)
	val2, err2 := st.Peek()
	assert.NoError(t, err2)
	assert.Equal(t, 42, *val2)

	assert.Equal(t, 2, st.Size())

	// Start popping
	pop2, err3 := st.Pop()
	assert.NoError(t, err3)
	assert.Equal(t, 42, *pop2)

	pop1, err4 := st.Pop()
	assert.NoError(t, err4)
	assert.Equal(t, 0, *pop1)

	assert.Equal(t, 0, st.Size())

	// Stack empty
	popNone, err5 := st.Pop()
	assert.Error(t, err5)
	assert.Nil(t, popNone)
}
