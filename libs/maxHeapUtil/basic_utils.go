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

func findParentNode(root, mh *MaxHeap) *MaxHeap {
	if root == nil || root == mh {
		return nil
	}
	queue := []*MaxHeap{root}
	for {
		firstNode := queue[0]
		if firstNode.Left == mh {
			return firstNode
		}
		if firstNode.Right == mh {
			return firstNode
		}
		queue = append(queue, firstNode.Left, firstNode.Right)
		queue = queue[1:]
	}
}

func InitMaxHeapFromIntArray(arr []int) *MaxHeap {
	if arr == nil {
		return nil
	}
	var res *MaxHeap
	for _, num := range arr {
		res = Add(res, num)
	}
	return res
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
