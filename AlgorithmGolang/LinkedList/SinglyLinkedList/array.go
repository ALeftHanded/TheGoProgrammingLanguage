package SinglyLinkedList

func ArrayToSinglyLinkedList(arr []interface{}) *SinglyLinkedList {
	sll := NewSinglyLinkedList()
	if arr == nil {
		return nil
	} else if len(arr) == 0 {
		return sll
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			sll.AddAtBegin(arr[i])
		}
	}
	return sll
}

func SinglyLinkedListToArray(sll *SinglyLinkedList) []interface{} {
	arr := []interface{}{}
	if sll == nil {
		return nil
	} else if sll.Head == nil {
		return arr
	} else {
		head := sll.Head
		for head != nil {
			arr = append(arr, head.Val)
			head = head.Next
		}
	}
	return arr
}
