package main

import (
	"slices"
	"testing"
)

// Defining the struct for the case data set
type testCases struct {
	name  string
	input []int
}

func TestBubbleSort(t *testing.T) {
	// Construct multiple sets of data in a slice
	// Alternatively could create an expected field with the arranged input but that's manual and I ain't go time for that.
	tests := []testCases{
		{
			name:  "TestOne",
			input: []int{16, 5, 3, 4, 11, 20},
		},
		{
			name:  "TestTwo",
			input: []int{30, 42, 109, 56, 3, 99},
		},
		{
			name:  "TestThree",
			input: []int{12, 1, 5, 55, 60, 49, 100},
		},
	}

	// Iterate through the data sets]
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sortedNumbers := BubbleSort(tc.input)
			// Using slices.Sorted to sort the data and return a brand-new sorted slice
			// this slice is then stored in the want to be compared with the response from BubbleSort
			want := slices.Sorted(slices.Values(tc.input))
			if !slices.IsSorted(sortedNumbers) {
				t.Errorf("got %v, want %v", tc.input, want)
			}
		})
	}
}

/*
Legacy Solution if using Go 1.22 or lower
// 1. Create a clean copy of the input data
want := slices.Clone(tc.cases)

// 2. Sort the copy in place (modifies 'want' directly)
slices.Sort(want)

// Now 'want' holds the sorted data, and tc.cases remains untouched!


*/
