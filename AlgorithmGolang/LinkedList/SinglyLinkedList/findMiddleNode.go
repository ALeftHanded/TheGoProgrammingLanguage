package SinglyLinkedList

// GetMiddleNode 快慢指针找中间节点
// 中间节点定义：
// k为正整数，ListNode的长度为2k+1则slow指向第k+1个节点
// k为正整数，ListNode的长度为2k则slow指向第k个节点
func GetMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
