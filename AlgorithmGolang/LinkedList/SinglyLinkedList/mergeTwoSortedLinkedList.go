package SinglyLinkedList

func MergeTwoSinglyLinkedLists(l1 *SinglyLinkedList, l2 *SinglyLinkedList) *SinglyLinkedList {
	var head1, head2 *ListNode
	var len1, len2 int
	if l1 != nil {
		head1 = l1.Head
		len1 = l1.Length
	} else if l2 != nil {
		head2 = l2.Head
		len2 = l2.Length
	} else {
		return nil
	}
	return &SinglyLinkedList{
		Head:   MergeTwoLists(head1, head2),
		Length: len1 + len2,
	}
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Discursive

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val.(int) <= list2.Val.(int) {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}

	// Iteration

	//preNode := &ListNode{Val: -1}
	//head := preNode
	//for list1 != nil || list2 != nil {
	//	if list1 == nil {
	//		head.Next = list2
	//		return preNode.Next
	//	} else if list2 == nil {
	//		head.Next = list1
	//		return preNode.Next
	//	} else if list1.Val.(int) <= list2.Val.(int) {
	//		head.Next = list1
	//		list1 = list1.Next
	//		head = head.Next
	//	} else {
	//		head.Next = list2
	//		list2 = list2.Next
	//		head = head.Next
	//	}
	//}
	//return nil
}
