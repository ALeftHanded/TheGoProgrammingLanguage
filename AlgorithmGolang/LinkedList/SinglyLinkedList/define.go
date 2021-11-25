package SinglyLinkedList

// Node for SinglyLinkedList.
type Node struct {
	Val  interface{}
	Next *Node
}

// SinglyLinkedList structure with length of the list and its head
type SinglyLinkedList struct {
	length int
	Head   *Node
}
