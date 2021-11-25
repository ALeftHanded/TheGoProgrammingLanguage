package SinglyLinkedList

func ArrayToSinglyLinkedList(arr []interface{}) *Node {
	var head *Node
	if len(arr) == 0 {
		return nil
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			node := &Node{
				Val:  arr[i],
				Next: head,
			}
			head = node
		}
	}
	return head
}

func SinglyLinkedListToArray(head *Node) []interface{} {
	arr := []interface{}{}
	if head == nil {
		return arr
	} else {
		for head != nil {
			arr = append(arr, head.Val)
			head = head.Next
		}
	}
	return arr
}
