package main

import (
	"fmt"

	goroutine "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go-Concurrency/1-Goroutines"
	channels "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go-Concurrency/2-Channels"
	buffered_channels "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go-Concurrency/3-Buffered_Channels"
	range_and_close "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go-Concurrency/4-Range_and_Close"
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

	c := make(chan int, 10)
	go range_and_close.Fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

