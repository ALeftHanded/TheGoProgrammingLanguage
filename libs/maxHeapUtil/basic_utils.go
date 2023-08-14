package maxHeapUtil

func Add(maxHeap *MaxHeap, num int) *MaxHeap {
	if maxHeap == nil {
		return &MaxHeap{Val: num}
	}
	// insert to the last node
	tmpNode := &MaxHeap{Val: num}
	maxHeap.insertLastNode(tmpNode)

	// adjust the heap if val of the parent node is less than val of the child node
	var parentNode *MaxHeap
	for {
		parentNode = findParentNode(maxHeap, tmpNode)
		// if parentNode == nil, it means that tmpNode is the root node
		if parentNode == nil {
			break
		}
		if tmpNode.Val <= parentNode.Val {
			break
		}
		tmpNode.Val, parentNode.Val = parentNode.Val, tmpNode.Val
		tmpNode = parentNode
	}
	return maxHeap

}

func (maxh *MaxHeap) insertLastNode(lastNode *MaxHeap) {
	queue := []*MaxHeap{maxh}
	var firstNode *MaxHeap
	for {
		firstNode = queue[0]
		if firstNode.Left == nil {
			firstNode.Left = lastNode
			break
		}
		if firstNode.Right == nil {
			firstNode.Right = lastNode
			break
		}
		queue = append(queue, firstNode.Left, firstNode.Right)
		queue = queue[1:]
	}
}

// mh is in root
func findParentNode(root, mh *MaxHeap) *MaxHeap {
	if root == nil || root == mh {
		return nil
	}
	queue := []*MaxHeap{root}
	for {
		firstNode := queue[0]
		if firstNode.Left == mh || firstNode.Right == mh {
			return firstNode
		}
		queue = append(queue, firstNode.Left, firstNode.Right)
		queue = queue[1:]
	}
}

func (maxh *MaxHeap) Equal(other *MaxHeap) bool {
	if maxh == nil && other == nil {
		return true
	}
	if maxh == nil || other == nil {
		return false
	}
	if maxh.Val != other.Val {
		return false
	}
	return maxh.Left.Equal(other.Left) && maxh.Right.Equal(other.Right)
}

func (maxh *MaxHeap) ValidCheck() bool {
	if maxh == nil {
		// Consider a nil heap as valid (temporary assumption)
		return true
	}
	// ! Check for the completeness of the binary tree
	if maxh.Left == nil && maxh.Right != nil {
		// Invalid case: Right child exists without a corresponding left child
		return false
	}
	if maxh.Left == nil && maxh.Right == nil {
		// Leaf node: return true as it's valid
		return true
	}
	// From this point on, a left child must exist
	if maxh.Val <= maxh.Left.Val {
		// Parent value must be greater than left child
		return false
	}
	if maxh.Right != nil && maxh.Val <= maxh.Right.Val {
		// Parent value must be greater than right child, if it exists
		return false
	}
	// ! Validate left subtree, and right subtree if it exists
	// * nil do not get ValidCheck method, which will cause panic
	return maxh.Left.ValidCheck() && (maxh.Right == nil || maxh.Right.ValidCheck())
}
