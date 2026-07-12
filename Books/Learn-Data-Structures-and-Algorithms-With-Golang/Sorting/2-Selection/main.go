package main

import "fmt"

func main () {
	elements := []int{11, 4, 8, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before Sorting ", elements)
	sortedElements := SelectionSort(elements)
	fmt.Println("After Sorting", sortedElements)
}

func SelectionSort(elements []int) []int {
	var i int
	for i = 0; i < len(elements)-1; i++  {
		// initialize the minimum element to be the one in position i.
		min := i
		var j int
		for j = i + 1; j <= len(elements)-1; j++ {
			// compare the element in position j with the minimum element
			if elements[j] < elements[min] {
				min = j
			}
		}
		// Send the elements to the swap function to change their positions
		swap(elements, i, min)
	}

	return elements
}

// swap method
func swap(elements []int, i , min int) {
	var temp int
	temp = elements[min]
	elements[min] = elements[i]
	elements[i] = temp
}