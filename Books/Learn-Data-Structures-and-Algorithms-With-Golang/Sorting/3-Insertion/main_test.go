package main

import (
	"testing"
	"slices"
)

type TestCases struct {
	name string
	input []int
}


func TestInsertionSort(t *testing.T) {
	tests := []TestCases{
		{
			name: "Test One",
			input: RandomNumberArrayGenerator(10),
		},
		{
			name: "Test Two",
			input: RandomNumberArrayGenerator(11),
		},
		{
			name: "Test Three",
			input: RandomNumberArrayGenerator(12),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T) {
			sortedArray := InsertionSort(test.input)
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedArray) {
				t.Errorf("got: %v, want %v", sortedArray, want)
			}
		})
	}
}

