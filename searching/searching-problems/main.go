package main

import "fmt"

/*
firstRepeatedElement
Given an unsorted list of n elements, find the first element which is repeated
*/
func firstRepeated(arr []int) int {
	hash := map[int]bool{}
	for _, num := range arr {
		if hash[num] {
			return num
		}
		hash[num] = true
	}
	return 0
}

func main() {
	firstRepeatedInput := []int{7, 1, 6, 3, 5, 1, 7, 4, 2}
	firstRepeatedRes := firstRepeated(firstRepeatedInput)
	fmt.Printf("\nFirstRepeated:\nInput: %v\nOutput: %v\n", firstRepeatedInput, firstRepeatedRes)
}
