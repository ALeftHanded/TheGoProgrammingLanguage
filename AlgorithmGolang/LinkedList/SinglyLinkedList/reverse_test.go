package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestReverseSingleLinkedList(t *testing.T) {
	type args struct {
		sll *SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want *SinglyLinkedList
	}{
		{
			name: "normal reverse test",
			args: args{
				sll: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}),
			},
			want: ArrayToSinglyLinkedList([]interface{}{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}),
		},
		{
			name: "nil reverse test",
			args: args{
				sll: nil,
			},
			want: nil,
		},
		{
			name: "empty reverse test",
			args: args{
				sll: NewSinglyLinkedList(),
			},
			want: NewSinglyLinkedList(),
		},
		{
			name: "single node SinglyLinkedList reverse test",
			args: args{
				sll: ArrayToSinglyLinkedList([]interface{}{0}),
			},
			want: ArrayToSinglyLinkedList([]interface{}{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSingleLinkedList(tt.args.sll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSingleLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal reverse test",
			args: args{
				head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}).Head,
			},
			want: ArrayToSinglyLinkedList([]interface{}{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}).Head,
		},
		{
			name: "nil reverse test",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "empty reverse test",
			args: args{
				head: NewListNode(),
			},
			want: NewListNode(),
		},
		{
			name: "single node SinglyLinkedList reverse test",
			args: args{
				head: ArrayToSinglyLinkedList([]interface{}{0}).Head,
			},
			want: ArrayToSinglyLinkedList([]interface{}{0}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal test",
			args: args{
				head:  ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}).Head,
				left:  3,
				right: 8,
			},
			want: ArrayToSinglyLinkedList([]interface{}{1, 2, 8, 7, 6, 5, 4, 3, 9}).Head,
		},
		{
			name: "single test",
			args: args{
				head:  ArrayToSinglyLinkedList([]interface{}{1}).Head,
				left:  1,
				right: 1,
			},
			want: ArrayToSinglyLinkedList([]interface{}{1}).Head,
		},
		{
			name: "left=0 right=length(listnode) test",
			args: args{
				head:  ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7}).Head,
				left:  1,
				right: 7,
			},
			want: ArrayToSinglyLinkedList([]interface{}{7, 6, 5, 4, 3, 2, 1}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
