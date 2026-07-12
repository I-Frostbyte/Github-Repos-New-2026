package main

import (
	"testing"
	"slices"
)

type TestCases struct {
	name string
	input []int
}

func TestSelectionSort(t *testing.T) {
	tests := []TestCases{
		{
			name: "TestOne",
			input: []int{2, 1, 5, 12, 16, 22, 34, 11, 3, 30},
		},
		{
			name: "TestTwo",
			input: []int{34, 48, 57, 12, 43, 4, 6, 17, 47},
		},
		{
			name: "TestThree",
			input: []int{46, 12, 22, 10, 8, 56, 98, 23},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedNumbers := SelectionSort(test.input)
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedNumbers) {
				t.Errorf("got: %v, want: %v", test.input, want)
			}
		})
	}
}