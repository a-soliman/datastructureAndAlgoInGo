package main

import "fmt"

func linearSearch(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}
	return false
}

func main() {
	list := []int{5, 9, 2, 10, 1, 7, 4, 8, 3, 6}

	linearSearchRes := linearSearch(list, 1)
	fmt.Printf("LinearSearch: %v\n", linearSearchRes)
}
