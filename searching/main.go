package main

import (
	"fmt"
)

func linearSearch(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}
	return false
}

func linearSortedSearch(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		} else if arr[i] > value {
			break
		}
	}
	return false
}

func main() {
	list := []int{5, 9, 2, 10, 1, 7, 4, 8, 3, 6}
	sortedList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	linearSearchRes := linearSearch(list, 1)
	fmt.Printf("LinearSearch: %v\n", linearSearchRes)

	linearSortedSearchRed := linearSortedSearch(sortedList, 5)
	fmt.Printf("LinearSortedSearch: %v\n", linearSortedSearchRed)

	// binarySearchRes := binarySearch(sortedList, 8)
	// fmt.Printf("\nBinarySearch: %v\n", binarySearchRes)

}
