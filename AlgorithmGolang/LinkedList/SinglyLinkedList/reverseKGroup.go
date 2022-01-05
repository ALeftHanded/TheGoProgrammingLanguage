package SinglyLinkedList

// ReverseKGroup 给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
// k是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
func ReverseKGroup(head *ListNode, k int) *ListNode {
	length := Count(head)
	splitKGroup := SplitIntoGroupByK(head, k)
	var res *ListNode
	if length%k != 0 {
		res = splitKGroup[len(splitKGroup)-1]
		splitKGroup = splitKGroup[:len(splitKGroup)-1]
	}
	for i := len(splitKGroup) - 1; i >= 0; i-- {
		res = AddListNodeAtEnd(ReverseList(splitKGroup[i]), res)
	}
	return res
}

func ReverseKGroupV2(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		// when length is not divisible by k, tail will be nil
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	pre := ReverseKGroupV2(tail, k)
	for i := 0; i < k; i++ {
		cur := head
		head = head.Next
		cur.Next = pre
		pre = cur
	}
	return pre
}
