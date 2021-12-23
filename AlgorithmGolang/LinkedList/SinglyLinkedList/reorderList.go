package SinglyLinkedList

func ReorderList(head *ListNode) *ListNode {
	midNode := GetMiddleNode(head)
	reverseListNode := ReverseList(midNode.Next)
	// 组装前清空head尾部
	midNode.Next = nil
	res := head
	for reverseListNode != nil {
		reverseCur := reverseListNode
		headCur := head

		head = head.Next
		reverseListNode = reverseListNode.Next

		headCur.Next = reverseCur
		reverseCur.Next = head
	}
	return res
}
