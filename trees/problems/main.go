package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

/*
CreteCompleteBinaryTreeFromArr
Create a complete binary tree from values given in an array.
*/
func btFromArrIterative(arr []int) *node {
	queue := []*node{}
	currentNode := &node{
		arr[0],
		nil,
		nil,
	}
	root := currentNode

	for i := 1; i < len(arr); i++ {
		newNode := &node{arr[i], nil, nil}
		queue = append(queue, newNode)

		if currentNode.right != nil {
			currentNode = queue[0]
			queue = queue[1:]
		}
		if currentNode.left == nil {
			currentNode.left = newNode
		} else if currentNode.right == nil {
			currentNode.right = newNode
		}
	}
	return root
}

func btFromArrRecursive(arr []int) *node {
	return btFromArrRecursiveUtil(arr, 0, len(arr))
}

func btFromArrRecursiveUtil(arr []int, idx int, size int) *node {
	n := &node{arr[idx], nil, nil}
	left, right := 2*idx+1, 2*idx+2
	if left < size {
		n.left = btFromArrRecursiveUtil(arr, left, size)
	}
	if right < size {
		n.right = btFromArrRecursiveUtil(arr, right, size)
	}
	return n
}

/*
PreOrderTraversal
Perform preOrderTraversal of BT
*/
func preOrderTraversal(root *node) []int {
	res := []int{}
	traversalUtil(root, &res)
	return res
}

func traversalUtil(node *node, res *[]int) {
	*res = append(*res, node.value)
	if node.left != nil {
		traversalUtil(node.left, res)
	}
	if node.right != nil {
		traversalUtil(node.right, res)
	}
}

/*
PostOrderTraversal
Perform postOrderTraversal on BT
*/
func postOrderTraversal(root *node) []int {
	res := []int{}
	postOrderUtil(root, &res)
	return res
}

func postOrderUtil(node *node, res *[]int) {
	if node.left != nil {
		postOrderUtil(node.left, res)
	}
	if node.right != nil {
		postOrderUtil(node.right, res)
	}
	*res = append(*res, node.value)
}

/*
InOrderTraversal
Perform inOrderTraversal on BT
*/
func inOrderTraversal(root *node) []int {
	res := []int{}
	inOrderUtil(root, &res)
	return res
}

func inOrderUtil(node *node, res *[]int) {
	if node.left != nil {
		inOrderUtil(node.left, res)
	}
	*res = append(*res, node.value)
	if node.right != nil {
		inOrderUtil(node.right, res)
	}
}

/*
LevelOrderTraversal
Perform LevelOrderTraversal on BT
*/

func levelOrderTraversal(root *node) []int {
	res := []int{}
	queue := []*node{root}
	currentNode := queue[0]
	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		res = append(res, currentNode.value)
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
	return res
}

/*
PrintLevelOrderLineByLine
Perform level order traversal on BT. such that such that all levels are printed line by line
*/
func levelOrderLineByLine(node *node) [][]int {
	res := [][]int{}
	levelOrderLineByLineUtil(node, 0, &res)
	return res
}

func levelOrderLineByLineUtil(node *node, level int, res *[][]int) {
	if node == nil {
		return
	}
	if len(*res)-1 < level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], node.value)
	levelOrderLineByLineUtil(node.left, level+1, res)
	levelOrderLineByLineUtil(node.right, level+1, res)
}

func levelOrderLineByLine2(n *node) {
	q1 := []*node{}
	q2 := []*node{}
	q1 = append(q1, n)
	var item *node

	for len(q1) > 0 || len(q2) > 0 {
		for len(q1) > 0 {
			item = q1[0]
			q1 = q1[1:]
			fmt.Print(item.value, " ")
			if item.left != nil {
				q2 = append(q2, item.left)
			}
			if item.right != nil {
				q2 = append(q2, item.right)
			}
		}
		fmt.Print("\n")
		for len(q2) > 0 {
			item = q2[0]
			q2 = q2[1:]
			fmt.Print(item.value, " ")
			if item.left != nil {
				q1 = append(q1, item.left)
			}
			if item.right != nil {
				q1 = append(q1, item.right)
			}
		}
		fmt.Print("\n")
	}
}

func levelOrderLineByLine3(n *node) {
	queue := []*node{n}
	var item *node
	count, newCount := 1, 0

	for len(queue) > 0 {
		for count > 0 {
			count = count - 1
			item = queue[0]
			queue = queue[1:]
			fmt.Print(item.value, " ")
			if item.left != nil {
				newCount = newCount + 1
				queue = append(queue, item.left)
			}
			if item.right != nil {
				newCount = newCount + 1
				queue = append(queue, item.right)
			}
		}
		if count == 0 {
			count = newCount
			newCount = 0
			fmt.Print("\n")
		}
	}
}

func main() {
	btFromArrIterativeInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	btFromArrIterativeOutput := btFromArrIterative(btFromArrIterativeInput)
	fmt.Printf("\nbtFromArrIterative: \nInput:%v\nOutput: %v\n", btFromArrIterativeInput, btFromArrIterativeOutput)

	btFromArrRecursiveInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	btFromArrRecursiveOutput := btFromArrRecursive(btFromArrRecursiveInput)
	fmt.Printf("\nbtFromArrRecursive: \nInput:%v\nOutput: %v\n", btFromArrRecursiveInput, btFromArrRecursiveOutput)

	preOrderTraversalInput := btFromArrRecursiveOutput
	preOrderTraversalOutput := preOrderTraversal(preOrderTraversalInput)
	fmt.Printf("\nPreOrderTraversal:\nInput: %v\nOutput:%v\n", preOrderTraversalInput, preOrderTraversalOutput)

	postOrderTraversalInput := btFromArrRecursiveOutput
	postOrderTraversalOutput := postOrderTraversal(postOrderTraversalInput)
	fmt.Printf("\npostOrderTraversal:\nInput: %v\nOutput:%v\n", postOrderTraversalInput, postOrderTraversalOutput)

	inOrderTraversalInput := btFromArrRecursiveOutput
	inOrderTraversalOutput := inOrderTraversal(inOrderTraversalInput)
	fmt.Printf("\ninOrderTraversal:\nInput: %v\nOutput:%v\n", inOrderTraversalInput, inOrderTraversalOutput)

	levelOrderTraversalInput := btFromArrRecursiveOutput
	levelOrderTraversalOutput := levelOrderTraversal(levelOrderTraversalInput)
	fmt.Printf("\nlevelOrderTraversal:\nInput: %v\nOutput:%v\n", levelOrderTraversalInput, levelOrderTraversalOutput)

	levelOrderLineByLineInput := btFromArrRecursiveOutput
	levelOrderLineByLineOutput := levelOrderLineByLine(levelOrderLineByLineInput)
	fmt.Printf("\nlevelOrderLineByLine:\nInput: %v\nOutput: %v\n", levelOrderLineByLineInput, levelOrderLineByLineOutput)

	fmt.Printf("\nLevelOrderLinByLine2:\n")
	levelOrderLineByLine2(btFromArrRecursiveOutput)
	fmt.Printf("\nLevelOrderLinByLine3:\n")
	levelOrderLineByLine3(btFromArrRecursiveOutput)
}
