package qs_imp

import (
	"fmt"
	"time"
)

func QuickSortMine(array []int) []int {
	start := time.Now()
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		less, greater := []int{}, []int{}
		for _, v := range array[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}

		sortedArray := QuickSortMine(less)
		sortedArray = append(sortedArray, pivot)
		for _, v := range QuickSortMine(greater) {
			sortedArray = append(sortedArray, v)
		}

		fmt.Println("--Mine:", time.Since(start))
		return sortedArray
	}
}

func QuickSort(arr []int, low, high int) []int {
	start := time.Now()
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}

	fmt.Println("--Online:", time.Since(start))
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
