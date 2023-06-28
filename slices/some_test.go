package slices_test

import (
	"strings"
	"testing"
	"time"

	"github.com/igorroncevic/go-utils/slices"
)

type someTestCase[T any] struct {
	Name           string
	Slice          []T
	Filtrator      func(T) bool
	ExpectedResult bool
}

func runSomeTestCases[T any](t *testing.T, testCases []someTestCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := slices.Some(tc.Slice, tc.Filtrator)

			AssertEqualFormatted(t, tc.ExpectedResult, result)
		})
	}
}

func TestSome(t *testing.T) {
	runSomeTestCases[int](t, getIntSomeTestCases())
	runSomeTestCases[string](t, getStringSomeTestCases())
	runSomeTestCases[testStruct](t, getStructSomeTestCases())
}

func getIntSomeTestCases() []someTestCase[int] {
	greaterThanZero := func(a int) bool { return a > 0 }

	return []someTestCase[int]{
		{
			Name:           "int - no elements in the slice",
			Slice:          []int{},
			Filtrator:      greaterThanZero,
			ExpectedResult: false,
		},
		{
			Name:           "int - 3 elements in the slice, none match",
			Slice:          []int{-1, -2, -3},
			Filtrator:      greaterThanZero,
			ExpectedResult: false,
		},
		{
			Name:           "int - 3 elements in the slice, one matches",
			Slice:          []int{-1, 0, 1},
			Filtrator:      greaterThanZero,
			ExpectedResult: true,
		},
	}
}

func getStringSomeTestCases() []someTestCase[string] {
	stringSomeFiltrator := func(a string) bool { return strings.Contains(a, "some") }

	return []someTestCase[string]{
		{
			Name:           "string - no elements in the slice",
			Slice:          []string{},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: false,
		},
		{
			Name:           "string - 3 elements in the slice, but none match",
			Slice:          []string{"this", "that", "the other"},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: false,
		},
		{
			Name:           "string - 3 elements in the slice, but only one matches",
			Slice:          []string{"somebody", "that i", "used to know"},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: true,
		},
	}
}

func getStructSomeTestCases() []someTestCase[testStruct] {
	now := time.Now()

	filledStruct := testStruct{
		StringField: "test",
		IntField:    101,
		TimeField:   now.Add(10 * time.Hour),
		BoolField:   true,
		MapField: map[string]interface{}{
			"other": "field",
		},
	}

	return []someTestCase[testStruct]{
		{
			Name:           "testStruct - no elements in the slice",
			Slice:          []testStruct{},
			Filtrator:      func(ts testStruct) bool { return ts.BoolField && ts.StringField != "" },
			ExpectedResult: false,
		},
		{
			Name:           "testStruct - 1 elements in the slice, but doesn't match",
			Slice:          []testStruct{getDefaultTestStruct(now)},
			Filtrator:      func(ts testStruct) bool { return ts.IntField > 100 },
			ExpectedResult: false,
		},
		{
			Name: "testStruct - filter by all fields",
			Slice: []testStruct{
				getDefaultTestStruct(now),
				filledStruct,
			},
			Filtrator: func(ts testStruct) bool {
				return ts.TimeField.After(now.Add(1 * time.Hour))
			},
			ExpectedResult: true,
		},
	}
}
