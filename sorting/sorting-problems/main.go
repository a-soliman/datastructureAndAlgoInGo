package main

import "fmt"

func copySlice(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	return copied
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

func main() {
	zeroOneArr := []int{1, 1, 1, 0, 0, 0, 1, 1, 0, 0}
	copied := copySlice(zeroOneArr)
	portionedArr := portionZeroOne(copied)
	fmt.Printf("PortionZeroOne:\nInput: %v\nOutput: %v\n", zeroOneArr, portionedArr)
}
