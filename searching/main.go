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

func binarySearch(arr []int, value int) bool {
	start, end := 0, len(arr)-1
	for end >= start {
		middle := (start + end) / 2
		if arr[middle] == value {
			return true
		} else if arr[middle] < value {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}

func binarySearchRec(arr []int, value int) bool {
	return binarySearchUtil(arr, 0, len(arr)-1, value)
}

func binarySearchUtil(arr []int, start int, end int, value int) bool {
	if start > end {
		return false
	}
	middle := (start + end) / 2
	if arr[middle] == value {
		return true
	}
	if arr[middle] < value {
		return binarySearchUtil(arr, middle+1, end, value)
	}
	return binarySearchUtil(arr, start, middle-1, value)
}

func main() {
	list := []int{5, 9, 2, 10, 1, 7, 4, 8, 3, 6}
	sortedList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	linearSearchRes := linearSearch(list, 1)
	fmt.Printf("LinearSearch: %v\n", linearSearchRes)

	linearSortedSearchRed := linearSortedSearch(sortedList, 5)
	fmt.Printf("LinearSortedSearch: %v\n", linearSortedSearchRed)

	binarySearchRes := binarySearch(sortedList, 8)
	fmt.Printf("\nBinarySearch: %v\n", binarySearchRes)

	binarySearchRecRes := binarySearchRec(sortedList, 1)
	fmt.Printf("\nBinarySearchRec: %v\n", binarySearchRecRes)

}
