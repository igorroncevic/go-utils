package slices_test

import (
	"strings"
	"testing"

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
