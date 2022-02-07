package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestMergeTwoSinglyLinkedLists(t *testing.T) {
	type args struct {
		l1 *SinglyLinkedList
		l2 *SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want *SinglyLinkedList
	}{
		{
			name: "both nil",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "both node nil",
			args: args{
				l1: NewSinglyLinkedList(),
				l2: NewSinglyLinkedList(),
			},
			want: NewSinglyLinkedList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoSinglyLinkedLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoSinglyLinkedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal test",
			args: args{
				list1: ArrayToSinglyLinkedList([]interface{}{1, 3, 5, 7, 9}).Head,
				list2: ArrayToSinglyLinkedList([]interface{}{2, 4, 6, 8}).Head,
			},
			want: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}).Head,
		},
		{
			name: "both nil",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "list1 nil",
			args: args{
				list1: nil,
				list2: ArrayToSinglyLinkedList([]interface{}{2, 4, 6, 8}).Head,
			},
			want: ArrayToSinglyLinkedList([]interface{}{2, 4, 6, 8}).Head,
		},
		{
			name: "list2 nil",
			args: args{
				list1: ArrayToSinglyLinkedList([]interface{}{1, 3, 5, 7, 9}).Head,
				list2: nil,
			},
			want: ArrayToSinglyLinkedList([]interface{}{1, 3, 5, 7, 9}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
