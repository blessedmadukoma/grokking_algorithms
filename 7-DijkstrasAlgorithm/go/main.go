package main

// import (
// 	"fmt"
// )

// type nodes struct {
// 	name string
// 	value int
// }
// type nodes struct {
// 	name string
// 	value int
// }

// func main() {
// 	fmt.Println("Implementation of Dijkstra's Algorithm in Go")

// 	minigraph := map[string]int{
// 		"a": 6,
// 		"b": 2,
// 	}

// 	// graph hashtable/map
// 	graph := map[string]minigraph{
// 		"a":      ["a"]: 6,
// 		"b":      2,
// 		"finish": 0,
// 	}

// 	// costs hashtable/map
// 	costs := map[string]int{
// 		"a":      6,
// 		"b":      2,
// 		"finish": 0,
// 	}
// 	// parents hashtable/map
// 	parents := map[string]string{
// 		"a":      "start",
// 		"b":      "start",
// 		"finish": "",
// 	}

// 	fmt.Println(graph)
// 	search("me", graph)

// 	// The search only runs 3 times because the key "me" has only 3 elements. To fix this, we need to find a way to check if any key has any element, then append to the queue before searching the queue
// }

// func search(key string, graph map[string][]string) {
// 	search_queue := make([]string, 0) // prevents memory leak
// 	var searched []string

// 	for _, v := range graph[key] {
// 		search_queue = append(search_queue, v)
// 	}

// 	for i := 0; i < len(search_queue); i++ {
// 		person, remains := pop(search_queue)
// 		isContained := Contains(searched, person)

// 		if isContained {
// 			fmt.Printf("`%s` already exists in the array!\n", person)
// 			continue
// 		} else {
// 			fmt.Printf("`%s` is the key\n", person)
// 			name, ok := findPerson(person)
// 			if ok {
// 				fmt.Println(name, "is found!")
// 				return
// 			}

// 			// Else: if the person is not found

// 			// changing the queue so the first element is not there again
// 			search_queue = remains

// 			fmt.Println("New search queue:", search_queue)

// 			// check if the key has values
// 			if len(graph[person]) > 0 {
// 				for _, v := range graph[person] {
// 					// append the values of the key "person" to the updated queue
// 					search_queue = append(search_queue, v)
// 				}
// 			} else {
// 				fmt.Printf("`%s` has no neighbor! \n", person)
// 			}

// 			searched = append(searched, person)

// 			// reset the search_queue after updating the queue
// 			i = -1

// 		}
// 	}
// }

// // pop deletes and returns the first element in the array or slice
// func pop(slice []string) (string, []string) {
// 	first, remains := slice[0], slice[1:]
// 	return first, remains
// }

// // Contains checks if an element is present in a slice
// func Contains(sl []string, name string) bool {
// 	for _, value := range sl {
// 		if &value == &name {
// 			return true
// 		}
// 	}
// 	return false
// }

// // findPerson checks if the element is who we are looking for
// func findPerson(name string) (string, bool) {
// 	if name == "Daniel" {
// 		return name, true
// 	}
// 	return "", false
// }
