package main

import (
	"fmt"
	"time"
)

func main() {
	// list := []int{0, 1, 5, 2, 3, 4, 8}

	// fmt.Println("Total for normal array:", normalSum(list))

	list := []int{1, 5, 2, 3, 4, 8}
	// fmt.Println("Total for recursive array:", recursiveSum(list))
	// fmt.Println("Total no of elements in recursive array:", rec_count_sum(list))
	fmt.Println("Maximum number recursively:", recursiveMax(list))
}

func normalSum(arr []int) int {
	start := time.Now()
	total := 0
	for _, v := range arr {
		total += v
	}
	fmt.Printf("Took %v\n", time.Since(start))
	return total
}

func recursiveSum(arr []int) int {
	start := time.Now()
	if len(arr) == 0 {
		return 0
	}

	fmt.Printf("Took %v\n", time.Since(start))
	return arr[0] + recursiveSum(arr[1:])
}

func rec_count_sum(arr []int) int {
	if len(arr) == 0{
		return 0
	}

	return 1 + rec_count_sum(arr[1:])
}

func recursiveMax(arr []int) int {
	start := time.Now()
	if len(arr) == 2{
		if arr[0] > arr[1]  {
			return arr[0]
		}
		return arr[1]
	}
	sub_max := recursiveMax(arr[1:])
	if arr[0] > sub_max {
		return arr[0]
	}
	fmt.Printf("Took %v\n", time.Since(start))
	return sub_max
}