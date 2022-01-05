package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "normal test",
			args: args{lists: []*ListNode{
				ArrayToSinglyLinkedList([]interface{}{1, 2, 3}).Head,
				ArrayToSinglyLinkedList([]interface{}{4}).Head,
				ArrayToSinglyLinkedList([]interface{}{5, 6}).Head,
				ArrayToSinglyLinkedList([]interface{}{0, 7}).Head,
				ArrayToSinglyLinkedList([]interface{}{-1, 8}).Head,
			}},
			want: ArrayToSinglyLinkedList([]interface{}{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}).Head,
		},
		{
			name: "empty test",
			args: args{lists: []*ListNode{
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
			}},
			want: ArrayToSinglyLinkedList([]interface{}{}).Head,
		},
		{
			name: "nil test",
			args: args{lists: []*ListNode{
				nil,
				nil,
			}},
			want: nil,
		},
		{
			name: "empty and nil test",
			args: args{lists: []*ListNode{
				nil,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
			}},
			want: ArrayToSinglyLinkedList([]interface{}{}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
