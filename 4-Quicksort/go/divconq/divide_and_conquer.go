package divconq

import (
	"fmt"
	"time"
)

func NormalSum(arr []int) int {
	start := time.Now()
	total := 0
	for _, v := range arr {
		total += v
	}
	fmt.Printf("Took %v\n", time.Since(start))
	return total
}

func RecursiveSum(arr []int) int {
	start := time.Now()
	if len(arr) == 0 {
		return 0
	}

	fmt.Printf("Took %v\n", time.Since(start))
	return arr[0] + RecursiveSum(arr[1:])
}

func Rec_count_sum(arr []int) int {
	if len(arr) == 0{
		return 0
	}

	return 1 + Rec_count_sum(arr[1:])
}

func RecursiveMax(arr []int) int {
	start := time.Now()
	if len(arr) == 2{
		if arr[0] > arr[1]  {
			return arr[0]
		}
		return arr[1]
	}
	sub_max := RecursiveMax(arr[1:])
	if arr[0] > sub_max {
		return arr[0]
	}
	fmt.Printf("Took %v\n", time.Since(start))
	return sub_max
}