package main

import (
	"fmt"

	goroutine "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go/1-Goroutines"
	channels "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go/2-Channels"
	buffered_channels "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go/3-Buffered_Channels"
)

func main () {
	newSortedArray := goroutine.CreateAndSortArray(12)

	sumOfSortedArray := channels.SumOfArrayElements(newSortedArray)

	fmt.Println("Sum of Sorted Array: ", sumOfSortedArray)

	newMultipleArrays := buffered_channels.CreateMultipleRandomArrays([]int{12, 13, 14, 15})

	fmt.Println("New Arrays")
	for _, array := range newMultipleArrays {
		fmt.Println(array)
	}
}

