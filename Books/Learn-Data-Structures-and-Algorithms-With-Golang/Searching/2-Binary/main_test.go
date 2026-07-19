package main

import (
	"slices"
	"testing"

	"github.com/I-Frostbyte/Github-Repos-New-2026/Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/generator"
)

type ArrGenRes struct {
	array        []int
	element      int
	isElementFound bool
	err          error
}

func NewInput(arrayLength, element int, isElementFound bool) ArrGenRes {
	generator := generator.NewGenerator()
	array, isElementFound, err := generator.RandomIntegerArrayGenerator(arrayLength, element, isElementFound)
	
	return ArrGenRes{
		array: *array,
		element: element,
		isElementFound: isElementFound,
		err: err,
	}
}

type TestCases struct {
	name  string
	input ArrGenRes
}

func TestBinarySearch(t *testing.T) {
	tests := []TestCases{
		{
			name: "Binary Search Test One",
			input: NewInput(25, 10, true),
		},
		{
			name: "Binary Search Test Two",
			input: NewInput(30, 22, true),
		},
		{
			name: "Binary Search Test Three",
			input: NewInput(35, 5, true),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wasElementFound := BinarySearch(slices.Sorted(slices.Values(test.input.array)), test.input.element)
			want := test.input.isElementFound

			if wasElementFound != want {
				t.Errorf("\n Got: %v \n Want: %v", wasElementFound, want)
				t.Logf("\n Got: %v \n Want: %v", wasElementFound, want)
			}
		})
	}
}
