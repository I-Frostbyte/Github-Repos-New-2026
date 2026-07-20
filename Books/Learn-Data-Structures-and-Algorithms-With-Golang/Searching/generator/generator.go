package generator

import (
	"fmt"
	"math/rand"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{

	}
}

// RandomIntegerArrayGenerator takes as argument the length of the new array,
// a target number and a boolean called checkForTarget.
// If checkForTarget is true, while creating the array, it would check
// if the target number was created.
// If checkForTarget is false, then the newArray is generated.
// The array generates numbers from 0 - 30.
// Duplicates can occur.
func (g *Generator) RandomIntegerArrayGenerator(arrayLength, target int, checkForTarget bool) (*[]int, bool, error) {
	if arrayLength < 2 {
		return nil, false, fmt.Errorf("Array Length to minimal to create new array: %v", arrayLength)
	}
	
	newArray := make([]int, arrayLength)

	found := false
	
	for index := range newArray {
		newArray[index] = rand.Intn(30)
		if checkForTarget {
			if newArray[index] == target {
				found = true
			}
		}
	}

	return &newArray, found, nil
}