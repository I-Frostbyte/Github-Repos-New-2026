package main

import (
	"testing"

	
	"github.com/I-Frostbyte/Github-Repos-New-2026/Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/generator"
	"github.com/I-Frostbyte/Github-Repos-New-2026/Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/generator"
)

type arrGenRes struct {
	array []int
	found bool
	element int
	err error
}

type TestCases struct {
	name string
	input arrGenRes
}

func TestLinearSearch(t *testing.T) {
	arrayGenerator := generator.NewGenerator([]int{})
	foundElementOne := 15
	foundElementTwo := 8
	foundElementThree := 27
	testOneInput, foundOne, _ := arrayGenerator.RandomIntegerArrayGenerator(25, foundElementOne, true)
	testTwoInput, foundTwo, _ := arrayGenerator.RandomIntegerArrayGenerator(30, foundElementTwo, true)
	testThreeInput, foundThree, _ := arrayGenerator.RandomIntegerArrayGenerator(35, foundElementThree, true)
	
	tests := []TestCases{
		{
			name: "Linear Search Test One",
			input: arrGenRes{
				array: *testOneInput,
				found: foundOne,
				element: 15,
			},
		},
		{
			name: "Linear Search Test Two",
			input: arrGenRes{
				array: *testTwoInput,
				found: foundTwo,
				element: 8,
			},
		},
		{
			name: "Linear Search Test Three",
			input: arrGenRes{
				array: *testThreeInput,
				found: foundThree,
				element: foundElementThree,
			},
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