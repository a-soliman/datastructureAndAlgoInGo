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

// Delete deletes the first found node with the given value. returns false if no node found with the value
func (bst *BST) Delete(value int) bool {}

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

// FindMax returns a node with the max value in the tree
func (bst *BST) FindMax() *BST {
	if bst.right != nil {
		return bst.right.FindMax()
	}
	return bst
}

// FindValuesInRange return []int containing all the found values within the range
func (bst *BST) FindValuesInRange(min, max int) []int {}

// Trim given a range, it deletes all the node that are out of the range, and returns the root
func (bst *BST) Trim(min, max int) *BST {}
