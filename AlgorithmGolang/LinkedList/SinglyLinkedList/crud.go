package SinglyLinkedList

import "fmt"

// Create: new && insert

// NewSinglyLinkedList init SinglyLinkedList
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// NewListNode init ListNode
func NewListNode() *ListNode {
	return &ListNode{}
}

// AddAtBegin add new SinglyLinkedList node at the beginning of the list
func (sll *SinglyLinkedList) AddAtBegin(val interface{}) {
	sll.Length++
	node := &ListNode{Val: val}
	node.Next = sll.Head
	sll.Head = node
}

// AddAtEnd add new SinglyLinkedList node at the end of the list
// todo Need rewrite this func.
func (sll *SinglyLinkedList) AddAtEnd(val interface{}) {
	sll.Length++
	node := &ListNode{Val: val}
	// Head is nil
	if sll.Head == nil {
		sll.Head = node
		return
	}

	head := sll.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
}

// Read

// Count return the length of the SinglyLinkedList
func (sll *SinglyLinkedList) Count() int {
	return sll.Length
}

// Display prints out the elements of the SinglyLinkedList.
func (sll *SinglyLinkedList) Display() {
	head := sll.Head
	fmt.Printf("----- Start display -----\n")
	if head == nil {
		fmt.Printf("----- SinglyLinkedList is empty! -----\n")
	}
	for head != nil {
		fmt.Printf("%v\n", head.Val)
		head = head.Next
	}
	fmt.Printf("----- End display -----\n")
}

// Update

// Delete

// DelAtBegin remove SinglyLinkedList node at the beginning of the list. Returns error if the list is empty.
func (sll *SinglyLinkedList) DelAtBegin() error {
	if sll.Length == 0 {
		return fmt.Errorf("SinglyLinkedList is empty, cannot delete")
	}
	sll.Head = sll.Head.Next
	sll.Length--
	return nil
}

// DelAtEnd remove SinglyLinkedList node at the end of the list. Returns error if the list is empty.
// todo Need rewrite this func.
func (sll *SinglyLinkedList) DelAtEnd() error {
	if sll.Length == 0 {
		return fmt.Errorf("SinglyLinkedList is empty, cannot delete")
	}
	sll.Length--
	// single node delete
	if sll.Head.Next == nil {
		sll.Head = nil
		return nil
	}

	head := sll.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	head.Next = nil
	return nil
}

// PopAtBegin remove SinglyLinkedList node at the beginning of the list and return its value. Returns -1 if the list is empty.
func (sll *SinglyLinkedList) PopAtBegin() interface{} {
	if sll.Length == 0 {
		return -1
	}
	res := sll.Head.Val
	sll.Head = sll.Head.Next
	sll.Length--
	return res
}

// PopAtEnd remove SinglyLinkedList node at the end of the list and return its value. Returns -1 if the list is empty.
func (sll *SinglyLinkedList) PopAtEnd() interface{} {
	if sll.Length == 0 {
		return -1
	}
	sll.Length--
	if sll.Head.Next == nil {
		res := sll.Head.Val
		sll.Head = nil
		return res
	}

	head := sll.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	res := head.Next.Val
	head.Next = nil
	return res
}
