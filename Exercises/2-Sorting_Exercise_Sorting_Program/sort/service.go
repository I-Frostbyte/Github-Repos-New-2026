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
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		SimpleSwap(input, i, min)
	}

	return input
}

// InsertionSort implements the Insertion Sort algorithm.
// In this algorithm there are two loops. The outer loop cycles through
// the elements of the array.
// For each index, the inner loop runs which compares the value of the
// element in that index to the one beneath it and if condition is true, swaps them.
// It then descends down the left of the array, comparing every element with its
// neighbor.
// Essentially for each index, the sort counts down to 0 while checking each element
// and swapping.
func (as *ArraySort) InsertionSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			j = j-1
		}
	}

	return input
}

// ShellSort implements the Shell Sort algorithm.
// This algorithm utilizes intervals to break down the array.
// These intervals allow for sections of the array to be sorted
// in ascending manner of the intervals.
// The intervals are created by seeing how many elements can be grouped
// within the array.
// The first loop goes through the array of intervals.
// The second loop (first inner) goes through the broken down array,
// comparing elements separated by the degree of the interval.
// The final loop compares the elements at these intervals and sorts them
// before decrementing by the value of the interval.
func (as *ArraySort) ShellSort(input []int) []int {
	arrayLength := len(input)
	intervals := []int{1}
	k := 1

	for {
		interval := power(2, k) + 1
		if interval > arrayLength-1 {
			break
		}
		// new interval appended to the start of the array and not the end.
		intervals = append([]int{interval}, intervals...)
		k++
	}

	for _, interval := range intervals {
		for i := interval; i < arrayLength; i += interval {
			j := i
			for j > 0 {
				if input[j-interval] > input[j] {
					input[j-interval], input[j] = input[j], input[j-interval]
				}
				j = j - interval
			}
		}
	}

	return input

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

func power(exponent, index int) int {
	power := 1
	for index > 0 {
		if index &1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}