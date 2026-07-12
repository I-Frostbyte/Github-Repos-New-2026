package main

import (
	"fmt"
	"math/rand"
)

func main () {
	elements := []int{11, 4, 8, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before Sorting ", elements)
	sortedElements := ShellSort(elements)
	fmt.Println("After Sorting", sortedElements)

	fmt.Println("\n Randomly Generatated numbers: \n", RandomNumberArrayGenerator(12))
}

func ShellSort(numbers []int) []int {
	var (
		arrLength = len(numbers)
		intervals = []int{1}
		k = 1
	)

	for {
		// Loop to create the intervals for the sort.
		// Using the intervals would allow for comparison of elements not near each other and make for faster sorting.
		// The looping is determined by how many elements can be grouped within the array.
		interval := power(2, k) + 1
		if interval > arrLength-1 {
			break
		}
		intervals = append([]int{interval}, intervals...)
		k++
	}
	
	// The first loop. This loops over the array of intervals.
	for _, interval := range intervals {
		// The second loop. This loop goes over the array of numbers at increments of the intervals.
		// It starts at the element at the interval number and loops through the array based
		// on the number of times the interval can cover the array.
		// This loop goes over the mini arrays created by the intervals.
		for i := interval; i < arrLength; i += interval {
			j := i
			// The third loop. This loop goes over the mini arrays and sorts them.
			for j > 0 {
				if numbers[j-interval] > numbers[j] {
					numbers[j-interval], numbers[j] = numbers[j], numbers[j-interval]
				}
				j = j - interval
			}
		}
	}

	return numbers
}

func power(exponent, index int) int {
	power := 1
	for index > 0 {
		if index&1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}

func RandomNumberArrayGenerator(arrayLength int) []int {
	randomArray := make([]int, arrayLength)
	for i := range randomArray {
		randomArray[i] = rand.Intn(100)
	}

	return randomArray
}