package main

import "fmt"

// The Selection Sort algorithm finds the lowest value in an array and moves it to the front of the array.

// To implement the Selection Sort algorithm in a programming language, we need:
// An array with values to sort.
// An inner loop that goes through the array, finds the lowest value, and moves it to the front of the array.
// This loop must loop through one less value each time it runs.
// An outer loop that controls how many times the inner loop must run.
// For an array with n values, this outer loop must run  n - 1 times.

func main() {
	var sampleArray = []int{2, 4, 1, 8, 0}
	fmt.Println(SelectionSort(sampleArray))
}

func SelectionSort(arr []int) []int {
	// len(arr)-1 gives you the number of iterations that will occur
	// Outer loop loops through each iteration (moving from the beginning to the end of the array)
	for i := 0; i < len(arr)-1; i++ {
		// This variable holds the index of the lowest number.
		// On the first loop iteration, the value is 0, and increases when the iteration is complete
		min_index := i
		// fmt.Printf("When i=%v, min_index=%v\n", i, min_index)

		// Inner loop loops through each element
		for j := i; j < len(arr); j++ {
			fmt.Printf("When i=%v, min_index=%v, and j=%v\n", i, min_index, j)
			// fmt.Printf("When arr[%v]=%v, and arr[%v]=%v\n", j, arr[j], min_index, arr[min_index])
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		arr[i], arr[min_index] = arr[min_index], arr[i]
		fmt.Println(arr)
	}
	return arr
}
