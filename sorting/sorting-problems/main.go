package main

import "fmt"

/*
Portion zero one, Given an array containing 0s and 1s. Write an algorithm to sort array so that 0s come first followed by 1s.
*/
func portionZeroOne(arr []int) []int {
	// add your code here
}

func main() {
	zeroOneArr := []int{1, 1, 1, 0, 0, 0, 1, 1, 0, 0}
	copied := make([]int, len(zeroOneArr))
	copy(copied, zeroOneArr)
	portionedArr := portionZeroOne(copied)
	fmt.Printf("PortionZeroOne:\nInput: %v\nOutput: %v\n", zeroOneArr, portionedArr)
}
