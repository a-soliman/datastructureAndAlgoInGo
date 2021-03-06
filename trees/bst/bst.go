package bst

import "math"

// BST node
type BST struct {
	value int
	left  *BST
	right *BST
}

// New returns a pointer to a new bst node
func New(value int) *BST {
	return &BST{value, nil, nil}
}

// NewFromUnsortedSlice returns a pointer to the root of a BST built out of a given unsorted []int
func NewFromUnsortedSlice(input []int) *BST {
	size := len(input)
	if size == 0 {
		return nil
	}
	root := New(input[0])
	for i, val := range input {
		if i > 0 {
			root.Insert(val)
		}
	}
	return root
}

// NewFromSortedSlice return a pointer to the root of a BST build out of a given sorted []int
func NewFromSortedSlice(input []int) *BST {
	size := len(input)
	midIdx := size / 2
	root := New(input[midIdx])
	treeFromSortedSliceUtil(root, &input, 0, midIdx-1, midIdx+1, size-1)
	return root
}

func treeFromSortedSliceUtil(root *BST, input *[]int, leftStartIdx int, leftEndIdx int, rightStartIdx int, rightEndIdx int) {
	// find left child
	if leftStartIdx <= leftEndIdx {
		midIdx := (leftStartIdx + leftEndIdx) / 2
		root.left = New((*input)[midIdx])
		treeFromSortedSliceUtil(root.left, input, leftStartIdx, midIdx-1, midIdx+1, leftEndIdx)
	}
	if rightStartIdx <= rightEndIdx {
		midIdx := (rightStartIdx + rightEndIdx) / 2
		root.right = New((*input)[midIdx])
		treeFromSortedSliceUtil(root.right, input, rightStartIdx, midIdx-1, midIdx+1, rightEndIdx)
	}
}

// Copy copies a tree and returns a pointer to the new trees root
func (bst *BST) Copy() *BST {
	root := New(bst.value)
	if bst.left != nil {
		root.left = bst.left.Copy()
	}
	if bst.right != nil {
		root.right = bst.right.Copy()
	}
	return root
}

// CopyMirror mirrors a tree into a new tree
func (bst *BST) CopyMirror() *BST {
	root := New(bst.value)
	if bst.left != nil {
		root.right = bst.left.CopyMirror()
	}
	if bst.right != nil {
		root.left = bst.right.CopyMirror()
	}
	return root
}

// Insert inserts a new bst node with the given value
func (bst *BST) Insert(value int) *BST {
	if value <= bst.value {
		if bst.left == nil {
			bst.left = New(value)
		} else {
			bst.left.Insert(value)
		}
	} else {
		if bst.right == nil {
			bst.right = New(value)
		} else {
			bst.right.Insert(value)
		}
	}
	return bst
}

// Find returns a node with its value equal to a given value
func (bst *BST) Find(value int) *BST {
	if value == bst.value {
		return bst
	} else if value < bst.value {
		if bst.left == nil {
			return nil
		}
		return bst.left.Find(value)
	} else {
		if bst.right == nil {
			return nil
		}
		return bst.right.Find(value)
	}
}

// Contains returns bool represting whether a node exists with the given value
func (bst *BST) Contains(value int) bool {
	if value == bst.value {
		return true
	} else if value < bst.value {
		if bst.left == nil {
			return false
		}
		return bst.left.Contains(value)
	} else {
		if bst.right == nil {
			return false
		}
		return bst.right.Contains(value)
	}
}

// FindMin returns a node with the min value in the tree
func (bst *BST) FindMin() *BST {
	if bst.left != nil {
		return bst.left.FindMin()
	}
	return bst
}

// FindMax returns the value of the  node with the max value in the tree
func (bst *BST) FindMax() *BST {
	if bst.right != nil {
		return bst.right.FindMax()
	}
	return bst
}

// IsValid returns true if a valid BST or else false
func (bst *BST) IsValid() bool {
	var checkNodeValid func(*BST, int, int) bool
	checkNodeValid = func(node *BST, min int, max int) bool {
		var valid bool
		valid = node.value >= min && node.value <= max
		if !valid {
			return false
		}
		if node.left != nil {
			valid = checkNodeValid(node.left, min, node.value)
			if !valid {
				return false
			}
		}
		if node.right != nil {
			valid = checkNodeValid(node.right, node.value, max)
			if !valid {
				return false
			}
		}
		return true
	}

	return checkNodeValid(bst, -math.MaxInt64, math.MaxInt64)
}

// InOrderTraversal return an int slice containing the tree values in order
func (bst *BST) InOrderTraversal() []int {
	res := []int{}
	inOrderTraversalUtil(bst, &res)
	return res
}

func inOrderTraversalUtil(node *BST, res *[]int) {
	if node.left != nil {
		inOrderTraversalUtil(node.left, res)
	}
	*res = append(*res, node.value)
	if node.right != nil {
		inOrderTraversalUtil(node.right, res)
	}
}

// Delete deletes the first found node with the given value. returns false if no node found with the value
func (bst *BST) Delete(value int) {
	bst = deleteNodeUtil(bst, value)
}

// Lca returns an int, and true if found the least common ansistor, or 0 and false if not found
func (bst *BST) Lca(first int, second int) (int, bool) {
	return lcaUtil(bst, first, second)
}

func lcaUtil(node *BST, first int, second int) (int, bool) {
	if node == nil {
		return 0, false
	}
	if node.value > first && node.value > second {
		return lcaUtil(node.left, first, second)
	}
	if node.value < first && node.value < second {
		return lcaUtil(node.right, first, second)
	}
	return node.value, true
}

func deleteNodeUtil(node *BST, value int) *BST {
	if node == nil {
		return nil
	}
	// var tempNode *BST = nil
	if node.value == value {
		if node.left == nil && node.right == nil {
			return nil
		}
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		maxNode := node.left.FindMax()
		maxValue := maxNode.value
		node.value = maxValue
		node.left = deleteNodeUtil(node.left, maxValue)
	} else if node.value > value {
		node.left = deleteNodeUtil(node.left, value)
	} else {
		node.right = deleteNodeUtil(node.right, value)
	}
	return node
}

// Trim given a range, it deletes all the node that are out of the range, and returns the root
func (bst *BST) Trim(min, max int) *BST {
	return trimUtil(bst, min, max)
}

func trimUtil(node *BST, min int, max int) *BST {
	if node == nil {
		return nil
	}
	node.left = trimUtil(node.left, min, max)
	node.right = trimUtil(node.right, min, max)

	if node.value < min {
		return node.right
	}
	if node.value > max {
		return node.left
	}

	return node
}

/*
Merge Two Balanced Binary Search Trees
You are given two balanced binary search trees e.g., AVL or Red Black Tree. Write a function that merges the two given balanced BSTs into a balanced binary search tree. Let there be m elements in first tree and n elements in the other tree. Your merge function should take O(m+n) time.

In the following solutions, it is assumed that sizes of trees are also given as input. If the size is not given, then we can get the size by traversing the tree
*/
func MergeTwoBST(n1 *BST, n2 *BST) *BST {
	treeOneSortedNodes, treeTwoSortedNodes := inOrderNodeTraversal(n1), inOrderNodeTraversal(n2)
	newTreeNodes := mergeNodes(treeOneSortedNodes, treeTwoSortedNodes)
	return bstFromSortedNodes(newTreeNodes, 0, len(newTreeNodes))
}

func inOrderNodeTraversal(n *BST) []*BST {
	res := []*BST{}
	inOrderNodeTraversalUtil(n, &res)
	return res
}

func inOrderNodeTraversalUtil(n *BST, res *[]*BST) {
	if n == nil {
		return
	}
	inOrderNodeTraversalUtil(n.left, res)
	*res = append(*res, n)
	inOrderNodeTraversalUtil(n.right, res)
}

func mergeNodes(nodeList1 []*BST, nodeList2 []*BST) []*BST {
	res := []*BST{}
	i, j := 0, 0
	size1, size2 := len(nodeList1), len(nodeList2)
	for i < size1 && j < size2 {
		if nodeList1[i].value < nodeList2[j].value {
			res = append(res, nodeList1[i])
			i++
		} else {
			res = append(res, nodeList2[j])
			j++
		}
	}
	for i < size1 {
		res = append(res, nodeList1[i])
		i++
	}
	for j < size2 {
		res = append(res, nodeList2[j])
		j++
	}
	return res
}

func bstFromSortedNodes(nodes []*BST, start int, end int) *BST {
	if start >= end {
		return nil
	}
	mid := (start + end) / 2
	root := nodes[mid]
	root.left = bstFromSortedNodes(nodes, start, mid)
	root.right = bstFromSortedNodes(nodes, mid+1, end)
	return root
}

// KthSmallestInBST returns the kth smallest element
func KthSmallestInBST(root *BST, k int) int {
	var res int = 0
	var idx int = 0
	kthSmallestUtil(root, &idx, k, &res)
	return res
}

func kthSmallestUtil(node *BST, idx *int, target int, res *int) {
	if *idx > target {
		return
	}
	if node.left != nil {
		kthSmallestUtil(node.left, idx, target, res)
	}
	*idx = *idx + 1
	if *idx == target {
		*res = node.value
		return
	}

	if node.right != nil {
		kthSmallestUtil(node.right, idx, target, res)
	}
}
