package main

import "fmt"

func main() {
	hashTables()
}

func hashTables() {
	book := map[string]float32{
		"apple": 1.04,
		"milk": 1.49,
		"avocado": 1.49,
	}

	fmt.Println(book)
	fmt.Println(book["apple"])
}