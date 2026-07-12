package main

import (
	"slices"
	"testing"
)

type TestCases struct {
	name string
	input []int
}

func TestShellSort(t *testing.T) {
	tests := []TestCases{
		{
			name: "Test One",
			input: RandomNumberArrayGenerator(32),
		},
		{
			name: "Test Two",
			input: RandomNumberArrayGenerator(35),
		},
		{
			name: "Test Three",
			input: RandomNumberArrayGenerator(40),
		},
	}

	for _, test :=  range tests {
		t.Run(test.name, func (t *testing.T) {
			sortedArray := ShellSort(test.input)
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedArray) {
				t.Errorf("got %v: want %v", sortedArray, want)
			}
		})
	}
}