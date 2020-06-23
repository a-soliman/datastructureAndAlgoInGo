package main

import (
	"fmt"
	"math"
	"sort"
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

/*
Group the numbers
You are given an integer array arr of size n. Group and rearrange them (in-place) into evens and odds in such a way that group of all even integers appears on the left side and group of all odd integers appears on the right side.
Example
Input: [1, 2, 3, 4]
Output: [4, 2, 1, 3]
Order does not matter. Other valid solutions are:
[2, 4, 1, 3]
[2, 4, 3, 1]
[4, 2, 3, 1]
*/
func groupNumbers(arr []int32) []int32 {
	size := len(arr)
	if size < 2 {
		return arr
	}
	i, j := 0, size-1
	for i < j {
		for i < size && isEven(arr[i]) {
			i++
		}
		for j >= 0 && !isEven(arr[j]) {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return arr
}

func isEven(num int32) bool {
	return num%2 == 0
}

/*
Top K
You are given an array of integers, arr, of size n, which is analogous to a continuous stream of integers input. Your task is to find K largest elements from a given stream of numbers.

By definition, we don't know the size of the input stream. Hence, produce K largest elements seen so far, at any given time. For repeated numbers, return them only once.
If there are less than K distinct elements in arr, return all of them.
Don't rely on size of input array arr.
Feel free to use built-in functions if you need a specific data-structure.


Example One
Input: arr = [1, 5, 4, 4, 2]; K = 2
Output: [4, 5]

Example Two
Input: arr = [1, 5, 1, 5, 1]; K = 3
Output: [5, 1]

Notes

Input Parameters: There is only one argument: Integer array arr.
Output: Return an integer array res, containing K largest elements. If there are less than K unique elements, return all of them. Order of elements in res does not matter.
Constraints:
1 <= n <= 10^5
1 <= K <= 10^5
arr may contain duplicate numbers.
arr may or may not be sorted
*/
func topK(arr []int32, k int32) []int32 {
	size := len(arr)
	dict := make(map[int32]bool)
	res := []int32{}
	var i int32
	buildHeap(arr)
	for i = 1; i <= k && size > 0; i++ {
		arr[0], arr[size-1] = arr[size-1], arr[0]
		_, ok := dict[arr[size-1]]
		if !ok {
			dict[arr[size-1]] = true
			res = append(res, arr[size-1])
		} else {
			i--
		}
		size--
		heapify(arr, size)
	}
	return res
}

func buildHeap(arr []int32) {
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

func heapify(arr []int32, size int) {
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

func getChild(arr []int32, size int, parentIdx int) int {
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
3 Sum
Given an integer array arr of size n, find all magic triplets in it.
Magic triplet is a group of three numbers whose sum is zero.
Note that magic triplets may or may not be made of consecutive numbers in arr.

Example One
Input: arr = [10, 3, -4, 1, -6, 9]
Output: [“10,-4,-6”, “3,-4,1”]

Example Two
Input: arr = [12, 34, -46]
Output: [“12,-46,34”]

Example Three
Input: arr = [0, 0, 0];
Output: [“0,0,0”]

Example Four
Input: arr = [-2, 2, 0 -2, 2];
Output: [“2,-2,0”]
*/

func findZeroSum(arr []int32) []string {
	size := len(arr)
	res := []string{}
	left, right := 0, size-1
	localSum := int32(0)
	var localRes string
	dict := make(map[string]bool)
	if size < 3 {
		return res
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i, num := range arr {
		left, right = i+1, size-1
		for left < right {
			localSum = num + arr[left] + arr[right]
			if localSum == 0 {
				localRes = fmt.Sprintf("%d, %d, %d", num, arr[left], arr[right])
				_, ok := dict[localRes]
				if !ok {
					res = append(res, localRes)
					dict[localRes] = true
				}
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if localSum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func dutchFlagSort(balls []string) {
	size := len(balls)
	nextLeft, nextRight := 0, size-1
	// rearrange R
	nextLeft = checkAndSwap(balls, nextLeft, nextRight, "R")
	nextRight = size - 1
	nextLeft = checkAndSwap(balls, nextLeft, nextRight, "G")
}

func checkAndSwap(arr []string, left int, right int, key string) int {
	size := len(arr)
	for left < right && left < size && right >= 0 {
		for arr[left] == key {
			left++
			if left >= size {
				break
			}
		}
		for arr[right] != key {
			right--
			if right == 0 {
				break
			}
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func mergeFirstIntoSecondfunc(arr1 []int32, arr2 []int32) {
	size1, size2 := len(arr1), len(arr2)
	largest1, largest2 := size1-1, size2-1
	nextPosition := size2 - 1

	// rearrange largest2
	for largest2 >= 0 {
		if arr2[largest2] == 0 {
			largest2--
		}
	}
	for largest1 >= 0 && largest2 >= 0 && nextPosition >= 0 {
		if largest1 >= largest2 {
			arr2[nextPosition] = arr1[largest1]
			largest1--
		} else {
			arr2[nextPosition] = arr2[largest2]
			largest2--
		}
		nextPosition--
	}
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
	fmt.Printf("\nMergeKSortedArrays: \nOutput: %v\n", mergeKArraysOutput)

	groupNumbersInput := []int32{2, 90}
	groupNumbersOutput := groupNumbers(groupNumbersInput)

	fmt.Printf("\nGroupNumbers:\nInput: %v \nOutput: %v\n", groupNumbersInput, groupNumbersOutput)

	topKInput := []int32{4, 8, 9, 6, 6, 2, 10, 2, 8, 1, 2, 9}
	topKOutput := topK(topKInput, 11)
	fmt.Printf("\nTopK: \nInput: %v \nOutput: %v\n", topKInput, topKOutput)

	findZeroSumInput := []int32{-2, 2, 0, -2, 2}
	findZeroSumOutput := findZeroSum(findZeroSumInput)
	fmt.Printf("\nFindZeroSumOutput: \nInput: %v \nOutput: %v\n", findZeroSumInput, findZeroSumOutput)

	// dutchFlagSortInput := []string{"G", "B", "G", "G", "R", "B", "R", "G"}
	dutchFlagSortInput := []string{"B", "R", "R", "R"}
	dutchFlagSort(dutchFlagSortInput)
	fmt.Printf("\nDutch National Flag: \nOutput: %v\n", dutchFlagSortInput)

	testArr1 := []int32{1, 3, 5}
	testArr2 := []int32{2, 4, 6, 0, 0, 0}
	mergeFirstIntoSecondfunc(testArr1, testArr2)
	fmt.Printf("\n Merge:  \nOutput: %v\n", testArr2)

}
