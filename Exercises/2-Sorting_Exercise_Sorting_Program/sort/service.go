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

// SelectionSort implements the Selection Sort algorithm.
// In this algorithm, the outer loop holds the lowest index of the array.
// While the inner loop, loops over the remaining indexes of the array.
// When a number is found lower than that in the current index of the inner loop,
// The index of that number is stored in the variable that holds the minimum.
// Note: As the inner loop progresses, the value of the minimum will change
// to hold the index of the lowest number to the right of the outer loop's current index.
// When the inner loop is done, the swap function swaps the values at the index stored in the minumum,
// with the current index held by the outer loop.
// Essentially the outer loop goes from index 0 to n-1 and for each index, places the lowest number
// found to the right of it (i.e. in the rest of the array).
func (as *ArraySort) SelectionSort(input []int) []int {
	for i := range input {
		min := i
		for j := i + 1; j <= len(input)-1; j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		SimpleSwap(input, i, min)
	}

	return input
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

func SimpleSwap(array []int, firstIndex, secondIndex int) {
	array[firstIndex], array[secondIndex] = array[secondIndex], array[firstIndex]
}