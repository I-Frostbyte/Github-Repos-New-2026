package channels

func SumOfArrayElements(array []int) int {
	sumChannel := make(chan int)
	
	go halfArraySum(array[:len(array)/2], sumChannel)
	go halfArraySum(array[len(array)/2:], sumChannel)

	rightHalfSum, leftHalfSum := <-sumChannel, <-sumChannel // receive from sumChannel

	return (rightHalfSum + leftHalfSum)
}

func halfArraySum(halfArray []int, sumChannel chan int) {
	sum := 0
	for _, v := range halfArray {
		sum += v
	}
	sumChannel <- sum // send sum to sumChannel
}