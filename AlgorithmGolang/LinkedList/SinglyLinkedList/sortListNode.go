package SinglyLinkedList

func SortListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, tail := SplitIntoTwoFromTheMiddle(head)
	newHeadSorted, tailSorted := SortListNode(newHead), SortListNode(tail)
	return MergeTwoLists(newHeadSorted, tailSorted)
}
