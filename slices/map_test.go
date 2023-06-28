package slices_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/igorroncevic/go-utils/slices"
)

type mapTestCase[T any] struct {
	Name           string
	Slice          []T
	Mapper         func(T) T
	ExpectedResult []T
}

func runMapTestCases[T any](t *testing.T, testCases []mapTestCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := slices.Map(tc.Slice, tc.Mapper)

			AssertEqualSlicesLength(t, tc.ExpectedResult, result)

			for i, _ := range result {
				AssertEqualSlicesFormatted(t, tc.ExpectedResult, result, i)
			}
		})
	}
}

func TestMap(t *testing.T) {
	runMapTestCases[int](t, getMapIntTestCases())
	runMapTestCases[string](t, getMapStringTestCases())
	runMapTestCases[testStruct](t, getMapStructTestCases())
}

func getMapIntTestCases() []mapTestCase[int] {
	return []mapTestCase[int]{
		{
			Name:           "int - no elements in the slice",
			Slice:          []int{},
			Mapper:         func(a int) int { return a + 1 },
			ExpectedResult: []int{},
		},
		{
			Name:           "int - increase all elements by 1",
			Slice:          []int{1, 2, 3},
			Mapper:         func(a int) int { return a + 1 },
			ExpectedResult: []int{2, 3, 4},
		},
	}
}

func getMapStringTestCases() []mapTestCase[string] {
	return []mapTestCase[string]{
		{
			Name:           "string - no elements in the slice",
			Slice:          []string{},
			Mapper:         func(a string) string { return fmt.Sprintf("%s - more text", a) },
			ExpectedResult: []string{},
		},
		{
			Name:           "string - append text to the end of the string",
			Slice:          []string{"a", "b", "c"},
			Mapper:         func(a string) string { return fmt.Sprintf("%s - more text", a) },
			ExpectedResult: []string{"a - more text", "b - more text", "c - more text"},
		},
	}
}

func getMapStructTestCases() []mapTestCase[testStruct] {
	now := time.Now()

	filledStruct := testStruct{
		StringField: "test",
		IntField:    101,
		TimeField:   now.Add(10 * time.Hour),
		BoolField:   true,
		MapField: map[string]interface{}{
			"other": "field 2",
		},
	}

	expectedFilledStruct := filledStruct
	expectedFilledStruct.StringField = "test - added text"
	expectedFilledStruct.IntField += 100
	expectedFilledStruct.TimeField = expectedFilledStruct.TimeField.Add(1 * time.Hour)
	expectedFilledStruct.BoolField = false
	delete(expectedFilledStruct.MapField, "other")

	return []mapTestCase[testStruct]{
		{
			Name:           "testStruct - no elements in the slice",
			Slice:          []testStruct{},
			Mapper:         func(ts testStruct) testStruct { return ts },
			ExpectedResult: []testStruct{},
		},
		{
			Name:           "testStruct - no change",
			Slice:          []testStruct{getDefaultTestStruct(now)},
			Mapper:         func(ts testStruct) testStruct { return ts },
			ExpectedResult: []testStruct{getDefaultTestStruct(now)},
		},
		{
			Name:  "testStruct - filter by all fields",
			Slice: []testStruct{filledStruct},
			Mapper: func(ts testStruct) testStruct {
				ts.StringField = fmt.Sprintf("%s - added text", ts.StringField)
				ts.IntField += 100
				ts.TimeField = ts.TimeField.Add(1 * time.Hour)
				ts.BoolField = false
				delete(ts.MapField, "other")

				return ts
			},
			ExpectedResult: []testStruct{expectedFilledStruct},
		},
	}
}
