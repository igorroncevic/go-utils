package slices_test

import (
	"strings"
	"testing"
	"time"

	"github.com/igorroncevic/go-utils/slices"
)

type everyTestCase[T any] struct {
	Name           string
	Slice          []T
	Filtrator      func(T) bool
	ExpectedResult bool
}

func runEveryTestCases[T any](t *testing.T, testCases []everyTestCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := slices.Every(tc.Slice, tc.Filtrator)

			AssertEqualFormatted(t, tc.ExpectedResult, result)
		})
	}
}

func TestEvery(t *testing.T) {
	runEveryTestCases[int](t, getIntEveryTestCases())
	runEveryTestCases[string](t, getStringEveryTestCases())
	runEveryTestCases[testStruct](t, getStructEveryTestCases())
}

func getIntEveryTestCases() []everyTestCase[int] {
	greaterThanZero := func(a int) bool { return a > 0 }

	return []everyTestCase[int]{
		{
			Name:           "int - no elements in the slice",
			Slice:          []int{},
			Filtrator:      greaterThanZero,
			ExpectedResult: true,
		},
		{
			Name:           "int - 3 elements in the slice, but only one matches",
			Slice:          []int{-1, 0, 1},
			Filtrator:      greaterThanZero,
			ExpectedResult: false,
		},
		{
			Name:           "int - 3 elements in the slice, all match",
			Slice:          []int{1, 2, 3},
			Filtrator:      greaterThanZero,
			ExpectedResult: true,
		},
	}
}

func getStringEveryTestCases() []everyTestCase[string] {
	stringSomeFiltrator := func(a string) bool { return strings.Contains(a, "some") }

	return []everyTestCase[string]{
		{
			Name:           "string - no elements in the slice",
			Slice:          []string{},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: true,
		},
		{
			Name:           "string - 3 elements in the slice, but none match",
			Slice:          []string{"this", "that", "the other"},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: false,
		},
		{
			Name:           "string - 3 elements in the slice, but only one matches",
			Slice:          []string{"somebody", "something", "somewhere"},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: true,
		},
	}
}

func getStructEveryTestCases() []everyTestCase[testStruct] {
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

	return []everyTestCase[testStruct]{
		{
			Name:           "testStruct - no elements in the slice",
			Slice:          []testStruct{},
			Filtrator:      func(ts testStruct) bool { return ts.BoolField && ts.StringField != "" },
			ExpectedResult: true,
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
				if _, ok := ts.MapField["other"]; ok {
					return true
				}

				return false
			},
			ExpectedResult: true,
		},
	}
}
