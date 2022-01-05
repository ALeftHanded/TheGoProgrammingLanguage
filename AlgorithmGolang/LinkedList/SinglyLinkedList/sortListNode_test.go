package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestSortListNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "normal test",
			args: args{head: ArrayToSinglyLinkedList([]interface{}{5, 4, 3, 2, 1}).Head},
			want: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortListNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
