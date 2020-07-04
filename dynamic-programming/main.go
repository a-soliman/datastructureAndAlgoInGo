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

func main() {
	fmt.Printf("FibonacciBottomUpSpaceEfficient: \nInput: %d\nOutput: %d\n", 6, fibonacciButtomUPSpaceEfficient(6))

}
