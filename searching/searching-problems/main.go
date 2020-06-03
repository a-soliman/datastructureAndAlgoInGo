package main

import (
	"fmt"
	"math"
	"sort"
)

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

/*
FindMissingNumber
In given list of n-1 elements, which are the in the range of 1 to n. There are no duplicates in the array. One of the integers is missing. Find the missing element
*/
func findMissingNumber(arr []int) int {
	n := arr[0]
	var total int
	var actualTotal int

	for _, num := range arr {
		actualTotal += num
		if num > n {
			n = num
		}
	}

	for i := 1; i <= n; i++ {
		total += i
	}

	return total - actualTotal
}

/*
MissingValues
Given an array, find the maximum and minimum value in the array and also find the values in range minimum and maximum that are absent in the array.
*/
func missingValues(arr []int) (int, int, []int) {
	hash := make(map[int]bool)
	min, max := arr[0], arr[0]
	missing := []int{}

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		hash[num] = true
	}
	for i := min + 1; i < max; i++ {
		_, ok := hash[i]
		if ok == false {
			missing = append(missing, i)
		}
	}
	return min, max, missing
}

/*
OddCountElements
Given an array in which all the elements appear even number of times except one. which appear odd number of times. Find the element which appear odd number of times.
*/
func oddCount(arr []int) int {
	numHash := make(map[int]int)
	var output int
	for _, num := range arr {
		numHash[num]++
	}
	for key, freq := range numHash {
		if freq%2 != 0 {
			output = key
			break
		}
	}
	return output
}

/*
SumDistinct
Given an array of size N. the elements in the array may be repeated. You need to find sum of distinct elements of the array.
If there is some value repeated continuously then they should be added once.
*/
func sumDistinct(arr []int) int {
	sum := 0
	size := len(arr)
	sort.Ints(arr)
	for i := 0; i < size; i++ {
		if i == size-1 && arr[i] != arr[i-1] {
			sum += arr[i]
		} else if arr[i] != arr[i+1] {
			sum += arr[i]
		}
	}
	return sum // Time O(n log(n)), space O(1)
}

/*
TwoElementsWhoseSumIsClosestToZero
In a given list of integers, both +v and -v. You need to find the two elements such that their sum is closest to zero
*/
func minAbsSumPair(arr []int) [2]int {
	res := [2]int{}
	size := len(arr)
	minDiff := math.MaxInt64
	left, right := 0, size-1
	sort.Ints(arr)
	for right >= left {
		sum := arr[left] + arr[right]
		if math.Abs(float64(sum-0)) < math.Abs(float64(0-minDiff)) {
			minDiff = sum
			res[0] = arr[left]
			res[1] = arr[right]
		}
		if sum == 0 {
			break
		} else if sum < 0 {
			left++
		} else {
			right--
		}
	}
	return res
}

/*
FindPair
Given an array of n numbers, find two elements such that their sum is equal to "value"
*/
func findPair(arr []int, value int) [2]int {
	res := [2]int{}
	hash := make(map[int]int)
	for _, num := range arr {
		remainder := value - num
		_, ok := hash[remainder]
		if ok {
			res[0] = num
			res[1] = remainder
			return res
		}
		hash[num] = num
	}
	res[0] = -1
	res[1] = -1
	return res
}

func main() {
	firstRepeatedInput := []int{7, 1, 6, 3, 5, 1, 7, 4, 2}
	firstRepeatedRes := firstRepeated(firstRepeatedInput)
	fmt.Printf("\nFirstRepeated:\nInput: %v\nOutput: %v\n", firstRepeatedInput, firstRepeatedRes)

	printDuplicates(firstRepeatedInput)

	removeDuplicatesRes := removeDuplicates(firstRepeatedInput)
	fmt.Printf("\nRemoveDuplicates:\nInput: %v\nOutput: %v\n", firstRepeatedInput, removeDuplicatesRes)

	findMissingNumberInput := []int{1, 3, 4, 5}
	findMissingNumberRes := findMissingNumber(findMissingNumberInput)
	fmt.Printf("\nFindMissingNumber:\nInput: %v\nOutput: %v\n", findMissingNumberInput, findMissingNumberRes)

	missingValuesInput := []int{6, 2, 1, 2}
	min, max, missing := missingValues(missingValuesInput)
	fmt.Printf("\nMissingValues:\nInput: %v\nOutput: min: %d, max: %d missing: %v\n", missingValuesInput, min, max, missing)

	oddCountInput := []int{1, 4, 2, 4, 3, 1, 2}
	oddCountOutput := oddCount(oddCountInput)
	fmt.Printf("\nOddCount: \nInput: %v\nOutput: %v\n", oddCountInput, oddCountOutput)

	sumDistinctInput := []int{1, 9, 2, 4, 3, 5, 4, 5}
	sumDistinctOutput := sumDistinct(sumDistinctInput)
	fmt.Printf("\nSumDistinct:\nInput: %v\nOutput: %d\n", sumDistinctInput, sumDistinctOutput)

	minAbsSumPairInput := []int{1, 5, -10, 3, 2, -6, 8, 9, 6}
	minAbsSumPairOutput := minAbsSumPair(minAbsSumPairInput)
	fmt.Printf("\nMinAbsSumPairO:\nInput: %v\nOutput: %d\n", minAbsSumPairInput, minAbsSumPairOutput)

	findPairInput := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	findPairOutput := findPair(findPairInput, 8)
	fmt.Printf("\nFindPair:\nInput: %v\nOutput: %d\n", findPairInput, findPairOutput)
}
