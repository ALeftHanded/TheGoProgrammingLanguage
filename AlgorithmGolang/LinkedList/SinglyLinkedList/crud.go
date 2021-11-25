package SinglyLinkedList

import "fmt"

// Create: new && insert

// NewSinglyLinkedList init SinglyLinkedList
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// AddAtBegin add new SinglyLinkedList node at the beginning of the list
func (sll *SinglyLinkedList) AddAtBegin(val interface{}) {
	sll.length++
	node := &Node{Val: val}
	node.Next = sll.Head
	sll.Head = node
}

// AddAtEnd add new SinglyLinkedList node at the end of the list
func (sll *SinglyLinkedList) AddAtEnd(val interface{}) {
	sll.length++
	node := &Node{Val: val}
	head := sll.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
}

// Read

// Count return the length of the SinglyLinkedList
func (sll *SinglyLinkedList) Count() int {
	return sll.length
}

// Display prints out the elements of the SinglyLinkedList.
func (sll *SinglyLinkedList) Display() {
	head := sll.Head
	for head != nil {
		fmt.Printf("%v", head.Val)
		head = head.Next
	}
}

// Update

// Delete

// DelAtBegin remove SinglyLinkedList node at the beginning of the list. Returns error if the list is empty.
func (sll *SinglyLinkedList) DelAtBegin() error {
	if sll.length == 0 {
		return fmt.Errorf("SinglyLinkedList is empty, cannot delete")
	}
	sll.Head = sll.Head.Next
	sll.length--
	return nil
}

// DelAtEnd remove SinglyLinkedList node at the end of the list. Returns error if the list is empty.
func (sll *SinglyLinkedList) DelAtEnd() error {
	if sll.length == 0 {
		return fmt.Errorf("SinglyLinkedList is empty, cannot delete")
	}
	head := sll.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	head.Next = nil
	sll.length--
	return nil
}

// PopAtBegin remove SinglyLinkedList node at the beginning of the list and return its value. Returns -1 if the list is empty.
func (sll *SinglyLinkedList) PopAtBegin() interface{} {
	if sll.length == 0 {
		return -1
	}
	res := sll.Head.Val
	sll.Head = sll.Head.Next
	sll.length--
	return res
}

// PopAtEnd remove SinglyLinkedList node at the end of the list and return its value. Returns -1 if the list is empty.
func (sll *SinglyLinkedList) PopAtEnd() interface{} {
	if sll.length == 0 {
		return -1
	}
	head := sll.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	res := head.Val
	head.Next = nil
	sll.length--
	return res
}
