package main

import "fmt"

func less(val1 int, val2 int) bool {
	return val1 < val2
}

func more(val1 int, val2 int) bool {
	return val1 > val2
}

func swap(list *[]int, idx1 int, idx2 int) {
	var temp int
	temp = (*list)[idx1]
	(*list)[idx1] = (*list)[idx2]
	(*list)[idx2] = temp
}

func bubbleSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if more(list[i], list[j]) {
				swap(&list, i, j)
			}
		}
	}
	return list
}

func insertionSort(list []int) []int {
	size := len(list)
	for i := 1; i < size-1; i++ {
		for idx := i; idx >= 0 && less(list[idx], list[idx-1]); idx-- {
			swap(&list, idx, idx-1)
		}
	}
	return list
}

func main() {
	list := []int{5, 9, 4, 10, 1, 7, 2, 8, 3, 6}
	bubbleSortedList := bubbleSort(list)
	insertionSortedList := insertionSort(list)
	fmt.Printf("BubbleSort: %v\n", bubbleSortedList)
	fmt.Printf("InsertionSort: %v\n", insertionSortedList)
}
