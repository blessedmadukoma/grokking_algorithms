package main

import "fmt"

func main() {
	// This function is to implement the selection sort
	list := []int{6, 2, 5, 3, 8, 1}

	fmt.Println("Old array:", list)

	newArr := selectionSort(list)
	fmt.Println("New array:", newArr)
}

func findSmallest(arr []int) (int, int) {
	smallest := arr[0]
	smallest_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return smallest_index, smallest
}

func selectionSort(array []int) []int {
	var newArr []int

	for range array {
		smallest_index, value := findSmallest(array)
		newArr = append(newArr, value)
		array = removeIndex(array, smallest_index)
	}
	return newArr
}

func removeIndex(slice []int, s int) []int {
	copy(slice[s:], slice[s+1:])
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}
