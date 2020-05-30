package main

import (
	"fmt"
	"math"
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

func quickSort(list []int) []int {
	size := len(list)
	quickSortUtil(&list, 0, size-1)
	return list
}

func quickSortUtil(list *[]int, lower int, upper int) {
	if upper <= lower {
		return
	}

	pivot, start, stop := (*list)[lower], lower, upper
	for lower <= upper {
		for (*list)[lower] <= pivot {
			lower++
		}
		for (*list)[upper] > pivot {
			upper--
		}
		if lower < upper {
			swap(list, lower, upper)
		}
	}
	swap(list, start, upper)

	quickSortUtil(list, start, upper-1)
	quickSortUtil(list, upper+1, stop)
}

func bucketSort(list []int) []int {
	var greatest int
	for _, num := range list {
		greatest = int(math.Max(float64(greatest), float64(num)))
	}
	buckets := make([]int, greatest+1)
	for _, num := range list {
		buckets[num]++
	}
	currentIdx := 0
	for i, count := range buckets {
		for count > 0 {
			list[currentIdx] = i
			count--
			currentIdx++
		}
	}
	return list
}

func heapSort(list []int) []int {
	size := len(list)
	if size < 2 {
		return list
	}
	buildMaxHeap(list)
	for i := size - 1; i > 0; i-- {
		swap(&list, 0, size-1)
		size--
		heapify(list, 0, size)
	}
	return list
}

func buildMaxHeap(list []int) {
	size := len(list)
	middle := size / 2
	for i := middle; i >= 0; i-- {
		heapify(list, i, size)
	}
}

func heapify(list []int, idx int, size int) {
	// find left and right nodes
	left, right := (idx*2)+1, (idx*2)+2
	max := idx

	if left < size && list[left] > list[max] {
		max = left
	}
	if right < size && list[right] > list[max] {
		max = right
	}
	if max != idx {
		swap(&list, idx, max)
		heapify(list, max, size)
	}
}

func main() {
	list := []int{5, 9, 2, 10, 1, 7, 4, 8, 3, 6}

	bubbleSortedList := bubbleSort(list)
	fmt.Printf("BubbleSort: %v\n", bubbleSortedList)

	insertionSortedList := insertionSort(list)
	fmt.Printf("InsertionSort: %v\n", insertionSortedList)

	selectionSortedList := selectionSort(list)
	fmt.Printf("SelectionSort: %v\n", selectionSortedList)

	mergeSortedList := mergeSort(list)
	fmt.Printf("MergeSort: %v\n", mergeSortedList)

	quickSortedList := quickSort(list)
	fmt.Printf("QuickSort: %v\n", quickSortedList)

	bucketSortedList := bucketSort(list)
	fmt.Printf("BucketSort: %v\n", bucketSortedList)

	heapSortedList := heapSort(list)
	fmt.Printf("HeapSort: %v\n", heapSortedList)
}
