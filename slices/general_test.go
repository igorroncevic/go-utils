package slices_test

import (
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

type testStruct struct {
	StringField string
	IntField    int
	TimeField   time.Time
	BoolField   bool
	MapField    map[string]interface{}
}

func getDefaultTestStruct(now time.Time) testStruct {
	return testStruct{
		StringField: "default",
		IntField:    1,
		TimeField:   now,
		BoolField:   true,
		MapField: map[string]interface{}{
			"some": "field",
		},
	}
}

func AssertEqualFormatted[T any](t *testing.T, expected, actual []T, index int) {
	assert.Equal(
		t, expected[index], actual[index],
		"values --> expected: '%+v', actual: '%+v', at index %d, \n\tslices --> expected: %+v, actual: %+v",
		expected[index], expected[index], index, expected, actual,
	)
}

func AssertEqualLength[T any](t *testing.T, expected, actual []T) {
	assert.Equal(
		t, len(expected), len(actual),
		"expected length = %d, actual length = %d\n\tslices --> expected: %+v, actual: %+v",
		len(expected), len(actual), expected, actual,
	)
}
