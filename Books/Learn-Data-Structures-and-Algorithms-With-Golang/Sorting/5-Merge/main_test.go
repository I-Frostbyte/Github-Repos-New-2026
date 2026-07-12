package main

import (
	"slices"
	"testing"
)

type TestCases struct {
	name string
	input []int
}

func TestMergeSort(t *testing.T) {
	tests := []TestCases{
		{
			name: "Test One",
			input: RandomlyGeneratedArray(45),
		},
		{
			name: "Test Two",
			input: RandomlyGeneratedArray(60),
		},
		{
			name: "Test Three",
			input: RandomlyGeneratedArray(90),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedArray := MergeSort(test.input)
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedArray) {
				t.Errorf("got %v: want %v", sortedArray, want)
			}
			t.Logf("\n got: %v \n want: %v", sortedArray, want)
		})
	}
}