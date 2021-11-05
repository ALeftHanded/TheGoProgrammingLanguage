package SinglyLinkedList

type Node struct {
	Val  int
	Next *Node
}

func NewSingly() *Node {
	return &Node{}
}

func ArrayToSinglyLinkedList(array []int) *Node {
	head := NewSingly()
	if len(array) == 0 {
		return head
	} else {
		for _, val := range array{

		}
	}

	return head
}