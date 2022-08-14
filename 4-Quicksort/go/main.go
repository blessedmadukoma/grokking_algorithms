package main

import (
	"fmt"
	"my_quicksort/divconq"
)

func main() {
	// Divide and Conquer implementation
	list := []int{0, 1, 5, 2, 3, 4, 8}
	fmt.Println("Total for normal array:", divconq.NormalSum(list))

	list = []int{1, 5, 2, 3, 4, 8}
	fmt.Println("Total for recursive array:", divconq.RecursiveSum(list))
	fmt.Println("Total no of elements in recursive array:", divconq.Rec_count_sum(list))
	fmt.Println("Maximum number recursively:", divconq.RecursiveMax(list))

	fmt.Println("-------")

	// Quicksort implementation

}
