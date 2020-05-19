package main

import "fmt"

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

func main() {
	zeroOneArr := []int{1, 1, 1, 0, 0, 0, 1, 1, 0, 0}
	copied := copySlice(zeroOneArr)
	portionedArr := portionZeroOne(copied)
	fmt.Printf("\nPortionZeroOne:\nInput: %v\nOutput: %v\n", zeroOneArr, portionedArr)

	zeroOneTwoArr := []int{1, 2, 1, 0, 2, 0, 1, 1, 2, 0}
	copiedZeroOneTwoArr := copySlice(zeroOneTwoArr)
	portionedArr = portionZeroOneTwo(copiedZeroOneTwoArr)
	fmt.Printf("\nPortionZeroOneTwo:\nInput: %v\nOutput: %v\n", zeroOneTwoArr, portionedArr)
}
