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

/*
printDuplicatesInAList
Given an array of n numbers, print the duplicate elements in the array
*/
func printDuplicates(arr []int) {
	hash := map[int]bool{}
	for _, num := range arr {
		if hash[num] {
			fmt.Println(num)
		} else {
			hash[num] = true
		}
	}
}

/*
RemoveDuplicates
Given an array of n numbers, remove the duplicate elements in the array
*/
func removeDuplicates(arr []int) []int {
	hash := map[int]bool{}
	res := []int{}
	for _, num := range arr {
		if !hash[num] {
			res = append(res, num)
			hash[num] = true
		}
	}
	return res
}

func main() {
	firstRepeatedInput := []int{7, 1, 6, 3, 5, 1, 7, 4, 2}
	firstRepeatedRes := firstRepeated(firstRepeatedInput)
	fmt.Printf("\nFirstRepeated:\nInput: %v\nOutput: %v\n", firstRepeatedInput, firstRepeatedRes)

	printDuplicates(firstRepeatedInput)

	removeDuplicatesRes := removeDuplicates(firstRepeatedInput)
	fmt.Printf("\nRemoveDuplicates:\nInput: %v\nOutput: %v\n", firstRepeatedInput, removeDuplicatesRes)
}