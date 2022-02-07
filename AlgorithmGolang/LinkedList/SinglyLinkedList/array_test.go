package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestArrayToSinglyLinkedList(t *testing.T) {
	type args struct {
		arr []interface{}
	}
	tests := []struct {
		name string
		args args
		want *SinglyLinkedList
	}{
		{
			name: "nil arr test",
			args: args{
				nil,
			},
			want: nil,
		},
		{
			name: "empty arr test",
			args: args{
				[]interface{}{},
			},
			want: &SinglyLinkedList{},
		},
		{
			name: "normal arr test",
			args: args{
				[]interface{}{1, 2, 3},
			},
			want: &SinglyLinkedList{
				Length: 3,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToSinglyLinkedList(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToSinglyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedListToArray(t *testing.T) {
	type args struct {
		sll *SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "normal SinglyLinkedList test",
			args: args{
				sll: &SinglyLinkedList{
					Length: 3, Head: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 3}},
					},
				},
			},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "nil SinglyLinkedList test",
			args: args{nil},
			want: nil,
		},
		{
			name: "empty SinglyLinkedList test",
			args: args{sll: &SinglyLinkedList{}},
			want: []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinglyLinkedListToArray(tt.args.sll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SinglyLinkedListToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
