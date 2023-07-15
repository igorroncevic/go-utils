package queue_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/igorroncevic/go-utils/queue"
)

func TestQueue(t *testing.T) {
	st := queue.New[int]()

	// Start enqueueing
	st.Enqueue(0)
	val1, err1 := st.Peek()
	assert.NoError(t, err1)
	assert.Equal(t, 0, *val1)

	st.Enqueue(42)
	val2, err2 := st.Peek()
	assert.NoError(t, err2)
	assert.Equal(t, 0, *val2)

	assert.Equal(t, 2, st.Len())

	// Start dequeueing
	deq2, err3 := st.Dequeue()
	assert.NoError(t, err3)
	assert.Equal(t, 0, *deq2)

	deq1, err4 := st.Dequeue()
	assert.NoError(t, err4)
	assert.Equal(t, 42, *deq1)

	assert.Equal(t, 0, st.Len())

	// Queue empty
	deqNone, err5 := st.Dequeue()
	assert.Error(t, err5)
	assert.Nil(t, deqNone)
}
