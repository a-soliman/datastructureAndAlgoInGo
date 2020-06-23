package bst

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
func NewFromSortedSlice([]int) *BST {}

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

// Delete deletes the first found node with the given value. returns false if no node found with the value
func (bst *BST) Delete(value int) bool {}

// Find returns a node with its value equal to a given value
func (bst *BST) Find(value int) *BST {}

// Contains returns bool represting whether a node exists with the given value
func (bst *BST) Contains(value int) bool {}

// FindMin returns a node with the min value in the tree
func (bst *BST) FindMin() *BST {}

// FindMax returns a node with the max value in the tree
func (bst *BST) FindMax() *BST {}

// FindValuesInRange return []int containing all the found values within the range
func (bst *BST) FindValuesInRange(min, max int) []int {}

// Trim given a range, it deletes all the node that are out of the range, and returns the root
func (bst *BST) Trim(min, max int) *BST {}
