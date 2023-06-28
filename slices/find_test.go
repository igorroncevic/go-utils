package slices_test

import (
	"testing"

	"github.com/igorroncevic/go-utils/slices"
	"github.com/igorroncevic/go-utils/util"
)

type findTestCase[T any] struct {
	Name           string
	Slice          []T
	ToFind         T
	ExpectedResult *T
	ExpectedError  error
}

func runFindTestCases[T comparable](t *testing.T, testCases []findTestCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := slices.Find(tc.Slice, tc.ToFind)

			AssertEqualFormatted[error](t, tc.ExpectedError, err)

			if tc.ExpectedResult != nil && result != nil {
				AssertEqualFormatted[T](t, *tc.ExpectedResult, *result)
			}
		})
	}
}

func TestFind(t *testing.T) {
	runFindTestCases[int](t, getIntFindTestCases())
	runFindTestCases[string](t, getStringFindTestCases())
}

func getIntFindTestCases() []findTestCase[int] {
	return []findTestCase[int]{
		{
			Name:           "int - no elements in the slice",
			Slice:          []int{},
			ToFind:         1,
			ExpectedResult: nil,
			ExpectedError:  slices.ErrNotFound,
		},
		{
			Name:           "int - 3 elements in the slice, one matches",
			Slice:          []int{-1, 0, 1},
			ToFind:         1,
			ExpectedResult: util.Ptr(1),
			ExpectedError:  nil,
		},
	}
}

func getStringFindTestCases() []findTestCase[string] {
	return []findTestCase[string]{
		{
			Name:           "string - no elements in the slice",
			Slice:          []string{},
			ToFind:         "hello",
			ExpectedResult: nil,
			ExpectedError:  slices.ErrNotFound,
		},
		{
			Name:           "string - 3 elements in the slice, but none match",
			Slice:          []string{"this", "that", "the other"},
			ToFind:         "hello",
			ExpectedResult: nil,
			ExpectedError:  slices.ErrNotFound,
		},
		{
			Name:           "string - 3 elements in the slice, but only one matches",
			Slice:          []string{"somebody", "that i", "used to know"},
			ToFind:         "somebody",
			ExpectedResult: util.Ptr("somebody"),
			ExpectedError:  nil,
		},
	}
}
