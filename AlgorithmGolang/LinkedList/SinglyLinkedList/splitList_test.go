package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestSplitIntoTwoFromTheMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name  string
		args  args
		want1 *ListNode
		want2 *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "normal test",
			args: args{
				head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}).Head,
			},
			want1: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5}).Head,
			want2: ArrayToSinglyLinkedList([]interface{}{6, 7, 8, 9}).Head,
		},
		{
			name:  "nil test",
			args:  args{nil},
			want1: nil,
			want2: nil,
		},
		{
			name:  "single node test",
			args:  args{head: ArrayToSinglyLinkedList([]interface{}{1}).Head},
			want1: ArrayToSinglyLinkedList([]interface{}{1}).Head,
			want2: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := SplitIntoTwoFromTheMiddle(tt.args.head)
			tt.args.head.Display()
			got1.Display()
			got2.Display()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SplitIntoTwoFromTheMiddle() got = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SplitIntoTwoFromTheMiddle() got1 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestSplitIntoTwoByParity(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name  string
		args  args
		want1 *ListNode
		want2 *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "normal test",
			args: args{
				head: ArrayToSinglyLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7}).Head,
			},
			want1: ArrayToSinglyLinkedList([]interface{}{1, 3, 5, 7}).Head,
			want2: ArrayToSinglyLinkedList([]interface{}{2, 4, 6}).Head,
		},
		{
			name:  "nil test",
			args:  args{nil},
			want1: nil,
			want2: nil,
		},
		{
			name:  "single node test",
			args:  args{head: ArrayToSinglyLinkedList([]interface{}{1}).Head},
			want1: ArrayToSinglyLinkedList([]interface{}{1}).Head,
			want2: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := SplitIntoTwoByParity(tt.args.head)
			tt.args.head.Display()
			got1.Display()
			got2.Display()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SplitIntoTwoByParity() got = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SplitIntoTwoByParity() got1 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
