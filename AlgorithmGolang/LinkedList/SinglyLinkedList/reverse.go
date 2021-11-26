package SinglyLinkedList

// ReverseList 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func ReverseList(sll *SinglyLinkedList) *SinglyLinkedList {
	var res *Node
	head := sll.Head
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = res
		res = cur
	}
	return &SinglyLinkedList{
		length: sll.length,
		Head: res,
	}
}