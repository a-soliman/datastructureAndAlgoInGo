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

/*
PrintSpiralTree
Given a binary tree, print the nodes breadth first in spiral order
*/
func printSpiralTree(n *node) {
	stack1, stack2 := []*node{n}, []*node{}
	var stackTop int
	var item *node
	for len(stack1) > 0 || len(stack2) > 0 {
		// print all in stack 1
		for i := 0; i < len(stack1); i++ {
			fmt.Print(stack1[i].value, " ")
		}
		fmt.Print("\n")

		for len(stack1) > 0 {
			stackTop = len(stack1) - 1
			item = stack1[stackTop]
			stack1 = stack1[0:stackTop]
			if item.right != nil {
				stack2 = append(stack2, item.right)
			}
			if item.left != nil {
				stack2 = append(stack2, item.left)
			}
		}

		// print all in stack2
		for i := 0; i < len(stack2); i++ {
			fmt.Print(stack2[i].value, " ")
		}
		fmt.Print("\n")

		for len(stack2) > 0 {
			stackTop = len(stack2) - 1
			item = stack2[stackTop]
			stack2 = stack2[0:stackTop]
			if item.left != nil {
				stack1 = append(stack1, item.left)
			}
			if item.right != nil {
				stack1 = append(stack1, item.right)
			}
		}

	}
}

/*
Print most left

*/
func printMostLeft(n *node) {
	queue := []*node{n}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			item := queue[0]
			queue = queue[1:]
			if item.left != nil {
				queue = append(queue, item.left)
			}
			if item.right != nil {
				queue = append(queue, item.right)
			}
			if size == len(queue)-1 {
				fmt.Print(item.value, "\n")
			}
			size--
		}
	}
}

func printMostLeftRec(n *node) {
	maxPrinted := 0
	printMostLeftUtil(n, 1, &maxPrinted)
}

func printMostLeftUtil(n *node, level int, maxPrinted *int) {
	if *maxPrinted < level {
		fmt.Print(n.value, "\n")
		*maxPrinted = *maxPrinted + 1
	}
	if n.left != nil {
		printMostLeftUtil(n.left, level+1, maxPrinted)
	}
	if n.right != nil {
		printMostLeftUtil(n.right, level+1, maxPrinted)
	}
}

func countUniValueTrees(n *node) int {
	if n.left == nil && n.right == nil { // the current node is a leaf
		return 1
	}
	left, right, total := 0, 0, 0
	if n.left != nil {
		left = countUniValueTrees(n.left)
	}
	if n.right != nil {
		right = countUniValueTrees(n.right)
	}
	total = left + right
	if n.value == n.left.value && n.value == n.right.value {
		total = left + right + 1
	}
	return total
}

/*
Given 2 trees, compare their leaf nodes left to right
*/

type leafNodeIterator struct {
	stack []*node
}

func (i *leafNodeIterator) isEmpty() bool {
	return len(i.stack) == 0
}

func (i *leafNodeIterator) hasNext() bool {
	return len(i.stack) > 0
}

func (i *leafNodeIterator) next() *node {
	var item *node

	for len(i.stack) > 0 {
		item = i.stack[len(i.stack)-1]
		i.stack = i.stack[0 : len(i.stack)-1]
		if item.left == nil && item.right == nil {
			break
		}
		if item.right != nil {
			i.stack = append(i.stack, item.right)
		}
		if item.left != nil {
			i.stack = append(i.stack, item.left)
		}
	}
	return item
}

func newLeafIterator(root *node) *leafNodeIterator {
	return &leafNodeIterator{
		[]*node{root},
	}
}

func compareLeafNodes(n1 *node, n2 *node) bool {
	tree1Iter := newLeafIterator(n1)
	tree2Iter := newLeafIterator(n2)
	for tree1Iter.hasNext() && tree2Iter.hasNext() {
		if tree1Iter.next().value != tree2Iter.next().value {
			return false
		}
	}
	return tree1Iter.isEmpty() && tree2Iter.isEmpty()
}

type llNode struct {
	value int
	prev  *llNode
	next  *llNode
}

/*
bstToDoublyLinkedList
given a bst convert it to doubly linkedlist
*/
func bstToLinkedList(n *node) (*llNode, *llNode) {
	_, head, tail := bstToLinkedListUtil(n, nil, nil)
	return head, tail
}

func bstToLinkedListUtil(n *node, first *llNode, last *llNode) (*llNode, *llNode, *llNode) {
	var prev, current, next *llNode
	if n.left != nil {
		prev, first, last = bstToLinkedListUtil(n.left, first, last)
	}
	current = &llNode{n.value, prev, nil}
	if current.prev != nil {
		current.prev.next = current
	}
	if n.right != nil {
		next, first, last = bstToLinkedListUtil(n.right, first, last)
		current.next = next
	}
	if current.next != nil {
		current.next.prev = current
	}
	if first == nil {
		first = current
	}
	last = current.next
	return current, first, last
}

/*
NthPreOrder
Given a bt, print the value of the node that will be at nth index when tree is traversed in pre order.
*/
func nthPreOrder(n *node, target int) int {
	counter := 0
	nthNode := node{}
	nthPreOrderUtil(n, target, &counter, &nthNode)
	return nthNode.value
}

func nthPreOrderUtil(n *node, target int, counter *int, res *node) {
	if n == nil {
		return
	}
	if n != nil {
		(*counter)++
	}
	if target == *counter {
		*res = *n
		return
	}
	nthPreOrderUtil(n.left, target, counter, res)
	nthPreOrderUtil(n.right, target, counter, res)
}

/*
NthPostOrder
Given a bt, print the values of node that will be at nth index when tree is traversed in post order
*/
func nthPostOrder(n *node, target int) int {
	counter := 0
	nthNode := node{}
	nthPostOrderUtil(n, target, &counter, &nthNode)
	return nthNode.value
}

func nthPostOrderUtil(n *node, target int, counter *int, res *node) {
	if n == nil {
		return
	}
	nthPostOrderUtil(n.left, target, counter, res)
	nthPostOrderUtil(n.right, target, counter, res)
	(*counter)++
	if *counter == target {
		*res = *n
		return
	}
}

/*
NthInOrder
Given a bt, print the values of node that will be at nth index when tree is traversed in order
*/
func nthInOrder(n *node, target int) int {
	counter := 0
	nthNode := node{}
	nthInOrderUtil(n, target, &counter, &nthNode)
	return nthNode.value
}

func nthInOrderUtil(n *node, target int, counter *int, res *node) {
	if n == nil {
		return
	}
	nthInOrderUtil(n.left, target, counter, res)
	(*counter)++
	if *counter == target {
		*res = *n
		return
	}
	nthInOrderUtil(n.right, target, counter, res)
}

/*
PrintAllPaths
Given a bt, print all the paths from the roots to the leaf
*/
func findAllPaths(n *node) [][]int {
	res := [][]int{}
	current := []int{}
	findAllPathsUtil(n, current, &res)
	return res
}

func findAllPathsUtil(n *node, current []int, res *[][]int) {
	current = append(current, n.value)
	if n.left == nil && n.right == nil {
		fullPath := make([]int, len(current))
		copy(fullPath, current)
		*res = append(*res, fullPath)
		return
	}
	if n.left != nil {
		findAllPathsUtil(n.left, current, res)
	}
	if n.right != nil {
		findAllPathsUtil(n.right, current, res)
	}
}

/*
NumberOfElement
find total number of nodes in bt
*/
func numOfElements(n *node) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	totalLeft := numOfElements(n.left)
	totalRight := numOfElements(n.right)
	return totalLeft + totalRight + 1
}

/*
SumOfAllNodes
Given a BT find the sum value of all it's nodes values
*/
func sumOfElements(n *node) int {
	if n == nil {
		return 0
	}
	sumOfLeft := sumOfElements(n.left)
	sumOfRight := sumOfElements(n.right)
	return sumOfLeft + sumOfRight + n.value
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

	fmt.Print("\nPrintSpiralTree:\n")
	printSpiralTree(btFromArrRecursiveOutput)

	fmt.Print("\nprintMostLeft:\n")
	printMostLeft(btFromArrRecursiveOutput)

	fmt.Print("\nprintMostLeftRec:\n")
	printMostLeftRec(btFromArrRecursiveOutput)

	fmt.Print("\ncountUniValueTrees: \n")
	countUniValueTreesOutput := countUniValueTrees(btFromArrRecursive([]int{5, 2, 7, 2, 2}))
	fmt.Printf("\ncountUniValueTrees:\nInput: %v\nOutput: %v\n", btFromArrRecursiveOutput, countUniValueTreesOutput)

	compareLeafNodesInput1 := btFromArrRecursive([]int{5, 2, 7, 2, 2})
	compareLeafNodesInput2 := btFromArrRecursive([]int{5, 3, 7, 2, 2})
	compareLeafNodesOutput := compareLeafNodes(compareLeafNodesInput1, compareLeafNodesInput2)
	fmt.Printf("\ncompareLeafNodes:\nOutput: %v\n", compareLeafNodesOutput)

	bst := &node{5, nil, nil}
	bst.left = &node{3, nil, nil}
	bst.right = &node{7, nil, nil}
	bst.left.left = &node{2, nil, nil}
	bst.left.right = &node{4, nil, nil}

	head, tail := bstToLinkedList(bst)
	fmt.Printf("\nbstToLinkedList:\nHead: %v\nTail: %v\n", (*head), (*tail))
	currentNode := head
	for currentNode != nil {
		if currentNode == head {
			fmt.Println("\nHead: ", currentNode.value)
		} else if currentNode == tail {
			fmt.Println("Tail: ", currentNode.value)
		} else {
			fmt.Println("Next: ", currentNode.value)
		}
		currentNode = currentNode.next
	}

	nthPreOrderInput := btFromArrRecursiveOutput
	nthPreOrderOutput := nthPreOrder(nthPreOrderInput, 5)
	fmt.Printf("\nNthPreOrder: \nInput: %v\nOutput: %v\n", nthPreOrderInput, nthPreOrderOutput)

	nthPostOrderInput := btFromArrRecursiveOutput
	nthPostOrderOutput := nthPostOrder(nthPostOrderInput, 5)
	fmt.Printf("\nNthPostOrder: \nInput: %v\nOutput: %v\n", nthPostOrderInput, nthPostOrderOutput)

	nthInOrderInput := btFromArrRecursiveOutput
	nthInOrderOutput := nthInOrder(nthInOrderInput, 5)
	fmt.Printf("\nNthInOrder: \nInput: %v\nOutput: %v\n", nthInOrderInput, nthInOrderOutput)

	findAllPathsInput := btFromArrRecursiveOutput
	findAllPathsOutput := findAllPaths(findAllPathsInput)
	fmt.Printf("\nFindAllPaths:\nInput: %v\nOutput: %#v\n", findAllPathsInput, findAllPathsOutput)

	numOfElementsInput := btFromArrRecursiveOutput
	numOfElementsOutput := numOfElements(numOfElementsInput)
	fmt.Printf("\nnumOfElements: \nInput: %v\nOutput: %d\n", numOfElementsInput, numOfElementsOutput)

	sumOfElementsInput := btFromArrRecursiveOutput
	sumOfElementsOutput := sumOfElements(sumOfElementsInput)
	fmt.Printf("\nsumOfElements: \nInput: %v\nOutput: %d\n", sumOfElementsInput, sumOfElementsOutput)

}
