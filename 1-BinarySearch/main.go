package main

import "fmt"

func main() {
	// This function is to implement the binary search
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index, value := binary_search(list, 3)
	fmt.Printf("Index: %d => Value: %d\n", index, value)
}

func binary_search(array []int, value int) (int, int) {
	low := 0
	high := len(array) - 1

	for low <= high {
		midpoint := (low + high)
		guess := array[midpoint]
		if guess == value {
			return midpoint, value
		} else if guess > value {
			high = midpoint - 1
		} else {
			low = midpoint + 1
		}
	}
	return 0, 0
}
