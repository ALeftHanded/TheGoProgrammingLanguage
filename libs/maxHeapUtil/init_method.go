package maxHeapUtil

func InitMaxHeapFromIntArrayBad(arr []int) *MaxHeap {
	if arr == nil {
		return nil
	}
	var res *MaxHeap
	for _, num := range arr {
		res = Add(res, num)
	}
	return res
}

func HeapifyInit(arr []int) *MaxHeap {
	if arr == nil {
		return nil
	}
	// 1. Identify last non-leaf node index
	// ? how many leaf node does a complete tree have?
	// ? why lastNonLeafNodeIndex is len(arr)/2 - 1
	// - Complete binary tree has about half nodes as leaf nodes.
	// - lastNonLeafNodeIndex is len(arr)/2 - 1 because children of len(arr)/2 will be out of range.

	lastNonLeafNodeIndex := len(arr)/2 - 1
	// 2. Heapify the array using a bottom-up approach
	for i := lastNonLeafNodeIndex; i >= 0; i-- {
		CheckAndSwap(arr, i)
	}
	// 3. Construct MaxHeap from the heapified array
	var nodeList []*MaxHeap
	for i := range arr {
		nodeList = append(nodeList, &MaxHeap{arr[i], nil, nil})
	}
	// Link parent nodes to their left and right children
	lastNonLeafNodeIndex = len(nodeList)/2 - 1
	for i := lastNonLeafNodeIndex; i >= 0; i-- {
		if i*2+2 < len(nodeList) {
			nodeList[i].Right = nodeList[2*i+2]
		}
		nodeList[i].Left = nodeList[2*i+1]
	}
	return nodeList[0]
}

func CheckAndSwap(arr []int, index int) {
	lastNonLeafNodeIndex := len(arr)/2 - 1
	if index > lastNonLeafNodeIndex {
		return
	}
	largestIndex := GetLargestIndex(arr, index)
	if largestIndex != index {
		arr[index], arr[largestIndex] = arr[largestIndex], arr[index]
		CheckAndSwap(arr, largestIndex)
	}
}

func GetLargestIndex(arr []int, index int) int {
	// heapify the array
	// ? how to find left and right children index and why they can be found
	// * if your node index is i, the i nodes before you all have two children
	// * so there are 2i nodes before your children, and your children index should be 2i+1 and 2i+2
	largestIndex, leftIndex, rightIndex := index, index*2+1, index*2+2
	if arr[leftIndex] > arr[index] {
		largestIndex = leftIndex
	}
	if rightIndex < len(arr) && arr[rightIndex] > arr[largestIndex] {
		largestIndex = rightIndex
	}
	return largestIndex
}
