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

func main() {
	btFromArrIterativeInput := []int{1, 2, 3, 4, 5, 6, 7}
	btFromArrIterativeOutput := btFromArrIterative(btFromArrIterativeInput)
	fmt.Printf("\nbtFromArrIterative: \nInput:%v\nOutput: %v\n", btFromArrIterativeInput, btFromArrIterativeOutput)
}
