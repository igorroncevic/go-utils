package set_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/igorroncevic/go-utils/set"
)

func TestSet(t *testing.T) {
	s := set.New[int]()

	// Add elements
	err := s.Add(1)
	assert.NoError(t, err, "error adding '1' to the set")

	err = s.Add(2)
	assert.NoError(t, err, "error adding '2' to the set")

	contains := s.Contains(1)
	assert.True(t, contains, "set does not contain value '1'")

	contains = s.Contains(99)
	assert.False(t, contains, "set does contain value '99'")

	checker := set.New[int]()

	s.Each(func(val int) {
		err := checker.Add(val)
		assert.NoError(t, err)
	})

	assert.Equal(t, s, checker)

	// Add already existing element
	err = s.Add(1)
	assert.Error(t, err, "no error while adding existing element ('1') to the set")

	assert.Equal(t, 2, s.Size(), "unexpected set size")

	// Remove elements
	err = s.Remove(1)
	assert.NoError(t, err, "error removing value '1'")

	assert.Equal(t, 1, s.Size(), "unexpected set size")

	err = s.Remove(99)
	assert.Error(t, err)

	assert.Equal(t, 1, s.Size(), "unexpected set size")

	s.Clear()
	assert.Equal(t, 0, s.Size(), "unexpected set size")
}
