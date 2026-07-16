package sort

import (
	"math/rand"
)
type ArraySort struct {
	input []int
}

func NewArraySort(
	input []int,
	) *ArraySort {
		return &ArraySort{
			input: input,
		}
}

// BubbleSort implements the Bubble Sort algorithm.
// In this algorithm, each number in the array is compared to the one next to it.
// The bigger number is shifted to the right and the loop restarts when a switch is made.
// This essentially moves the biggest number to the end and every other number builds up to it.
func (as *ArraySort) BubbleSort(input []int) []int {
	isSwapped := true

	for isSwapped {
		isSwapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i-1], input[i] = input[i], input[i-1]
				isSwapped = true
			}
		}
	}

	return input
}

func (as *ArraySort) SelectionSort(input []int) []int {
	panic ("unimplemented")
}

func (as *ArraySort) InsertionSort(input []int) []int {
	panic ("unimplemented")
}

func (as *ArraySort) ShellSort(input []int) []int {
	panic ("unimplemented")
}

func (as *ArraySort) MergeSort(input []int) []int {
	panic ("unimplemented")
}

func (as *ArraySort) QuickSort(input []int) []int {
	panic ("unimplemented")
}

func RandomArrayGeneratorForTests(arrLength int) []int {
	newArray := make([]int, arrLength)

	for i := range newArray {
		newArray[i] = rand.Intn(50)
	}

	return newArray
}