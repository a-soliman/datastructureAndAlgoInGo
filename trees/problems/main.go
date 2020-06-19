package main

import "fmt"

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

func main() {
	btFromArrIterativeInput := []int{1, 2, 3, 4, 5, 6, 7}
	btFromArrIterativeOutput := btFromArrIterative(btFromArrIterativeInput)
	fmt.Printf("\nbtFromArrIterative: \nInput:%v\nOutput: %v\n", btFromArrIterativeInput, btFromArrIterativeOutput)

	btFromArrRecursiveInput := []int{1, 2, 3, 4, 5, 6, 7}
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
}
