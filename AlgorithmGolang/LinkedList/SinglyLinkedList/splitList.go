package SinglyLinkedList

// SplitIntoTwoFromTheMiddle Split Node into two nodes from the middle of the list.
func SplitIntoTwoFromTheMiddle(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	middle := GetMiddleNode(head)
	tail := middle.Next
	middle.Next = nil
	return head, tail
}

// SplitIntoTwoByParity Split Node into two nodes according to the parity of the list.
func SplitIntoTwoByParity(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	odd, even := head, head.Next
	if even == nil {
		return odd, even
	}
	oddTmp, evenTmp := odd, even
	// if length of list is even, then end condition will be oddTmp.Next.Next == nil && evenTmp.Next == nil
	// if length of list is odd, then end condition will be evenTmp.Next.Next == nil && oddTmp.Next.Next == endNode
	for oddTmp.Next.Next != nil && evenTmp.Next.Next != nil {
		oddCur, evenCur := oddTmp.Next.Next, evenTmp.Next.Next
		oddTmp.Next, evenTmp.Next = oddCur, evenCur
		oddTmp, evenTmp = oddCur, evenCur
	}
	oddTmp.Next = oddTmp.Next.Next
	evenTmp.Next = nil

	return odd, even
}

// SplitIntoGroupByK Split Node into groups which the length of group is k.
func SplitIntoGroupByK(head *ListNode, k int) []*ListNode {
	var res []*ListNode
	length := Count(head)
	if length <= k {
		return append(res, head)
	}
	for i := 0; i < length/k; i++ {
		cur, tmpSplit := head, head
		for j := 0; j < k-1; j++ {
			cur = cur.Next
		}
		head = cur.Next
		cur.Next = nil
		res = append(res, tmpSplit)
	}
	res = append(res, head)
	return res
}
