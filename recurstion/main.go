package main

import (
	"fmt"
	"strconv"
)

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
		if num <= 2 {
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

func committees(n int, k int) int {
	// base case
	if k == 0 || k == n {
		return 1
	}
	return committees(n-1, k-1) + committees(n-1, k)
}

func towerOfHanoi(n int, src rune, dist rune, helper rune) {
	// base case only one desc to move
	if n == 1 {
		fmt.Printf("move a desc from %v to %v\n", string(src), string(dist))
	} else {
		towerOfHanoi(n-1, src, helper, dist)
		fmt.Printf("move a desc from %v to %v\n", string(src), string(dist))
		towerOfHanoi(n-1, helper, dist, src)
	}
}

func binaryStrings(n int) []string {
	res := []string{}
	var helper func(string)
	helper = func(str string) {
		if len(str) == n {
			res = append(res, str)
			return
		}
		helper(str + "0")
		helper(str + "1")
	}
	helper("")
	return res
}

func decStrings(n int) []string {
	res := []string{}
	var helper func(string)
	helper = func(str string) {
		if len(str) == n {
			res = append(res, str)
			return
		}
		for i := 0; i < 10; i++ {
			helper(str + strconv.Itoa(i))
		}
	}
	helper("")
	return res
}

func permutation(str string) []string {
	res := []string{}
	permutationHelper("", str, &res)
	return res
}

func permutationHelper(currentPerm string, remaining string, res *[]string) {
	if len(remaining) == 0 {
		*res = append(*res, currentPerm)
		return
	}
	var newRemaining string
	for i, char := range remaining {
		newRemaining = remaining[0:i] + remaining[i+1:len(remaining)]
		permutationHelper(currentPerm+string(char), newRemaining, res)
	}
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

	fibonacciInput := 4
	fibonacciOutput := fibonacci(fibonacciInput)
	fmt.Printf("\nFibonacci:\nInput: %d\nOutput: %d\n", fibonacciInput, fibonacciOutput)

	committeesStudents := 4
	committeesRows := 2
	numOfCommittees := committees(committeesStudents, committeesRows)
	fmt.Printf("\nCommittees: \nStudents: %d, Rows: %d\nOutput: %d\n", committeesStudents, committeesRows, numOfCommittees)

	towerOfHanoi(3, 'A', 'B', 'C')

	binaryStringsInput := 3
	binaryStringsOutput := binaryStrings(binaryStringsInput)
	fmt.Printf("\nBinaryStrings:\nInput: %d\nOutput: %v\n", binaryStringsInput, binaryStringsOutput)

	decStringsInput := 1
	decStringsOutput := decStrings(decStringsInput)
	fmt.Printf("\nDecStrings:\nInput: %d\nOutput: %v\n", decStringsInput, decStringsOutput)

	permutationInput := "abc"
	permutationOutput := permutation(permutationInput)
	fmt.Printf("\nPerutation:\nInput: %s\nOutput: %v\n", permutationInput, permutationOutput)
}
