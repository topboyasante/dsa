package main

import "fmt"

// An array is a collection of elements of the same type

func main() {
	// We have an array that has 5 integers, with values 2,4,6,8,100
	var sampleArray = [5]int{2, 4, 1, 8, 0}

	// We can change the elements in an array
	sampleArray[0] = 20
	sampleArray[1] = -5

	// We can't add more elements to an array, we would have to create a new array
	// and copy the previous elements into the new one

	minVal := sampleArray[0]

	for _, v := range sampleArray {
		if v < minVal {
			minVal = v
		}
	}

	fmt.Println(minVal)

}
