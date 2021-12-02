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