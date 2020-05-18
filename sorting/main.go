package main

import (
	"fmt"
)

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
	for i := 1; i < size; i++ {
		for idx := i; idx > 0; idx-- {
			if idx > 0 && less(list[idx], list[idx-1]) {
				swap(&list, idx, idx-1)
			}
		}
	}
	return list
}

func selectionSort(list []int) []int {
	var largestVal int
	var largestIdx int
	unsortedSize := len(list) - 1

	for {
		if unsortedSize == 0 {
			break
		}
		largestVal = list[0]
		largestIdx = 0
		for i := 1; i <= unsortedSize; i++ {
			if more(list[i], largestVal) {
				largestVal = list[i]
				largestIdx = i
			}
		}
		if largestIdx != unsortedSize {
			swap(&list, largestIdx, unsortedSize)
		}
		unsortedSize--
	}
	return list
}

func mergeSort(list []int) []int {
	size := len(list)
	if size == 1 {
		return list
	}
	mid := size / 2
	left := mergeSort(list[0:mid])
	right := mergeSort(list[mid:size])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	leftSize, rightSize := len(left), len(right)
	leftStart, rightStart := 0, 0
	var sortedList []int

	for leftStart < leftSize && rightStart < rightSize {
		if left[leftStart] < right[rightStart] {
			sortedList = append(sortedList, left[leftStart])
			leftStart++
		} else {
			sortedList = append(sortedList, right[rightStart])
			rightStart++
		}
	}
	for leftStart < leftSize {
		sortedList = append(sortedList, left[leftStart])
		leftStart++
	}
	for rightStart < rightSize {
		sortedList = append(sortedList, right[rightStart])
		rightStart++
	}
	return sortedList
}

func main() {
	list := []int{5, 9, 4, 10, 1, 7, 2, 8, 3, 6}

	bubbleSortedList := bubbleSort(list)
	fmt.Printf("BubbleSort: %v\n", bubbleSortedList)

	insertionSortedList := insertionSort(list)
	fmt.Printf("InsertionSort: %v\n", insertionSortedList)

	selectionSortedList := selectionSort(list)
	fmt.Printf("SelectionSort: %v\n", selectionSortedList)

	mergeSortedList := mergeSort(list)
	fmt.Printf("MergeSort: %v\n", mergeSortedList)
}
