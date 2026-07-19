package main

import (
	"testing"

	
	"github.com/I-Frostbyte/Github-Repos-New-2026/Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/generator"
)

type arrGenRes struct {
	array []int
	isElementFound bool
	element int
	err error
}

func NewInput(arrLength, elementToFind int, isElementFound bool) arrGenRes {
	arrayGenerator := generator.NewGenerator()
	testInput, isElementFound, err := arrayGenerator.RandomIntegerArrayGenerator(arrLength, elementToFind, isElementFound)
	return arrGenRes{
		array: *testInput,
		isElementFound: isElementFound,
		element: elementToFind,
		err: err,
	}
}

type TestCases struct {
	name string
	input arrGenRes
}

func TestLinearSearch(t *testing.T) {
	tests := []TestCases{
		{
			name: "Linear Search Test One",
			input: NewInput(25, 15, true),
		},
		{
			name: "Linear Search Test Two",
			input: NewInput(30, 8, true),
		},
		{
			name: "Linear Search Test Three",
			input: NewInput(35, 27, true),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wasElementFound := LinearSearch(test.input.array, test.input.element)
			want := test.input.isElementFound
			if wasElementFound != want {
				t.Errorf("\n Got: %v \n Want: %v \n", wasElementFound, want)
				t.Logf("\n Got: %v \n Want: %v \n", wasElementFound, want)
			}
		})
	}
}