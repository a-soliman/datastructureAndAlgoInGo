package main

import "fmt"

func factorial(n int) int {
	if n <= 2 {
		return 2
	}
	return n * factorial(n-1)
}

func main() {
	factorialInput := 5
	factorialOutput := factorial(factorialInput)
	fmt.Printf("\nFactorial:\nInput: %d\nOutput: %d\n", factorialInput, factorialOutput)
}
