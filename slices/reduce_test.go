package slices_test

import (
	"fmt"
	"testing"

	"github.com/igorroncevic/go-utils/slices"
)

type reduceTestCase[T, M any] struct {
	Name           string
	Slice          []T
	Reducer        func(M, T) M
	InitialValue   M
	ExpectedResult M
}

func runReduceTestCases[T, M any](t *testing.T, testCases []reduceTestCase[T, M]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := slices.Reduce(tc.Slice, tc.Reducer, tc.InitialValue)

			AssertEqualFormatted(t, tc.ExpectedResult, result)
		})
	}
}

func TestReduce(t *testing.T) {
	runReduceTestCases[int](t, getIntReduceTestCases())
	runReduceTestCases[string](t, getStringReduceTestCases())
}

func getIntReduceTestCases() []reduceTestCase[int, int] {
	addElem := func(acc, elem int) int { return acc + elem }
	subElem := func(acc, elem int) int { return acc - elem }

	return []reduceTestCase[int, int]{
		{
			Name:           "int - no elements in the slice",
			Slice:          []int{},
			Reducer:        addElem,
			InitialValue:   0,
			ExpectedResult: 0,
		},
		{
			Name:           "int - adds one",
			Slice:          []int{1, 2, 3},
			Reducer:        addElem,
			InitialValue:   0,
			ExpectedResult: 6,
		},
		{
			Name:           "int - subtracts one",
			Slice:          []int{1, 2, 3, 4},
			Reducer:        subElem,
			InitialValue:   0,
			ExpectedResult: -10,
		},
	}
}

func getStringReduceTestCases() []reduceTestCase[string, string] {
	appendElem := func(acc, elem string) string {
		if acc == "" {
			return elem
		}

		return fmt.Sprintf("%s, %s", acc, elem)
	}

	return []reduceTestCase[string, string]{
		{
			Name:           "string - no elements in the slice",
			Slice:          []string{},
			Reducer:        appendElem,
			InitialValue:   "",
			ExpectedResult: "",
		},
		{
			Name:           "string - 3 elements in the slice",
			Slice:          []string{"this", "that", "the other"},
			Reducer:        appendElem,
			InitialValue:   "",
			ExpectedResult: "this, that, the other",
		},
	}
}
