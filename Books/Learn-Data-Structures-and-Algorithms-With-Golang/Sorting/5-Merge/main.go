package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newArray := RandomlyGeneratedArray(32)
	fmt.Println("Array Before Merge Sort: ", newArray)
	sortedArray := MergeSort(newArray)
	fmt.Println("Array After Merge Sort: ", sortedArray)
}

func MergeSort(arrayToSort []int) []int {
	if len(arrayToSort) < 2 {
		return arrayToSort
	}

	middle := (len(arrayToSort)) / 2

	return JoinArrays(MergeSort(arrayToSort[:middle]), MergeSort(arrayToSort[middle:]))
}

func JoinArrays(leftArray, rightArray []int) []int {
	lengthOfJoinedArray, i, j := len(leftArray)+len(rightArray), 0, 0

	joinedArray := make([]int, lengthOfJoinedArray)

	for k := 0; k < lengthOfJoinedArray; k++ {
		if i > len(leftArray)-1 && j <= len(rightArray)-1 {
			joinedArray[k] = rightArray[j]
			j++
		} else if j > len(rightArray)-1 && i <= len(leftArray)-1 {
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

func RandomlyGeneratedArray(arrayLength int) []int {
	newArray := make([]int, arrayLength)

	for i := range newArray {
		newArray[i] = rand.Intn(150)
	}

	return newArray
}

/*

                       [11 4 8 6 19 21 71 13 15 2]
                      /                           \
           [11 4 8 6 19]                     [21 71 13 15 2]
           /          \                      /             \
      [11 4]      [8 6 19]            [21 71]        [13 15 2]
      /   \         /    \            /    \          /      \
   [11] [4]      [8]   [6 19]      [21] [71]      [13]   [15 2]
                        /  \                              /   \
                      [6] [19]                         [15]  [2]

*/
