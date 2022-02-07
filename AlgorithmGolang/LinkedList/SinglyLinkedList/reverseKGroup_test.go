package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal test",
			args: args{head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}).Head, k: 3},
			want: ArrayToSinglyLinkedList([]interface{}{3, 2, 1, 6, 5, 4, 9, 8, 7, 10, 11}).Head,
		},
		{
			name: "length less than k",
			args: args{head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Head, k: 11},
			want: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Head,
		},
		{
			name: "k is divisible by length",
			args: args{head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}).Head, k: 3},
			want: ArrayToSinglyLinkedList([]interface{}{3, 2, 1, 6, 5, 4, 9, 8, 7}).Head,
		},
		{
			name: "k is 1",
			args: args{head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}).Head, k: 1},
			want: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
