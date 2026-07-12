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
	// Construct multiple sets of data in a slice
	// Alternatively could create an expected field with the arranged input but that's manual and I ain't go time for that.
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

	// Iterate through the data sets]
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedNumbers := SelectionSort(test.input)
			// Using slices.Sorted to sort the data and return a brand-new sorted slice
			// this slice is then stored in the want to be compared with the response from BubbleSort
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedNumbers) {
				t.Errorf("got: %v, want: %v", test.input, want)
			}
		})
	}
}