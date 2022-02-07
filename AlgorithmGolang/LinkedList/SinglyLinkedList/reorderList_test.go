package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal test",
			args: args{
				head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3}).Head,
			},
			want: ArrayToSinglyLinkedList([]interface{}{
				1, 3, 2}).Head,
		},
		{
			name: "single node test",
			args: args{
				head: ArrayToSinglyLinkedList([]interface{}{1}).Head,
			},
			want: ArrayToSinglyLinkedList([]interface{}{1}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				got.Display()
				tt.want.Display()
				t.Errorf("ReorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
