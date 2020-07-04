package main

import "fmt"

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

func main() {
	fmt.Printf("FibonacciBottomUpSpaceEfficient: \nInput: %d\nOutput: %d\n", 6, fibonacciButtomUPSpaceEfficient(6))
	fmt.Printf("\nFibonacciMemoization: \nInput: %d\nOutput: %d\n", 6, fibonacciMemoization(6))
	fmt.Printf("\nStairCase: \nInput: %d\nOutput: %d\n", 4, stairCase(4))
	fmt.Printf("\nCombinations: \nInput 4,3\nOutput: %d\n", combinations(4, 3))
}
