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
func rearrangeTwoSortedArrays(arr1 []int, arr2[int]) {

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
	fmt.Printf("\nRearrangeTwoSortedArrays:\nOutput:\n%v\n%v", rearrangeTwoSortedArraysInput1, rearrangeTwoSortedArraysInput2)

}
