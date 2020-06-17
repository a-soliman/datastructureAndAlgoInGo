package main

import (
	"fmt"
	"strconv"
	"strings"
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

func combination(str string) []string {
	res := []string{}
	combinationHelper("", str, 0, &res)
	return res
}

func combinationHelper(currentCombination string, str string, idx int, res *[]string) {
	if idx >= len(str) {
		*res = append(*res, currentCombination)
		return
	}
	combinationHelper(currentCombination, str, idx+1, res)
	combinationHelper(currentCombination+string(str[idx]), str, idx+1, res)
}

func nQueens(n int) (bool, [][]int) {
	board := *makeBoard(n)
	success := nQUtil(0, board)
	return success, board
}

func makeBoard(n int) *[][]int {
	board := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}
	return &board
}

func nQUtil(col int, board [][]int) bool {
	size := len(board)
	if col == size {
		return true
	}
	for i := 0; i < size; i++ {
		if !isValid(col, i, board) {
			continue
		}
		board[i][col] = 1
		success := nQUtil(col+1, board)
		if success {
			return true
		}
		board[i][col] = 0
	}
	return false
}

func isValid(col int, row int, board [][]int) bool {
	size := len(board)
	// horizontal
	i, j := 0, 0
	for i < size {
		if board[row][i] == 1 {
			return false
		}
		i++
	}
	// vertical
	i = 0
	for i < size {
		if board[i][col] == 1 {
			return false
		}
		i++
	}
	// diagonal 1 find top left
	i, j = row, col
	for i > 0 && j > 0 {
		i--
		j--
	}

	for i < size && j < size {
		if board[i][j] == 1 {
			return false
		}
		i++
		j++
	}

	// diagonal 2 find top right
	i, j = row, col
	for i > 0 && j < size-1 {
		i--
		j++
	}

	for i < size && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i++
		j--
	}
	return true
}

/*
LetterCasePermutation
*/
func letterCasePermutation(str string) []string {
	res := []string{}
	letterCasePermutationHelper(str, "", &res)
	return res
}

func letterCasePermutationHelper(remaining string, currentPermutation string, res *[]string) {
	remainingSize := len(remaining)
	if remainingSize == 0 {
		*res = append(*res, currentPermutation)
		return
	}
	char := string(remaining[0])
	remaining = remaining[1:remainingSize]
	if strings.ToLower(char) != strings.ToUpper(char) {
		letterCasePermutationHelper(remaining, currentPermutation+strings.ToUpper(char), res)
	}
	letterCasePermutationHelper(remaining, currentPermutation+char, res)
}

/*
Generate Parentheses
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/
func generateParenthesis(n int) []string {
	res := []string{}
	generateParenthesisUtil(n, 0, "", 0, 0, &res)
	return res
}

func generateParenthesisUtil(limit int, idx int, current string, opened int, closed int, res *[]string) {
	if len(current) == limit*2 {
		*res = append(*res, current)
		return
	}
	if closed < limit && opened > closed {
		generateParenthesisUtil(limit, idx+1, current+")", opened, closed+1, res)
	}
	if opened < limit {
		generateParenthesisUtil(limit, idx+1, current+"(", opened+1, closed, res)
	}
}

/*
Generate All Possible Expressions That Evaluate To The Given Target Value

Generate All Possible Expressions That Evaluate To The Given Target Value



Given a string s that consists of digits (“0”..”9”) and target, a non-negative integer, find all expressions that can be built from string s that evaluate to the target.

When building expressions, you have to insert one of the following operators between each pair of consecutive characters in s: “join” or “*” or “+”. For example, by inserting different operators between the two characters of string “12” we can get either 12 (1 joined with 2) or 2 (1*2) or 3 (1+2).

Other operators such as “-” or “÷” are NOT supported.

Expressions that evaluate to the target but only utilize a part of s do not count: entire s has to be consumed.

Precedence of the operators is conventional: “join” has the highest precedence, “*” – medium and “+” has the lowest precedence. For example, 1+2*34=(1+(2*(34)))=1+68=69.

You have to return ALL expressions that can be built from string s and evaluate to the target.



Example One

Input:
s="222", target=24.
Output:
["22+2", "2+22"] and ["2+22", "22+2"] are both correct outputs.

22+2=24: we inserted the “join” operator between the first two characters and the “+” operator between the last two characters of s.
2+22=24: we inserted the “+” operator between the first two characters and the “join” operator between the last two characters of s.

Example Two
Input: s="1234", target=11.
Output: ["1+2*3+4"]

Example Three
Input:
s="99", target=1.
Output:
[]
Notes
Input Format: Function has two arguments: s and target.
*/

// func generateAllExpressions(s string, target int64) []string {
// 	res := []string{}
// 	generateAllExpressionsUtil(s, "", 0, 0, "", int(target), &res)
// 	return res
// }

// func generateAllExpressionsUtil(master string, current string, idx int, sum int, lastVal string, target int, res *[]string) {
// 	size := len(master)
// 	if idx >= size {
// 		*res = append(*res, current)
// 		return
// 	}
// 	currentItem := string(master[idx])
// 	currentItemVal, _ := strconv.Atoi(currentItem)
// 	sum = sum + currentItemVal
// 	lastChar := ""

// 	if idx > 0 {
// 		lastChar = string(current[len(current)-1])
// 	}

// 	lastValInt, _ := strconv.Atoi(lastVal)
// 	if lastChar == "*" {
// 		sum = sum - lastValInt - currentItemVal + (currentItemVal * lastValInt)
// 	}
// 	if idx < size-1 {
// 		generateAllExpressionsUtil(master, current+currentItem+"+", idx+1, sum, currentItem, target, res)
// 		generateAllExpressionsUtil(master, current+currentItem+"*", idx+1, sum, currentItem, target, res)
// 	}
// 	generateAllExpressionsUtil(master, current+currentItem, idx+1, sum, currentItem, target, res)
// }

func decomposePalindromStrings(str string) []string {
	res := []string{}
	decomposeUtil(str, 0, []string{}, &res)
	return res
}

func decomposeUtil(str string, idx int, current []string, res *[]string) {
	mainSize := len(str)
	if idx == mainSize {
		*res = append(*res, strings.Join(current, "|"))
		return
	}
	for i := idx; i < mainSize; i++ {
		if isPalindrom(str, idx, i) {
			palindromicSnippet := str[idx : i+1]
			current = append(current, palindromicSnippet)
			decomposeUtil(str, i+1, current, res)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrom(s string, start int, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
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

	combinationInput := "123"
	combinationOutput := combination(combinationInput)
	fmt.Printf("\nCombination:\nInput: %s\nOutput: %#v\n", combinationInput, combinationOutput)

	nQueensInput := 5
	success, board := nQueens(nQueensInput)
	fmt.Printf("\nNQueens:\nInput: %d\nOutput: Success: %v\nBoard: %v\n", nQueensInput, success, board)

	letterCasePermutationInput := "a1b2"
	letterCasePermutationOutput := letterCasePermutation(letterCasePermutationInput)
	fmt.Printf("\nLetterCasePermutation:\nInput: %s\nOutput: %v\n", letterCasePermutationInput, letterCasePermutationOutput)

	generateParenthesisInput := 3
	generateParenthesisOutput := generateParenthesis(generateParenthesisInput)
	fmt.Printf("\nGenerateParenthesis:\nInput: %d\nOutput: %#v\n", generateParenthesisInput, generateParenthesisOutput)

	// test := generateAllExpressions("1234", 11)
	// fmt.Printf("\nGenerateAllExpressions:\n %#v\n", test)

	decomposePalindromStringsInput := "racacbrbd"
	decomposePalindromStringsOutput := decomposePalindromStrings(decomposePalindromStringsInput)
	fmt.Printf("\ndecomposePalindromStrings:\n %#v\n", decomposePalindromStringsOutput)

}
