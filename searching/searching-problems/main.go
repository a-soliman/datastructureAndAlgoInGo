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

/*
ArithmeticProgressionTriplet
Given a sorted array, find all Arithmetic progression triplet possible
*/
func apTriplets(arr []int) [][3]int {
	size := len(arr)
	var (
		i, j, k, sum, target int
	)
	res := [][3]int{}
	for i = 1; i < size; i++ {
		j, k = i-1, i+1
		target = 2 * arr[i]
		for j >= 0 && k < size {
			sum = arr[j] + arr[k]
			if sum == target {
				res = append(res, [3]int{arr[j], arr[i], arr[k]})
				j--
				k++
			} else if sum < target {
				k++
			} else {
				j--
			}
		}
	}
	return res
}

/*
GeometricProgressionTriplet
Given a sorted array, find all Geometric progression triplet possible
*/
func gpTriplets(arr []int) [][3]int {
	size := len(arr)
	res := [][3]int{}
	var (
		i, j, k, sum, target int
	)

	for i = 1; i < size-1; i++ {
		j, k = i-1, i+1
		target = arr[i] * arr[i]
		for j >= 0 && k < size {
			sum = arr[j] * arr[k]
			if sum == target {
				res = append(res, [3]int{arr[j], arr[i], arr[k]})
				j--
				k++
			} else if sum < target {
				k++
			} else {
				j--
			}
		}
	}
	return res
}

/*
MajorityElementInAnArray
Given an array of n elements. Find the majority element. which appears more than n/2 times.
Return 0 incase there is no majority.
*/
func getMajority(arr []int) int {
	size := len(arr)
	maj, freq, count := arr[0], 1, 0
	for i := 1; i < size; i++ {
		if arr[i] == maj {
			freq++
		} else {
			freq--
			if freq == 0 {
				maj = arr[i]
				freq = 1
			}
		}
	}
	for _, num := range arr {
		if num == maj {
			count++
		}
	}
	if count >= size/2 {
		return maj
	}
	return 0
}

/*
MajorityElementInSortedArr
Given a sorted array of n elements. Find the majority element. which appears more than n/2 times.
Return 0 incase there is no majority.
*/
func getMajorityInSorted(arr []int) int {
	size := len(arr) - 1
	mid := size / 2
	candidate := arr[mid]
	start, end := getStart(arr, mid, candidate), getEnd(arr, mid, candidate)
	if end-start+1 > mid {
		return candidate
	}
	return 0
}

func getStart(arr []int, end, val int) int {
	start, mid := 0, 0
	for end > start {
		mid = (start + end) / 2
		if arr[mid] == val && (mid == 0 || arr[mid-1] != val) {
			return mid
		}
		start = mid + 1
	}
	return end
}

func getEnd(arr []int, start, val int) int {
	end, mid := len(arr)-1, 0
	for end > start {
		mid = (start + end) / 2
		if arr[mid] == val && (mid == end || arr[mid+1] != val) {
			return mid
		}
		end = mid - 1
	}
	return start
}

/*
kthLargestNumber
Given an unsorted array of n distinct elements. How will you identify the kth largest element with minimum number of comparisons?
*/
func kthLargest(arr []int, k int) int {
	size := len(arr)
	buildHeap(arr)
	for i := 1; i < k; i++ {
		arr[0], arr[size-1] = arr[size-1], arr[0]
		size--
		heapify(arr, size)
	}
	return arr[0]
}

func buildHeap(arr []int) {
	var (
		parentIndex int
		childIndex  int
	)
	for i := 1; i < len(arr); i++ {
		childIndex = i
		parentIndex = getParent(childIndex)
		for parentIndex >= 0 && arr[parentIndex] < arr[childIndex] {
			arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]
			childIndex = parentIndex
			parentIndex = getParent(childIndex)
		}
	}
}

func heapify(arr []int, size int) {
	parentIndex := 0
	childIndex := getChild(arr, size, parentIndex)
	for childIndex < size && arr[parentIndex] < arr[childIndex] {
		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]
		parentIndex = childIndex
		childIndex = getChild(arr, size, parentIndex)
	}
}

func getParent(childIdx int) int {
	if childIdx%2 == 0 {
		return (childIdx / 2) - 1
	}
	return childIdx / 2
}

func getChild(arr []int, size int, parentIdx int) int {
	childIdx1 := (parentIdx * 2) + 1
	childIdx2 := childIdx1 + 1

	if childIdx1 > size {
		return math.MaxInt64
	}
	if childIdx2 > size {
		return childIdx1
	}
	if arr[childIdx1] > arr[childIdx2] {
		return childIdx1
	}
	return childIdx2
}

/*
FindMedianOfAnUnsortedList
in an unsorted list of numbers of size n, if all the elements of the array are sorted then find the element which lie at the index n/2.
*/
func findMedian(arr []int) int {
	size := len(arr)
	if size < 2 {
		return arr[0]
	}
	return quickSearch(arr, (size+1)/2)
}

func quickSearch(arr []int, k int) int {
	left, right := []int{}, []int{}
	size := len(arr)

	for i := 1; i < size; i++ {
		if arr[i] > arr[0] {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	leftSize := len(left)
	if leftSize == k-1 {
		return arr[0]
	}
	if leftSize > k-1 {
		return quickSearch(left, k)
	}
	return quickSearch(right, k-leftSize-1)
}

/*
FindMaximaInBitonicList
A bitonic list comprises of an increasing sequence of integers immediately followed by a decreasing sequence. Find maxima in bitonic list
*/
func findBitonicArrMax(arr []int) int {
	start, end, mid := 0, len(arr)-1, 0
	currentVal := 0

	for start <= end {
		mid = (start + end) / 2
		currentVal = arr[mid]
		if currentVal > arr[mid+1] && currentVal > arr[mid-1] {
			break
		}
		if currentVal < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return currentVal
}

/*
StockPurchaseSell
in a given list, in which nth element is the price of the stock on nth day. You are asked to buy once and sell once. on what date you
will be buying nad at what date you will be selling to get maximum profit.
*/
func maxProfit(stocks []int) (buy int, sell int, profit int) {
	currentMin, currentMax, currentProfit := stocks[0], stocks[0], 0

	for _, stock := range stocks {
		if stock > currentMax {
			currentMax = stock
		} else if stock < currentMin {
			currentProfit = currentMax - currentMin
			if currentProfit > profit {
				profit = currentProfit
				buy = currentMin
				sell = currentMax
			}
			currentMin = stock
			currentMax = stock
		}
	}
	if currentProfit > profit {
		profit = currentProfit
		buy = currentMin
		sell = currentMax
	}
	return
}

/*
FindMedianOfTwoSortedLists
Given Two sorted lists, find the median of the arrays if they are compained to form a bigger list
*/
func findMedianOfTwoSortedLists(arr1 []int, arr2 []int) int {
	size1, size2 := len(arr1), len(arr2)
	medianIdx := int(math.Ceil(float64((size1 + size2) / 2)))
	i, j, count := 0, 0, 0
	res := 0

	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			res = arr1[i]
			i++
		} else {
			res = arr2[j]
			j++
		}
		if count == medianIdx {
			break
		}
		count++
	}
	return res
}

/*
Search01List
In given list of 0's and 1s in which all the 0's come before 1's. Write an algorithm to find the index of the first 1.
*/
func binarySearch01(arr []int) int {
	start, end, mid := 0, len(arr)-1, 0
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] == 1 && arr[mid-1] == 0 {
			break
		} else if arr[mid] == 0 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return mid
}

/*
FindMaxInARotatedArray
Given a sorted list S of N integers. S is rotated an unknown number of times. Find maximum value element in the array.
*/
func rotationMax(arr []int) int {
	start, end := 0, len(arr)-1
	idx := rotationMaxIndexUtil(arr, start, end)
	return arr[idx]
}

func rotationMaxIndexUtil(arr []int, start int, end int) int {
	if start >= end {
		return start
	}

	mid := (start + end) / 2
	if arr[mid] > arr[mid+1] {
		return mid
	} else if arr[start] < arr[mid] {
		return rotationMaxIndexUtil(arr, mid+1, end)
	} else {
		return rotationMaxIndexUtil(arr, start, mid-1)
	}
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

	apTripletsInput := []int{1, 2, 3, 4, 9, 17, 23}
	apTripletsOutput := apTriplets(apTripletsInput)
	fmt.Printf("\nAPTriplets:\nInput: %v\nOutput: %v\n", apTripletsInput, apTripletsOutput)

	gpTripletsInput := []int{1, 2, 3, 4, 9, 17, 23}
	gpTripletsOutput := gpTriplets(gpTripletsInput)
	fmt.Printf("\nGPTriplets:\nInput: %v\nOutput: %v\n", gpTripletsInput, gpTripletsOutput)

	getMajorityInput := []int{1, 5, 5, 13, 5, 31, 5}
	getMajorityOutput := getMajority(getMajorityInput)
	fmt.Printf("\nGetMajority:\nInput: %v\nOutput: %v\n", getMajorityInput, getMajorityOutput)

	getMajorityInSortedInput := []int{1, 5, 5, 5, 13, 31}
	getMajorityInSortedOutput := getMajorityInSorted(getMajorityInSortedInput)
	fmt.Printf("\ngetMajorityInSorted:\nInput: %v\nOutput: %v\n", getMajorityInSortedInput, getMajorityInSortedOutput)

	kthLargestInput := []int{5, 3, 2, 10, 9, 8, 12}
	kthLargestOutput := kthLargest(kthLargestInput, 2)
	fmt.Printf("\nkthLargest:\nInput: %v\nOutput: %v\n", kthLargestInput, kthLargestOutput)

	findMedianInput := []int{11, 1, 2, 5, 4, 13, 10}
	findMedianOutput := findMedian(findMedianInput)
	fmt.Printf("\nFindMedian:\nInput: %v\nOutput: %v\n", findMedianInput, findMedianOutput)

	findBitonicArrMaxInput := []int{1, 5, 10, 13, 20, 30, 8, 6, 5}
	findBitonicArrMaxOutput := findBitonicArrMax(findBitonicArrMaxInput)
	fmt.Printf("\nFindBitonicArrMax:\nInput: %v\nOutput: %v\n", findBitonicArrMaxInput, findBitonicArrMaxOutput)

	maxProfitInput := []int{10, 150, 6, 67, 61, 16, 86, 6, 67, 78, 150, 3, 28, 143}
	buy, sell, profit := maxProfit(maxProfitInput)
	fmt.Printf("\nMaxProfit:\nInput: %v\nOutput: %d, %d, %d \n", maxProfitInput, buy, sell, profit)

	findMedianOfTwoSortedListsInput1 := []int{1, 2, 3, 7}
	findMedianOfTwoSortedListsInput2 := []int{4, 5, 6}
	findMedianOfTwoSortedListsOutput := findMedianOfTwoSortedLists(findMedianOfTwoSortedListsInput1, findMedianOfTwoSortedListsInput2)

	fmt.Printf("\nFindMedianOfTwoSortedLists:\nInput 1: %v\nInput 2: %v\nOutput: %v \n", findMedianOfTwoSortedListsInput1, findMedianOfTwoSortedListsInput2, findMedianOfTwoSortedListsOutput)

	search01ListInput := []int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	search01ListOutput := binarySearch01(search01ListInput)
	fmt.Printf("\nSearch01List:\nInput: %v\nOutput: %v \n", search01ListInput, search01ListOutput)

	findMaxInARotatedArrayInput := []int{11, 12, 13, 8, 9, 10, 3, 5, 7}
	findMaxInARotatedArrayOutput := rotationMax(findMaxInARotatedArrayInput)
	fmt.Printf("\nFindMaxInARotatedArray:\nInput: %v\nOutput: %v \n", findMaxInARotatedArrayInput, findMaxInARotatedArrayOutput)

}
