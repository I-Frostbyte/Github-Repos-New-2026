package main

import (
	"fmt"
	"math/rand"
)

func main () {
	newArray := []int{11, 4, 8, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Array before Sorting: ", newArray)
	sortedArray := QuickSort(newArray, 0, 9)
	fmt.Println("Array after Sorting: ", sortedArray)
}

func QuickSort(numbersArray []int, below, upper int) []int {
	if below < upper {
		part := divideParts(numbersArray, below, upper)
		QuickSort(numbersArray, below, part-1)
		QuickSort(numbersArray, part+1, upper)
	}

	return numbersArray
}

func divideParts(sliceArray []int, below, upper int) int {
	center := sliceArray[upper]
	i := below

	for j := below; j < upper; j++ {
		if sliceArray[j] <= center {
			swap(&sliceArray[i], &sliceArray[j])
			i += 1
		}
	}

	swap(&sliceArray[i], &sliceArray[upper])
	return i
}

func swap (sliceArrayPointerOne, sliceArrayPointerTwo *int) {
	val := *sliceArrayPointerOne
	*sliceArrayPointerOne = *sliceArrayPointerTwo
	*sliceArrayPointerTwo = val
}

func RandomlyGeneratedArray(arrayLength int) []int {
	newArray := make([]int, arrayLength)

	for i := range newArray {
		newArray[i] = rand.Intn(100)
	}

	return newArray
}