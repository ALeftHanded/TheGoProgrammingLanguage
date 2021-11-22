package SinglyLinkedList

type Node struct {
	Val  int
	Next *Node
}

func NewSingly() *Node {
	return &Node{}
}

func ArrayToSinglyLinkedList(arr []int) *Node {
	var head *Node
	if len(arr) == 0 {
		return nil
	} else {
		// arr reversed
		for l, r := 0, len(arr)-1; l<r; l, r = l+1, r-1{
			arr[l], arr[r] = arr[r], arr[l]
		}
		for _, val := range arr {
			node := &Node{
				Val: val,
				Next: head,
			}
			head = node
		}
	}
	return head
}