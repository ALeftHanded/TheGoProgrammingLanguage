package SinglyLinkedList

func ReverseKGroup(head *ListNode, k int) *ListNode {
	length := Count(head)
	splitKGroup := SplitIntoGroupByK(head, k)
	var res *ListNode
	if length%k != 0 {
		res = splitKGroup[len(splitKGroup)-1]
		splitKGroup = splitKGroup[:len(splitKGroup)-1]
	}
	for i := len(splitKGroup) - 1; i >= 0; i-- {
		res = AddListNodeAtEnd(ReverseList(splitKGroup[i]), res)
	}
	return res
}

func ReverseKGroupV2(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		// when length is not divisible by k, tail will be nil
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	pre := ReverseKGroupV2(tail, k)
	for i := 0; i < k; i++ {
		cur := head
		head = head.Next
		cur.Next = pre
		pre = cur
	}
	return pre
}
