package sort

import (
	"testing"
	"slices"
)

type TestCases struct {
	name string
	input []int
}

func TestBubbleSort(t *testing.T) {
	ArraySorter := NewArraySort([]int{})

	tests := []TestCases{
		{
			name: "Bubble Sort: Test One",
			input: RandomArrayGeneratorForTests(20),
		},
		{
			name: "Bubble Sort: Test Two",
			input: RandomArrayGeneratorForTests(25),
		},
		{
			name: "Bubble Sort: Test Three",
			input: RandomArrayGeneratorForTests(30),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedArray := ArraySorter.BubbleSort(test.input)
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedArray) {
				t.Errorf("\n Got: %v \n Want: %v \n", sortedArray, want)
				t.Logf("\n Got: %v \n Want: %v \n", sortedArray, want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	ArraySorter := NewArraySort([]int{})

	tests := []TestCases{
		{
			name: "Selection Sort Test One",
			input: RandomArrayGeneratorForTests(25),
		},
		{
			name: "Selection Sort Test Two",
			input: RandomArrayGeneratorForTests(30),
		},
		{
			name: "Selection Sort Test Three",
			input: RandomArrayGeneratorForTests(35),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedArray := ArraySorter.SelectionSort(test.input)
			want := slices.Sorted(slices.Values(test.input))
			if !slices.IsSorted(sortedArray) {
				t.Errorf("\n Got: %v \n Want: %v \n", sortedArray, want)
				t.Logf("\n Got: %v \n Want: %v \n", sortedArray, want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {

}

func TestShellSort(t *testing.T) {

}

func TestMergeSort(t *testing.T) {

}

func TestQuickSort(t *testing.T) {

}