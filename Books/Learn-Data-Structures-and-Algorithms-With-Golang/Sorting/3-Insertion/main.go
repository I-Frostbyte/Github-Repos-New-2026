package main

import (
	"fmt"
	"math/rand"
)

func main () {
	randomArray := RandomNumberArrayGenerator(12)
	fmt.Println("Randomly Generated Array: ", randomArray)
	sortedArray := InsertionSort(randomArray)
	fmt.Println("Sorted Array: ", sortedArray)
}

func RandomNumberArrayGenerator(arrayLength int) []int {
	generatedArray := make([]int, arrayLength)
	for i := range arrayLength {
		generatedArray[i] = rand.Intn(150)
	}

	return generatedArray
}

func InsertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		j := i
		for j > 0 {
			// compares every number at this index with every index beneath it and swapping them where necessary.
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
			j = j-1
		}
	}

	return numbers
}