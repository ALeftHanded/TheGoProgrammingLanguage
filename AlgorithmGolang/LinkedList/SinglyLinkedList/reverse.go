package SinglyLinkedList


func ReverseSingleLinkedList(sll *SinglyLinkedList) *SinglyLinkedList {
	if sll == nil {
		return nil
	}
	head := ReverseList(sll.Head)
	return &SinglyLinkedList{
		Length: sll.Length,
		Head:   head,
	}
}


// ReverseList 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func ReverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = res
		res = cur
	}
	return res
}

// ReverseBetween 给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表 。
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	/* Fixme high memory usage
	headToLeft := head
	count:= 1
	// 保证head指向left侧
	for count < left {
		count++
		head = head.Next
	}
	if headToLeft == head {
		headToLeft = nil
	} else {
		cur := headToLeft
		for cur.Next != head {
			cur = cur.Next
		}
		cur.Next = nil
	}
	leftToRight := head
	// 保证head指向right侧
	for count < right {
		count++
		head = head.Next
	}
	rightToTail := head.Next
	head.Next = nil
	return AddListNodeAtEnd(
		headToLeft, AddListNodeAtEnd(
			ReverseList(leftToRight), rightToTail))
	*/
	var headToLeft *ListNode
	count:= 1
	if left > 1{
		headToLeft = head
	}
	// 保证head指向left-1侧
	for count < left - 1 {
		count++
		head = head.Next
	}

	var leftToRightReverse *ListNode
	for count <= right {
		cur := head
		head = head.Next
		cur.Next = leftToRightReverse
		leftToRightReverse = cur
		count++
	}
	// 保证head指向right+1侧
	return AddListNodeAtEnd(
		headToLeft, AddListNodeAtEnd(
			leftToRightReverse, head))
}