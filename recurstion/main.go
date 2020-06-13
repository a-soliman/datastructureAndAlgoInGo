package main

import "fmt"

func factorial(n int) int {
	if n <= 2 {
		return 2
	}
	return n * factorial(n-1)
}

func toPower(n int, k int) int {
	if k == 0 {
		return 1
	}
	return n * toPower(n, k-1)
}

func setSubSetCount(size int) int {
	if size == 0 {
		return 1
	}
	return 2 * setSubSetCount(size-1)
}

func fibonacci(n int) int {
	hash := make(map[int]int)

	var helper func(int, map[int]int) int

	helper = func(num int, hash map[int]int) int {
		if num < 2 {
			return num
		}
		_, ok := hash[num]
		if ok {
			return hash[num]
		}
		hash[num] = helper(num-1, hash) + helper(num-2, hash)
		return hash[num]
	}

	return helper(n, hash)
}

func main() {
	factorialInput := 5
	factorialOutput := factorial(factorialInput)
	fmt.Printf("\nFactorial:\nInput: %d\nOutput: %d\n", factorialInput, factorialOutput)

	toPowerInput1 := 2
	toPowerInput2 := 4
	toPowerOutput := toPower(toPowerInput1, toPowerInput2)
	fmt.Printf("\ntoPower:\nInput: N = %d, k = %d\nOutput: %d\n", toPowerInput1, toPowerInput2, toPowerOutput)

	setSubSetCountInput := 3
	setSubSetCountOutput := setSubSetCount(setSubSetCountInput)
	fmt.Printf("\nSetSubSetCount:\nInput: %d\nOutput: %d\n", setSubSetCountInput, setSubSetCountOutput)

	fibonacciInput := 100
	fibonacciOutput := fibonacci(fibonacciInput)
	fmt.Printf("\nFibonacci:\nInput: %d\nOutput: %d\n", fibonacciInput, fibonacciOutput)
}
