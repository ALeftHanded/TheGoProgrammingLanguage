package SinglyLinkedList

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	type args struct {
		headA      *ListNode
		headB      *ListNode
		commonPart *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal test",
			args: args{
				headA:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6}).Head,
				headB:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6, 1, 234, 35}).Head,
				commonPart: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4}).Head,
			},
		},
		{
			name: "same value test",
			args: args{
				headA:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6}).Head,
				headB:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6}).Head,
				commonPart: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4}).Head,
			},
		},
		{
			name: "no common part test",
			args: args{
				headA:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6}).Head,
				headB:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6}).Head,
				commonPart: nil,
			},
		},
		{
			name: "one head nil test",
			args: args{
				headA:      nil,
				headB:      ArrayToSinglyLinkedList([]interface{}{9, 4, 6}).Head,
				commonPart: ArrayToSinglyLinkedList([]interface{}{9}).Head,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.headA != nil {
				tt.args.headA.AddListNodeAtEnd(tt.args.commonPart)
			} else {
				tt.args.headA = tt.args.commonPart
			}
			if tt.args.headB != nil {
				tt.args.headB.Next.AddListNodeAtEnd(tt.args.commonPart)
			} else {
				tt.args.headB = tt.args.commonPart
			}
			if got := GetIntersectionNode(tt.args.headA, tt.args.headB); got != tt.args.commonPart {
				t.Errorf("GetIntersectionNode() = %v, want %v", got, tt.args.commonPart)
			}
		})
	}
}
