package main

import (
	"slices"
	"testing"
)

type TestCases struct {
	name string
	input []int
}

func TestQuickSort(t *testing.T) {
	tests := []TestCases{
		{
			name: "Test One",
			input: RandomlyGeneratedArray(15),
		},
		{
			name: "Test Two",
			input: RandomlyGeneratedArray(25),
		},
		{
			name: "Test Three",
			input: RandomlyGeneratedArray(50),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedArray := QuickSort(test.input, 0, (len(test.input)-1))
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedArray) {
				t.Errorf("\n got: %v \n want: %v", sortedArray, want)
			}
			t.Logf("\n got: %v \n want: %v", sortedArray, want)
		})
	}
}