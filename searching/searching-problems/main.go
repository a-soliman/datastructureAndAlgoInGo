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

/*
FindMinDiff
Given an array of integers, find the element pair with minimum difference (return the diff)
*/
func findMinDiff(arr []int) int {
	minDiff := math.MaxInt64
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		sum := arr[i] - arr[i-1]
		if sum < minDiff {
			minDiff = sum
		}
	}
	return minDiff
}

/*
minDiffPairOfTwoArrays
Given two arrays, find the minimum difference pair such that it should take one element from each array
*/
func minDiffPairOfTwoArrays(arr1 []int, arr2 []int) [2]int {
	minDiff := math.MaxInt64
	diff := 0
	sort.Ints(arr1)
	sort.Ints(arr2)
	i, j := 0, 0
	out1, out2 := 0, 0
	size1, size2 := len(arr1), len(arr2)

	for i < size1 && j < size2 {
		diff = abs(arr1[i], arr2[j])

		if minDiff > diff {
			minDiff = diff
			out1 = arr1[i]
			out2 = arr2[j]
		}
		if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return [2]int{out1, out2}
}

/*
ClosestPair
Given an array of positive integers. You need to find a pair in array whose sum is closest to given number.
*/
func closestSum(arr []int, value int) [2]int {
	sort.Ints(arr) // O(n log(n))
	res := [2]int{}
	i, j := 0, len(arr)-1
	closestS := math.MaxInt64
	sum := 0
	localSum := 0
	for i < j { // O(n)
		localSum = arr[i] + arr[j]
		sum = abs(value, localSum)
		if closestS > sum {
			closestS = sum
			res[0] = arr[i]
			res[1] = arr[j]
		}
		if localSum > value {
			j--
		} else if localSum < value {
			i++
		} else if localSum == value {
			break
		}
	}
	return res
}

/*
SumPairRestArr
Given an array find if there is a pair whose sum is equal to the sum of rest of the elements of the array
*/
func sumPairRestArr(arr []int) [2]int {
	res := [2]int{-1, -1}
	sort.Ints(arr)
	i, j := 0, len(arr)-1
	totalSum := sumArr(arr)
	sum := 0
	rest := 0
	for i < j {
		sum = arr[i] + arr[j]
		rest = totalSum - sum
		if sum == rest {
			res[0] = arr[i]
			res[1] = arr[j]
			break
		} else if sum < rest {
			i++
		} else {
			j--
		}
	}
	return res
}

/*
ZeroSumTriplets
Given an array of integers, you need to find all the triplets whose sum = 0
*/
func zeroSumTriplets(arr []int) [][3]int {
	res := [][3]int{}
	sort.Ints(arr)
	size := len(arr)
	start, stop := 0, size-1
	sum := 0
	for i := 0; i < size/2; i++ {
		start = i + 1
		stop = size - 1

		for start < stop {
			sum = arr[i] + arr[start] + arr[stop]
			if sum == 0 {
				res = append(res, [3]int{arr[i], arr[start], arr[stop]})
				start++
				stop--
			} else if sum > 0 {
				stop--
			} else {
				start++
			}
		}
	}

	return res
}

func abs(num1 int, num2 int) int {
	return int(math.Abs(float64(num1) - float64(num2)))
}

func sumArr(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func findDuplicatesInSortedArr(arr1 []int, arr2 []int) []int {
	size1, size2 := len(arr1), len(arr2)
	i, j := 0, 0
	res := []int{}
	for i < size1 && j < size2 {
		if arr1[i] == arr2[j] {
			if len(res) == 0 || res[len(res)-1] != arr1[i] { // Can use set instead
				res = append(res, arr1[i])
			}
			i++
			j++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			i++
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

	findMinDiffInput := []int{1, 3, 2, 7, 8, 9}
	findMinDiffOutput := findMinDiff(findMinDiffInput)
	fmt.Printf("\nFindMinDiff:\nInput: %v\nOutput: %d\n", findMinDiffInput, findMinDiffOutput)

	minDiffPairOfTwoArraysInput1 := []int{1, 3, 2, 7, 8, 9}
	minDiffPairOfTwoArraysInput2 := []int{5, 10, 15}
	minDiffPairOfTwoArraysOutput := minDiffPairOfTwoArrays(minDiffPairOfTwoArraysInput1, minDiffPairOfTwoArraysInput2)
	fmt.Printf("\nMinDiffPairOfTwoArrays:\nInput: %v, %v\nOutput: %d\n", minDiffPairOfTwoArraysInput1, minDiffPairOfTwoArraysInput2, minDiffPairOfTwoArraysOutput)

	closestSumInput := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	closestSumOutput := closestSum(closestSumInput, 6)
	fmt.Printf("\nClosestSum:\nInput: %v\nOutput: %d\n", closestSumInput, closestSumOutput)

	sumPairRestArrInput := []int{1, 2, 4, 3, 7, 3}
	sumPairRestArrOutput := sumPairRestArr(sumPairRestArrInput)
	fmt.Printf("\nSumPairRestArr:\nInput: %v\nOutput: %d\n", sumPairRestArrInput, sumPairRestArrOutput)

	zeroSumTripletsInput := []int{1, 2, -4, 3, 7, -3}
	zeroSumTripletsOutput := zeroSumTriplets(zeroSumTripletsInput)
	fmt.Printf("\nZeroSumTriplets:\nInput: %v\nOutput: %d\n", zeroSumTripletsInput, zeroSumTripletsOutput)

	findDuplicatesInSortedArrInput1 := []int{2, 3, 3, 5, 4, 4, 6, 7, 7, 8, 12}
	findDuplicatesInSortedArrInput2 := []int{5, 5, 6, 8, 8, 9, 10, 16}
	findDuplicatesInSortedArrOutput := findDuplicatesInSortedArr(findDuplicatesInSortedArrInput1, findDuplicatesInSortedArrInput2)
	fmt.Printf("\nfindDuplicatesInSortedArr:\nInput: %v, %v\nOutput: %d\n", findDuplicatesInSortedArrInput1, findDuplicatesInSortedArrInput2, findDuplicatesInSortedArrOutput)

}
