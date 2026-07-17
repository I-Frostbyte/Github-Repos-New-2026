package sort

import (
	"math/rand"

	// "google.golang.org/genproto/googleapis/appengine/legacy"
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

// MergeSort implements the Merge Sort algorithm.
// This algorithm makes use of recursiveness.
// This function has a supplementary joiner function.
// The MergeSort breaks down the input array in two and 
// calls the joiner function and passes in as arguments,
// the two halves of the input array but passes them as arguments
// of the MergeSort, initiating the recursiveness.
// This breaks the array down till it gets to an array of just two elements.
// It sorts those two and then recursively begins to sort the others, moving
// up a tree like structure. Like so:
/*

                       [a b c d e f g h i j]
                      /                     \
            [a b c d e]                  [f g h i j]
            /         \                  /         \
        [a b]     [c d e]            [f g]     [h i j]
        /   \      /    \            /   \     /     \
      [a]   [b]  [c]  [d e]        [f]   [g]  [h]  [i j]
                       /  \                        /   \
                      [d] [e]                     [i]  [j]

*/
func (as *ArraySort) MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	middle := (len(input)) / 2

	return JoinArrays(as.MergeSort(input[:middle]), as.MergeSort(input[middle:]))
}

func (as *ArraySort) QuickSort(input []int) []int {
	panic ("unimplemented")
}

// RandomArrayGeneratorForTests generates a random array of integers between 0 & 50.
// These integers can repeat themselves.
// This array is solely for testing purposes.
func RandomArrayGeneratorForTests(arrLength int) []int {
	newArray := make([]int, arrLength)

	for i := range newArray {
		newArray[i] = rand.Intn(50)
	}

	return newArray
}

// SimpleSwap takes an array of integers, a first index and a second index.
// It then swaps the element in the first index with the element of the second index.
func SimpleSwap(array []int, firstIndex, secondIndex int) {
	array[firstIndex], array[secondIndex] = array[secondIndex], array[firstIndex]
}

// JoinArrays takes two arrays and joins them while sorting
// the elements of the joined array in ascending order.
func JoinArrays(leftArray, rightArray []int) []int {
	lengthOfJoinedArray, i, j := len(leftArray)+len(rightArray), 0, 0

	joinedArray := make([]int, lengthOfJoinedArray)

	for k := range joinedArray {
		if i > len(leftArray)-1 && j <= len(rightArray)-1 {
			joinedArray[k] = rightArray[j]
			j++
		} else if i <= len(leftArray)-1 && j > len(rightArray)-1 {
			joinedArray[k] = leftArray[i]
			i++
		} else if leftArray[i] < rightArray[j] {
			joinedArray[k] = leftArray[i]
			i++
		} else {
			joinedArray[k] = rightArray[j]
			j++
		}
	}

	return joinedArray
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