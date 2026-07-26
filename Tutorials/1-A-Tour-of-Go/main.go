package main

import (
	"fmt"

	goroutine "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go/1-Goroutines"
	channels "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go/2-Channels"
)

func main () {
	newSortedArray := goroutine.CreateAndSortArray(12)

	sumOfSortedArray := channels.SumOfArrayElements(newSortedArray)

	fmt.Println("Sum of Sorted Array: ", sumOfSortedArray)
}

