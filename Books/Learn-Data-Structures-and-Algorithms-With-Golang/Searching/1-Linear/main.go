package main

func main () {

}

func LinearSearch(numberArray []int, findElement int) bool {
	for _, element := range numberArray {
		if element == findElement {
			return true
		}
	}

	return false
}