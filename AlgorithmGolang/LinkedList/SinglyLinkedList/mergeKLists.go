package SinglyLinkedList

import "container/heap"

// MergeKLists 给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
func MergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	default:
		mid := len(lists) >> 1
		return MergeTwoLists(MergeKLists(lists[:mid]), MergeKLists(lists[mid:]))
	}
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].Val.(int) < h[j].Val.(int)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MergeKListsWithMinHeapOp(minHeap []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)
	for _, node := range minHeap {
		if node != nil {
			heap.Push(h, node)
		}
	}

	var res, current *ListNode

	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		if res == nil {
			res = minNode
			current = res
		} else {
			current.Next = minNode
			current = current.Next
		}

		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}
	return res
}

func MergeKListsWithMinHeap(minHeap []*ListNode) *ListNode {
	// remove nil in min heap array
	for i, node := range minHeap {
		if node == nil {
			minHeap = append(minHeap[:i], minHeap[i+1:]...)
		}
	}

	var res, h *ListNode
	// heapify
	heapify(minHeap)
	res, minHeap = heapPop0(minHeap)
	h = res
	minHeap = heapPush(minHeap, h.Next)

	for len(minHeap) > 0 {
		var tmpNodes *ListNode
		tmpNodes, minHeap = heapPop0(minHeap)
		h.Next = tmpNodes
		h = h.Next
		if tmpNodes.Next != nil {
			minHeap = heapPush(minHeap, tmpNodes.Next)
		}
	}
	return res
}

func heapify(minHeap MinHeap) {
	lastNonLeafNodeIndex := len(minHeap)/2 - 1
	for i := lastNonLeafNodeIndex; i >= 0; i-- {
		checkAndSwap(minHeap, i)
	}
}

func checkAndSwap(heap MinHeap, index int) {
	lastNonLeafNodeIndex := len(heap)/2 - 1
	if index > lastNonLeafNodeIndex {
		return
	}
	minIndex := index
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	if heap.Less(leftIndex, index) {
		minIndex = leftIndex
	}
	if rightIndex < len(heap) {
		if heap.Less(rightIndex, leftIndex) {
			minIndex = rightIndex
		}
	}
	if minIndex != index {
		heap.Swap(index, minIndex)
		checkAndSwap(heap, minIndex)
	}
}
func heapPop0(minHeap MinHeap) (*ListNode, []*ListNode) {

	minHeap.Swap(0, len(minHeap)-1)
	lastElementIndex := 0
	leftIndex := 2*0 + 1
	rightIndex := 2*0 + 2
	for {
		if leftIndex < len(minHeap)-1 {
			if rightIndex < len(minHeap)-1 {
				if minHeap.Less(leftIndex, rightIndex) {
					minHeap.Swap(leftIndex, lastElementIndex)
					lastElementIndex = leftIndex
				} else {
					minHeap.Swap(rightIndex, lastElementIndex)
					lastElementIndex = rightIndex
				}
			} else {
				minHeap.Swap(leftIndex, lastElementIndex)
				lastElementIndex = leftIndex
			}
		} else {
			break
		}
		leftIndex = 2*lastElementIndex + 1
		rightIndex = 2*lastElementIndex + 2
	}
	return minHeap[len(minHeap)-1], minHeap[:len(minHeap)-1]
}

func heapPush(minHeap MinHeap, pushNode *ListNode) []*ListNode {
	if pushNode == nil {
		return minHeap
	}
	minHeap = append(minHeap, pushNode)
	pushNodeIndex := len(minHeap) - 1
	parentIndex := (pushNodeIndex - 1) / 2
	for parentIndex >= 0 {
		if minHeap.Less(pushNodeIndex, parentIndex) {
			minHeap.Swap(pushNodeIndex, parentIndex)
		} else {
			break
		}
		pushNodeIndex = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
	return minHeap
}
