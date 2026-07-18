package main

import (
	"testing"

	
	"github.com/I-Frostbyte/Github-Repos-New-2026/Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/generator"
)

type arrGenRes struct {
	array []int
	found bool
	element int
	err error
}

func NewInput(arrLength, elementToFind int, lookForElement bool) arrGenRes {
	arrayGenerator := generator.NewGenerator()
	testInput, elementFound, err := arrayGenerator.RandomIntegerArrayGenerator(arrLength, elementToFind, lookForElement)
	return arrGenRes{
		array: *testInput,
		found: elementFound,
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
			foundElement := LinearSearch(test.input.array, test.input.element)
			want := test.input.found
			if foundElement != want {
				t.Errorf("\n Got: %v \n Want: %v \n", foundElement, want)
				t.Logf("\n Got: %v \n Want: %v \n", foundElement, want)
			}
		})
	}
}