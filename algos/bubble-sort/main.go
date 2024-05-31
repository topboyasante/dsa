package main

import "fmt"

// Bubble Sort is an algorithm that sorts an array from the lowest value to the highest value.

// To implement the Bubble Sort algorithm in a programming language, we need:
// An array with values to sort.
// An inner loop that goes through the array and swaps values if the first value is higher than the next value.
// This loop must loop through one less value each time it runs.
// An outer loop that controls how many times the inner loop must run.
// For an array with n values, this outer loop must run n-1 times.

// Time complexity = O(n^2)

func main() {
	var sampleArray = []int{2, 4, 1, 8, 0}
	fmt.Println(BubbleSort(sampleArray))
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		// len(arr)-1 gives you the number of iterations that will occur
		for j := 0; j < len(arr)-i-1; j++ {
			// Every time an iteration occurs, one number is pushed to its right position.
			// In the next iteration, we don't need that number (because its at the right position)
			// len(arr)-i-1 shows us the number of elements needed for the next iteration
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}
