package slices_test

import (
	"strings"
	"testing"
	"time"

	"github.com/igorroncevic/go-utils/slices"
)

type filterTestCase[T any] struct {
	Name           string
	Slice          []T
	Filtrator      func(T) bool
	ExpectedResult []T
}

func runFilterTestCases[T any](t *testing.T, testCases []filterTestCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := slices.Filter(tc.Slice, tc.Filtrator)

			AssertEqualSlicesLength(t, tc.ExpectedResult, result)

			for i, _ := range result {
				AssertEqualSlicesFormatted(t, tc.ExpectedResult, result, i)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	runFilterTestCases[int](t, getIntFilterTestCases())
	runFilterTestCases[string](t, getStringFilterTestCases())
	runFilterTestCases[testStruct](t, getStructFilterTestCases())
}

func getIntFilterTestCases() []filterTestCase[int] {
	const (
		MaxUint = ^uint(0)
		MinUint = 0
		MaxInt  = int(MaxUint >> 1)
		MinInt  = -MaxInt - 1
	)

	greaterThanZero := func(a int) bool { return a > 0 }

	return []filterTestCase[int]{
		{
			Name:           "int - no elements in the slice",
			Slice:          []int{},
			Filtrator:      greaterThanZero,
			ExpectedResult: []int{},
		},
		{
			Name:           "int - 3 elements in the slice, but only one matches",
			Slice:          []int{-1, 0, 1},
			Filtrator:      greaterThanZero,
			ExpectedResult: []int{1},
		},
		{
			Name:           "int - working with boundary values",
			Slice:          []int{MinUint, MaxInt, MinInt},
			Filtrator:      greaterThanZero,
			ExpectedResult: []int{MaxInt},
		},
	}
}

func getStringFilterTestCases() []filterTestCase[string] {
	stringSomeFiltrator := func(a string) bool { return strings.Contains(a, "some") }

	return []filterTestCase[string]{
		{
			Name:           "string - no elements in the slice",
			Slice:          []string{},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: []string{},
		},
		{
			Name:           "string - 3 elements in the slice, but none match",
			Slice:          []string{"this", "that", "the other"},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: []string{},
		},
		{
			Name:           "string - 3 elements in the slice, but only one matches",
			Slice:          []string{"somebody", "that i", "used to know"},
			Filtrator:      stringSomeFiltrator,
			ExpectedResult: []string{"somebody"},
		},
		{
			Name:           "string - utf8 letters",
			Slice:          []string{"Äußerst", "Öffnen", "Übermorgen", "szabályai", "folyó", "Őrség", "Lőrinc"},
			Filtrator:      func(a string) bool { return strings.ContainsAny(a, "ÄÖÜáóŐő") },
			ExpectedResult: []string{"Äußerst", "Öffnen", "Übermorgen", "szabályai", "folyó", "Őrség", "Lőrinc"},
		},
	}
}

func getStructFilterTestCases() []filterTestCase[testStruct] {
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

	return []filterTestCase[testStruct]{
		{
			Name:           "testStruct - no elements in the slice",
			Slice:          []testStruct{},
			Filtrator:      func(ts testStruct) bool { return ts.BoolField && ts.StringField != "" },
			ExpectedResult: []testStruct{},
		},
		{
			Name:           "testStruct - 1 elements in the slice, but doesn't match",
			Slice:          []testStruct{getDefaultTestStruct(now)},
			Filtrator:      func(ts testStruct) bool { return ts.IntField > 100 },
			ExpectedResult: []testStruct{},
		},
		{
			Name: "testStruct - filter by all fields",
			Slice: []testStruct{
				getDefaultTestStruct(now),
				filledStruct,
			},
			Filtrator: func(ts testStruct) bool {
				if _, ok := ts.MapField["other"]; ok {
					return ts.StringField != "" && ts.IntField > 100 && ts.TimeField.After(now) && ts.BoolField
				}

				return false
			},
			ExpectedResult: []testStruct{filledStruct},
		},
	}
}
