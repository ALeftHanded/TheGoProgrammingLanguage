package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestArrayToSinglyLinkedList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "null arr test",
			args: args{
				[]int{},
			},
			want: nil,
		},
		{
			name: "normal arr test",
			args: args{
				[]int{1,2,3},
			},
			want: &Node{
				Val: 1,
				Next: &Node{
					Val: 2,
					Next: &Node{Val: 3},
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
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "normal SinglyLinkedList test",
			args: args{
				&Node{
					Val: 1,
					Next: &Node{
						Val: 2,
					Next: &Node{Val: 3}},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "nil SinglyLinkedList test",
			args: args{nil},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinglyLinkedListToArray(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SinglyLinkedListToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}