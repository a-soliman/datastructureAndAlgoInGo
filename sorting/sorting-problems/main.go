package main

import (
	"fmt"
	"math"
)

func copySlice(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	return copied
}

func swap(arr []int, idx1 int, idx2 int) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}

/*
Portion zero one, Given an array containing 0s and 1s. Write an algorithm to sort array so that 0s come first followed by 1s.
*/
func portionZeroOne(arr []int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		for arr[left] == 0 {
			left++
		}
		for arr[right] == 1 {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return arr
}

/*
Portion zero, one, and two
Given an array containing 0s, 1s, and 2s. Write an algorithm to sort the array so that 0s come first, followed by 1s and then 2s in the end.
*/
func portionZeroOneTwo(arr []int) []int {
	left, right := 0, len(arr)-1
	for i := 0; i < len(arr); i++ {
		for arr[i] == 0 && left < i {
			swap(arr, i, left)
			left++
		}
		for arr[i] == 2 && right > i {
			swap(arr, i, right)
			right--
		}
	}
	return arr
}

/*
Range portion
Given an array of integers and a range. Write an algorithm to portion array so that values smaller than the range come to left,
then values under the range followed by values greater than the range
*/
func rangePortion(arr []int, lower int, higher int) []int {
	start, end := 0, len(arr)-1

	for i := 0; i < len(arr); i++ {
		if arr[i] < lower && i > start {
			swap(arr, start, i)
			start++
		} else if arr[i] > higher && i < end {
			swap(arr, end, i)
			end--
		}
	}

	return arr
}

/*
MinimumSwaps
Given an array of integers and a value intiger. Find the minimum swaps required to bring all the values less than the given value to the start of the array.
*/

func minSwaps(arr []int, value int) int {
	swaps := 0
	left, right := 0, len(arr)-1
	for right > left {
		for arr[left] < value {
			left++
		}
		for arr[right] > value {
			right--
		}
		if left < right {
			swaps++
			left++
			right--
		}
	}
	return swaps
}

/*
AbsoluteSort
Given an array of integers and a value intiger. Sort the array according to the absolute difference from the given value.
*/
func absoluteSort(arr []int, value int) []int {
	idx := 0
	for i := 0; i < len(arr); i++ {
		idx = i
		for idx >= 1 && isAbsLess(arr[idx], arr[idx-1], value) {
			swap(arr, idx, idx-1)
			idx--
		}
	}
	return arr
}

func isAbsLess(val1 int, val2 int, ref int) bool {
	return math.Abs(float64(ref-val1)) < math.Abs(float64(ref-val2))
}

/*
SortByOrder
Given two arrays. Sort the first array according to the order defined in the second array
*/

func sortByOrder(arr []int, order []int) []int {
	if len(arr) < 2 || len(order) < 1 {
		return arr
	}
	count := map[int]int{}
	for _, v := range arr {
		count[v]++
	}
	currentIdx := 0
	// loop over the order slice and construct the arr
	for _, currentOrder := range order {
		_, ok := count[currentOrder]
		if ok {
			for count[currentOrder] > 0 {
				arr[currentIdx] = currentOrder
				currentIdx++
				count[currentOrder]--
			}
			if count[currentOrder] == 0 {
				delete(count, currentOrder)
			}
		}
	}
	// append the rest of the map to the arr
	for key := range count {
		for count[key] > 0 {
			arr[currentIdx] = key
			currentIdx++
			count[key]--
		}
		delete(count, key)
	}
	return arr
}

/*
SortEvenOdd
Given an int array containing even and odd values, separate even numbers from the odd numbers
*/
func sortEvenOdd(arr []int) []int {
	evenEnd := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			swap(arr, i, evenEnd)
			evenEnd++
		}
	}
	return arr
}

/*
RearrangeTwoSortedArrays
Given two sorted int arrays, rearrange them so that the largest element in first array is smaller that the first element of the 2nd array.
both arrays should still be sorted
*/
func rearrangeTwoSortedArrays(arr1 []int, arr2 []int) {
	if len(arr1) == 0 || len(arr2) == 0 {
		return
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[0] {
			arr1[i], arr2[0] = arr2[0], arr1[i]
			bubbleUp(arr2)
		}
	}
	return
}

func bubbleUp(arr []int) {
	temp := arr[0]
	i := 1

	for i < len(arr) && arr[i] < temp {
		arr[i-1] = arr[i]
		i++
	}
	arr[i-1] = temp
}

/*
CheckReversed
Given an Array of ints, check if reversing a subarray will sort the original array
*/
func checkRevered(arr []int) bool {
	var (
		rangeStart int
		rangeEnd   int
	)
	// find rangeStart
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			rangeStart = i
			rangeEnd = i
			break
		}
	}
	// find rangeEnd
	for i := rangeStart + 1; i < len(arr); i++ {
		if arr[i] < arr[rangeEnd] {
			rangeEnd = i
		}
	}
	if !isSorted(arr[0:rangeStart]) || arr[rangeEnd] < arr[rangeStart-1] {
		return false
	}
	if !isSorted(arr[rangeEnd+1:]) || arr[rangeStart] > arr[rangeEnd+1] {
		return false
	}
	if !isReversedSorted(arr[rangeStart : rangeEnd+1]) {
		return false
	}
	return true
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func isReversedSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

/*
Merge K sorted arrays
This is a popular Facebook problem.

Given K sorted arrays arr, of size N each, merge them into a new array res, such that res is a sorted array.

 Assume N is very large compared to K. N may not even be known. The arrays could be just sorted streams, for instance, timestamp streams.

All arrays might be sorted in increasing manner or decreasing manner. Sort all of them in the manner they appear in input.

Repeats are allowed.
Negative numbers and zeros are allowed.
Assume all arrays are sorted in the same order. Preserve that sort order in output.
It is possible to find out the sort order from at least one of the arrays.
*/
func mergeKArrays(arr [][]int32) []int32 {
	res := []int32{}
	arraysCount := len(arr)
	increasing := false

	if arraysCount == 1 {
		return arr[0]
	}

	if arraysCount >= 2 {
		increasing = findOrder(arr)
		res = merge2SortedArrays(arr[0], arr[1], increasing)
	}

	for i := 2; i < arraysCount; i++ {
		res = merge2SortedArrays(res, arr[i], increasing)
	}
	return res
}

func findOrder(arr [][]int32) bool {
	res := true
	currentArr := []int32{}
	for i := 0; i < len(arr); i++ {
		currentArr = arr[i]
		if len(currentArr) >= 2 {
			for j := 1; j < len(currentArr); j++ {
				if currentArr[j] != currentArr[j-1] {
					return currentArr[j-1] < currentArr[j]
				}
			}
		}
	}
	return res
}

func merge2SortedArrays(arr1 []int32, arr2 []int32, increasing bool) []int32 {
	res := []int32{}
	size1, size2 := len(arr1), len(arr2)
	i, j := 0, 0

	for i < size1 && j < size2 {
		if increasing {
			if arr1[i] < arr2[j] {
				res = append(res, arr1[i])
				i++
			} else {
				res = append(res, arr2[j])
				j++
			}
		} else {
			if arr1[i] > arr2[j] {
				res = append(res, arr1[i])
				i++
			} else {
				res = append(res, arr2[j])
				j++
			}
		}
	}
	for i < size1 {
		res = append(res, arr1[i])
		i++
	}
	for j < size2 {
		res = append(res, arr2[j])
		j++
	}
	return res
}

func main() {
	zeroOneArr := []int{1, 1, 1, 0, 0, 0, 1, 1, 0, 0}
	copied := copySlice(zeroOneArr)
	portionedArr := portionZeroOne(copied)
	fmt.Printf("\nPortionZeroOne:\nInput: %v\nOutput: %v\n", zeroOneArr, portionedArr)

	zeroOneTwoArr := []int{1, 2, 1, 0, 2, 0, 1, 1, 2, 0}
	copiedZeroOneTwoArr := copySlice(zeroOneTwoArr)
	portionedArr = portionZeroOneTwo(copiedZeroOneTwoArr)
	fmt.Printf("\nPortionZeroOneTwo:\nInput: %v\nOutput: %v\n", zeroOneTwoArr, portionedArr)

	rangePortionInput := []int{7, 1, 6, 3, 5, 4, 2}
	copiedRangePortionInput := copySlice(rangePortionInput)
	portionedArr = rangePortion(copiedRangePortionInput, 3, 5)
	fmt.Printf("\nRangePortion:\nInput: %v\nOutput: %v\n", rangePortionInput, portionedArr)

	minSwapsInput := []int{2, 7, 5, 6, 1, 3, 4, 9, 10, 8}
	value := 5
	minSwapsOutput := minSwaps(minSwapsInput, value)
	fmt.Printf("\nMinSwaps:\nInput: %v, %d\nOutput: %d\n", minSwapsInput, value, minSwapsOutput)

	absoluteSortInput := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	copiedAbsoluteSortInput := copySlice(absoluteSortInput)
	value = 5
	absoluteSortOutput := absoluteSort(copiedAbsoluteSortInput, value)
	fmt.Printf("\nAbsoluteSort:\nInput: %v, %d\nOutput: %d\n", copiedAbsoluteSortInput, value, absoluteSortOutput)

	sortByOrderInput := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8}
	order := []int{2, 1, 8, 3}
	copiedSortByOrderInput := copySlice(sortByOrderInput)
	sortByOrderOutput := sortByOrder(copiedSortByOrderInput, order)
	fmt.Printf("\nSortByOrder:\nInput: %v,\nOrder: %v\nOutput: %d\n", sortByOrderInput, order, sortByOrderOutput)

	evenOddInput := []int{2, 7, 5, 6, 1, 3, 4, 9, 10, 8}
	copiedEvenOddInput := copySlice(evenOddInput)
	evenOddOutput := sortEvenOdd(copiedEvenOddInput)
	fmt.Printf("\nSortEvenOdd:\nInput: %v\nOutput: %v\n", evenOddInput, evenOddOutput)

	rearrangeTwoSortedArraysInput1 := []int{1, 5, 9, 10, 15, 20}
	rearrangeTwoSortedArraysInput2 := []int{2, 3, 8, 13}
	rearrangeTwoSortedArrays(rearrangeTwoSortedArraysInput1, rearrangeTwoSortedArraysInput2)
	fmt.Printf("\nRearrangeTwoSortedArrays:\nOutput:\n%v\n%v\n", rearrangeTwoSortedArraysInput1, rearrangeTwoSortedArraysInput2)

	checkReveredInput := []int{1, 3, 8, 5, 4, 3, 10, 11, 12, 18, 28}
	checkReveredOutput := checkRevered(checkReveredInput)
	fmt.Printf("\nCheckRevered:\nInput: %v\nOutput: %v\n", checkReveredInput, checkReveredOutput)

	mergeKArraysInput := [][]int32{}

	row1 := []int32{3, 3, 12, 20, 22, 25, 34}
	row2 := []int32{4, 10, 12, 20, 28, 32, 36}
	row3 := []int32{5, 6, 10, 19, 22, 28, 34}
	row4 := []int32{8, 17, 17, 25, 34, 34, 42}
	row5 := []int32{6, 7, 14, 17, 18, 25, 26}
	row6 := []int32{8, 10, 15, 19, 28, 32, 40}
	row7 := []int32{5, 13, 17, 19, 25, 26, 27}
	row8 := []int32{1, 9, 12, 20, 26, 28, 30}
	row9 := []int32{0, 8, 13, 19, 21, 25, 28}
	row10 := []int32{3, 12, 18, 21, 27, 32, 32}

	mergeKArraysInput = append(mergeKArraysInput, row1, row2, row3, row4, row5, row6, row7, row8, row9, row10)
	mergeKArraysOutput := mergeKArrays(mergeKArraysInput)
	fmt.Printf("\nMergeKSortedArrays: \nOutput: %v", mergeKArraysOutput)

}
