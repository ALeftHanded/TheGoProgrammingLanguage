package SinglyLinkedList

// ListNode for SinglyLinkedList.
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// SinglyLinkedList structure with length of the list and its head
type SinglyLinkedList struct {
	length int
	Head   *ListNode
}
