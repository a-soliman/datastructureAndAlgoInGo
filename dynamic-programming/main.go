package main

import (
	"fmt"
	"math"
)

/*
Fibonacci bottom-up space efficient
*/
func fibonacciButtomUPSpaceEfficient(n int) int {
	const tableSize = 3
	table := make([]int, tableSize)
	table[0] = 0
	table[1] = 1
	for i := 2; i <= n; i++ {
		table[i%tableSize] = table[(i-1)%tableSize] + table[(i-2)%tableSize]
	}
	return table[n%tableSize]
}

/*
Fibonacci top-bottom memoization
*/
func fibonacciMemoization(n int) int {
	hash := make(map[int]int)
	hash[0], hash[1] = 0, 1
	return fibonacciMemoizationUtil(n, &hash)
}

func fibonacciMemoizationUtil(n int, hash *map[int]int) int {
	_, found := (*hash)[n]
	if found {
		return (*hash)[n]
	}
	(*hash)[n] = fibonacciMemoizationUtil(n-1, hash) + fibonacciMemoizationUtil(n-2, hash)
	return (*hash)[n]
}

/*
StairCase
count the n steps ways a child can climb a stears of size n, if the child can take 1, 2, or 3 steps at a time
*/
func stairCase(n int) int {
	const tableSize = 4
	table := make([]int, tableSize)
	table[0], table[1], table[2] = 1, 2, 4
	for i := 3; i < n; i++ {
		table[i%tableSize] = table[(i-1)%tableSize] + table[(i-2)%tableSize] + table[(i-3)%tableSize]
	}
	return table[n-1]

}

/*
Minimum cost stair climb
given a cost array of steps count the minimum cost to reach the final step, you can move 1 or 2 steps at a time
*/
func minCostStairClimb(stairCost []int) int {
	costSize := len(stairCost)
	resSize := costSize + 2
	res := make([]int, resSize)
	res[1] = stairCost[0]
	for i := 2; i < resSize-1; i++ {
		res[i] = int(math.Min(float64(stairCost[i-1]+res[i-1]), float64(stairCost[i-1]+res[i-2])))
	}
	res[resSize-1] = int(math.Min(float64(res[resSize-2]), float64(res[resSize-3])))
	return res[resSize-1]
}

func combinations(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	table := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]int, k+1)
	}
	for row := 0; row <= n; row++ {
		table[row][0] = 1
	}
	for col := 0; col <= k; col++ {
		table[col][col] = 1
	}
	for row := 2; row <= n; row++ {
		for col := 1; col < row; col++ {
			table[row][col] = table[row-1][col] + table[row-1][col-1]
		}
	}
	fmt.Println(table)
	return table[n][k]
}

/*
CountUniquePathsUsing DP
given a m by n matrix starting from the top left, count how many unique paths can you reach the bottom left cell, using DP
*/
func countUniquePaths(m, n int) int {
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}
	for row := 0; row < m; row++ {
		table[row][0] = 1
	}
	for col := 0; col < n; col++ {
		table[0][col] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			table[row][col] = table[row-1][col] + table[row][col-1]
		}
	}
	return table[m-1][n-1]
}

/*
MaxPathSum
given a m by n grid, starting from the top left, return the sum of the max path to reach the bottom right cell
*/
func maxPathSum(grid [][]int) int {
	rowsLen := len(grid)
	colsLen := len(grid[0])
	res := make([][]int, rowsLen)
	for row := 0; row < rowsLen; row++ {
		res[row] = make([]int, colsLen)
	}
	res[0][0] = grid[0][0]

	// prefill base cases
	for col := 1; col < colsLen; col++ {
		res[0][col] = grid[0][col] + res[0][col-1]
	}
	for row := 1; row < rowsLen; row++ {
		res[row][0] = grid[row][0] + res[row-1][0]
	}

	for row := 1; row < rowsLen; row++ {
		for col := 1; col < colsLen; col++ {
			topValue := res[row-1][col]
			leftValue := res[row][col-1]
			res[row][col] = grid[row][col] + int(math.Max(float64(topValue), float64(leftValue)))
		}
	}
	return res[rowsLen-1][colsLen-1]
}

/*
MinCoinChange
*/
func minCoinChange(amount int, coins []int) int {
	table := make([]int, amount+1)
	for i := 1; i < len(table); i++ {
		table[i] = math.MaxInt64
	}
	for i := 1; i < len(table); i++ {
		for _, coin := range coins {
			if i-coin >= 0 && table[i-coin] != math.MaxInt64 {
				potential := table[i-coin] + 1
				if potential < table[i] {
					table[i] = potential
				}
			}
		}
	}
	if table[amount] == math.MaxInt64 {
		return -1
	}
	return table[amount]
}

/*
Levenshtein Distance

Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)
You have the following 3 operations permitted on a word:
a) Insert a character
b) Delete a character
c) Replace a character
The minimum number of steps required to convert word1 to word2 with the given set of allowed operations is called edit distance.
e.g. Minimum edit distance between the words 'kitten' and 'sitting', is 3.

kitten → sitten (substitution of "s" for "k")
sitten → sittin (substitution of "i" for "e")
sittin → sitting (insertion of "g" at the end)
*/
func levenshteinDistance(strWord1 string, strWord2 string) int32 {
	rowSize, colSize := len(strWord1)+1, len(strWord2)+1
	table := make([][]int32, rowSize)
	for i := 0; i < rowSize; i++ {
		table[i] = make([]int32, colSize)
	}

	// fill the base cases
	for col := 1; col < colSize; col++ {
		table[0][col] = int32(col)
	}
	for row := 1; row < rowSize; row++ {
		table[row][0] = int32(row)
	}

	// iterate and build
	for row := 1; row < rowSize; row++ {
		for col := 1; col < colSize; col++ {
			if strWord1[row-1] == strWord2[col-1] {
				table[row][col] = table[row-1][col-1]
			} else {
				top, topLeft, left := table[row-1][col], table[row-1][col-1], table[row][col-1]
				min := getMin(top, topLeft, left)
				table[row][col] = min + 1
			}
		}
	}
	return table[rowSize-1][colSize-1]
}

func getMin(val1, val2, val3 int32) int32 {
	if val1 <= val2 && val1 <= val3 {
		return val1
	}
	if val2 <= val1 && val2 <= val3 {
		return val2
	}
	return val3
}

/*
Knight's tour!

Given a phone keypad as shown below:

1 2 3
4 5 6
7 8 9
– 0 –

How many different phone numbers of given length can be formed starting from the given digit? The constraint is that the movement from one digit to the next is similar to the movement of the Knight in chess.
For example if we are at 1 then the next digit can be either 6 or 8, if we are at 6 then the next digit can be 1, 7 or 0.
Repetition of digits is allowed, e.g. 1616161616 is a valid number.
The problem requires us to just give the count of different phone numbers and not necessarily list the numbers.
Find a polynomial-time solution, based on Dynamic Programming.

Example One
Input: startdigit = 1, phonenumberlength = 2
Output: 2
Two possible numbers of length 2: 16, 18.

Example Two
Input: startdigit = 1, phonenumberlength = 3
Output: 5
The possible numbers of length 3: 160, 161, 167, 181, 183
*/

func numPhoneNumbers(startdigit int32, phonenumberlength int32) int64 {
	numPad := buildNumPad()
	hash := make(map[int32]map[int32]int64)
	return numPhoneUtil(startdigit, phonenumberlength, &hash, &numPad)
}

func numPhoneUtil(startdigit int32, phonenumberlength int32, hash *map[int32]map[int32]int64, numPad *[][]int32) int64 {
	if phonenumberlength == 1 {
		return 1
	}

	levelExists, startdigitExists := false, false
	_, levelExists = (*hash)[phonenumberlength]
	if !levelExists {
		(*hash)[phonenumberlength] = make(map[int32]int64)
	}
	_, startdigitExists = (*hash)[phonenumberlength][startdigit]
	if !startdigitExists {
		(*hash)[phonenumberlength][startdigit] = 0
		neighbors := (*numPad)[startdigit]
		for _, neighbor := range neighbors {
			(*hash)[phonenumberlength][startdigit] += numPhoneUtil(neighbor, phonenumberlength-1, hash, numPad)
		}
	}
	return (*hash)[phonenumberlength][startdigit]
}

func buildNumPad() [][]int32 {
	numPad := make([][]int32, 10)
	numPad[0] = []int32{4, 6}
	numPad[1] = []int32{6, 8}
	numPad[2] = []int32{7, 9}
	numPad[3] = []int32{4, 8}
	numPad[4] = []int32{0, 3, 9}
	numPad[5] = []int32{}
	numPad[6] = []int32{0, 1, 7}
	numPad[7] = []int32{2, 6}
	numPad[8] = []int32{1, 3}
	numPad[9] = []int32{2, 4}
	return numPad
}

func main() {
	fmt.Printf("FibonacciBottomUpSpaceEfficient: \nInput: %d\nOutput: %d\n", 6, fibonacciButtomUPSpaceEfficient(6))
	fmt.Printf("\nFibonacciMemoization: \nInput: %d\nOutput: %d\n", 6, fibonacciMemoization(6))
	fmt.Printf("\nStairCase: \nInput: %d\nOutput: %d\n", 4, stairCase(4))
	fmt.Printf("\nCombinations: \nInput 4,3\nOutput: %d\n", combinations(4, 3))
	fmt.Printf("\nCountUniquePaths: \nInput: 2, 3\nOutput: %d \n", countUniquePaths(2, 3))
	maxPathSumInput := [][]int{{1, 3, 1}, {1, 5, 1}, {10, 2, 1}}
	maxPathSumOutput := maxPathSum(maxPathSumInput)
	fmt.Printf("\nMaxPathSum: \nInput: %v\nOutput: %d\n", maxPathSumInput, maxPathSumOutput)
	minCostStairClimbInput := []int{10, 15, 20, 25, 10}
	minCostStairClimbOutput := minCostStairClimb(minCostStairClimbInput)
	fmt.Printf("\nMinCostStairClimb: \nInput: %v\nOutput: %d\n", minCostStairClimbInput, minCostStairClimbOutput)
	minCoinAmount, minCoinCoins := 9, []int{1, 5, 7}
	minCoinOutput := minCoinChange(minCoinAmount, minCoinCoins)
	fmt.Printf("\nMinCoinChange: \nInput <amount> : %d, \nInput <coins> : %v, \nOutput: %d\n", minCoinAmount, minCoinCoins, minCoinOutput)

	levenshteinDistanceOutput := levenshteinDistance("pizza", "yolo")
	fmt.Printf("\nLevenshteinDistance: \n Output: %d\n", levenshteinDistanceOutput)
	test := numPhoneNumbers(1, 4)
	fmt.Println("test ", test)
}
