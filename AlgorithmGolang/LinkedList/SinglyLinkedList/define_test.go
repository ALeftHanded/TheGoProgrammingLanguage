package SinglyLinkedList

import (
	"reflect"
	"testing"
)

func TestNewSingly(t *testing.T) {
	tests := []struct {
		name string
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "just a test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingly(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingly() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

