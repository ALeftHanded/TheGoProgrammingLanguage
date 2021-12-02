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
		// TODO: Add test cases.
		{
			name: "normal reverse test",
			args: args{
				sll: ArrayToSinglyLinkedList([]interface{}{1,2,3,4,5,6,7,8,9,0}),
			},
			want: ArrayToSinglyLinkedList([]interface{}{0,9,8,7,6,5,4,3,2,1}),
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