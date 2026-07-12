package main

import "fmt"

func main() {
	integers := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}

	sortedIntegers := BubbleSort(integers)

	fmt.Println("Bubble Sort Outcome: ", sortedIntegers)
}

func BubbleSort(numbers []int) []int {

	isSwapped := true

	// Keep looping as long as least one swap happened during the previous pass
	for isSwapped {
		// if array is perfectly sorted, it would break the first loop
		isSwapped = false
		for i := 1; i < len(numbers); i++ {
			// second loop goes through all the numbers in the array and "bubbles" the biggest number to the end.
			// each time a swap is made, the isSwapped is set to true which would trigger the outer loop AFTER
			// the inner loop is done.
			if numbers[i-1] > numbers[i] {
				temp := numbers[i]
				numbers[i] = numbers[i-1]
				numbers[i-1] = temp
				isSwapped = true
			}
		}
	}

	return numbers
}

/*
Was trying to be too smart but good syntax to remember:
	// get the length of the passed array
	arrlen := len(numbers)

	// create an empty array of corresponding length to hold the sorted elements
	sortedNumbers := make([]int, arrlen)

	// Use Go's multiple assignment to swap in one line
    numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
    isSwapped = true
*/
