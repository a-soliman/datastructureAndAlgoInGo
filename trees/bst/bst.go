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
