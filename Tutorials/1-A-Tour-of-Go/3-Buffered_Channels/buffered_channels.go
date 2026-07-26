package bufferedchannels

import (
	goroutines "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/1-A-Tour-of-Go/1-Goroutines"
)


func CreateMultipleRandomArrays(arrayLengths []int) [][]int {
	bufferedChannel := make(chan []int, len(arrayLengths))
	for _, arrlen := range arrayLengths {
		newSortedArray := goroutines.CreateAndSortArray(arrlen)
		bufferedChannel <- newSortedArray
	}

	var newArrays [][]int
	for i := 0; i < len(arrayLengths); i++ {
		newArrays = append(newArrays, <-bufferedChannel)
	}
	
	return newArrays
}