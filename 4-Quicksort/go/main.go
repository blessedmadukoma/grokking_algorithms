package main

import (
	"fmt"
	"my_quicksort/divconq"
	"my_quicksort/qs_imp"
)

func main() {
	// Divide and Conquer implementation
	list := []int{0, 1, 5, 2, 3, 4, 8}
	fmt.Println("Total for normal array before Divide and Conquer:", divconq.NormalSum(list))

	// list = []int{1, 5, 2, 3, 4, 8}
	fmt.Println("Total for recursive array for divide and conquer:", divconq.RecursiveSum(list))
	fmt.Println("Total no of elements in recursive array for divide and conquer:", divconq.Rec_count_sum(list))
	fmt.Println("Maximum number recursively for divide and conquer:", divconq.RecursiveMax(list))

	fmt.Println("-------")

	// Quicksort implementation
	list = []int{1, 10, 5, 2, 3, 4, 8}
	fmt.Println("Quicksort in action:", qs_imp.QuickSort(list, 0, 6))
	fmt.Println("Quicksort in action:", qs_imp.QuickSortMine(list))
}
