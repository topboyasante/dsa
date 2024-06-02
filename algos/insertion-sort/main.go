package main

import "fmt"

func main() {
	var sampleArray = []int{2, 4, 1, 8, 0}
	fmt.Println(InsertionSort(sampleArray))
}

func InsertionSort(arr []int) []int {
	for i := 1; i <= len(arr)-1; i++ {
		// The current value we are on in our array
		currentValue := arr[i]
		// Holds the "empty" index
		emptyIndex := i

		// This loop will run as long as the "empty" index is greater than 0 and the 
		// element before the "empty" index is greater than the current value we are on in our array.
		// (because we are assuming that the first element in the array is sorted)
		// This assumption is because, if an array has one element, then its already sorted
		for emptyIndex > 0 && arr[emptyIndex-1] > currentValue {
			// this will shift the value before the empty place to the empty place
			// now that place will be empty
			arr[emptyIndex] = arr[emptyIndex-1]
			emptyIndex = emptyIndex - 1
		}
		// put the current value in the empty index
		arr[emptyIndex] = currentValue
	}
	return arr
}
