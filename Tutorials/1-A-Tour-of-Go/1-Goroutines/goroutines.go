package goroutines

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
)

func RandomArrayGenerator(arrayLength int) []int {
	newArray := make([]int, arrayLength)

	for index := range newArray {
		newArray[index] = rand.Intn(75)
	}

	return newArray
}

func CreateAndSortArray(arrayLength int) {
	var newRandomArray []int = RandomArrayGenerator(12)

	fmt.Println("Created Array: ", newRandomArray)

	// Initiate WaitGroup Counter.
	// WaitGroup counters wait for a group of goroutines or tasks to finish.
	wg := &sync.WaitGroup{}

	var sortedArray []int

	// Adding a +ve count to the WaitGroup
	wg.Add(1)
	go func(array []int, sortedArray *[]int) {
		// Decrementing the count in the WaitGroup by adding -1.
		defer wg.Done()
		*sortedArray = slices.Sorted(slices.Values(array))
	}(newRandomArray, &sortedArray)

	// Wait blocks until the WaitGroup task counter is zero.
	wg.Wait()

	fmt.Println("Sorted Array: ", sortedArray)

	// return newRandomArray
}