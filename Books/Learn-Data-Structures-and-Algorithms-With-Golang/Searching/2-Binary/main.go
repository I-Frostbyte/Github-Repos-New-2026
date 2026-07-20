package main

import (
	"sort"
)

func main() {}

// BinarySearch utilizes the sort.Search algorithm.
// sort.Search searches an orderd array and returns
// the index where the number was found or where the number
// would be found if it were to exist.
// BinarySearch builds on that by using that index to compare
// the value at the index with the element to be found.
// It then returns a true or false based on the response.
func BinarySearch(inputArray []int, findElement int) bool {
	found := false

	index := sort.Search(len(inputArray), func(i int) bool {
		return inputArray[i] >= findElement
	})

	if index < len(inputArray) && inputArray[index] == findElement {
		found = true
	}

	return found
}