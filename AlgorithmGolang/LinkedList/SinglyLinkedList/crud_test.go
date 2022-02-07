package SinglyLinkedList

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *SinglyLinkedList
	}{
		{
			name: "normal test",
			want: &SinglyLinkedList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSinglyLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSinglyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_AddAtBegin(t *testing.T) {
	type fields struct {
		length int
		Head   *ListNode
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "empty list first add",
			fields: fields{},
			args: args{
				val: 123,
			},
		},
		{
			name: "normal list first add",
			fields: fields{
				length: 3,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
			args: args{
				val: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sll := &SinglyLinkedList{
				Length: tt.fields.length,
				Head:   tt.fields.Head,
			}
			sll.Display()
			fmt.Printf("\n----- Exec AddAtBegin -----\n\n")
			sll.AddAtBegin(tt.args.val)
			sll.Display()
		})
	}
}

func TestSinglyLinkedList_AddAtEnd(t *testing.T) {
	type fields struct {
		length int
		Head   *ListNode
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "empty list add at end",
			fields: fields{},
			args: args{
				val: 123,
			},
		},
		{
			name: "normal list add at end",
			fields: fields{
				length: 3,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
			args: args{
				val: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sll := &SinglyLinkedList{
				Length: tt.fields.length,
				Head:   tt.fields.Head,
			}
			sll.Display()
			fmt.Printf("\n----- Exec AddAtEnd -----\n\n")
			sll.AddAtEnd(tt.args.val)
			sll.Display()
		})
	}
}

func TestListNode_AddListNodeAtEnd(t *testing.T) {
	type fields struct {
		Val  interface{}
		Next *ListNode
	}
	type args struct {
		tail *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "normal test",
			fields: fields{
				Val:  123,
				Next: &ListNode{Val: 1231312},
			},
			args: args{ArrayToSinglyLinkedList([]interface{}{1, 2, 34, 4, 52, 53}).Head},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			ln.Display()
			ln.AddListNodeAtEnd(tt.args.tail)
			ln.Display()
		})
	}
}

func TestSinglyLinkedList_DelAtBegin(t *testing.T) {
	type fields struct {
		length int
		Head   *ListNode
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty SinglyLinkedList del",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "single node SinglyLinkList del",
			fields: fields{
				length: 1,
				Head:   &ListNode{Val: 1},
			},
			wantErr: false,
		},
		{
			name: "normal SinglyLinkList del",
			fields: fields{
				length: 1,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sll := &SinglyLinkedList{
				Length: tt.fields.length,
				Head:   tt.fields.Head,
			}

			sll.Display()
			fmt.Printf("\n----- Exec DelAtBegin -----\n\n")
			if err := sll.DelAtBegin(); (err != nil) != tt.wantErr {
				t.Errorf("SinglyLinkedList.DelAtBegin() error = %v, wantErr %v", err, tt.wantErr)
			}
			sll.Display()
		})
	}
}

func TestSinglyLinkedList_DelAtEnd(t *testing.T) {
	type fields struct {
		length int
		Head   *ListNode
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty SinglyLinkedList del",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "single node SinglyLinkList del",
			fields: fields{
				length: 1,
				Head:   &ListNode{Val: 1},
			},
			wantErr: false,
		},
		{
			name: "normal SinglyLinkList del",
			fields: fields{
				length: 3,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sll := &SinglyLinkedList{
				Length: tt.fields.length,
				Head:   tt.fields.Head,
			}
			if err := sll.DelAtEnd(); (err != nil) != tt.wantErr {
				t.Errorf("SinglyLinkedList.DelAtEnd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSinglyLinkedList_PopAtBegin(t *testing.T) {
	type fields struct {
		length int
		Head   *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "empty SinglyLinkedList pop",
			fields: fields{},
			want:   -1,
		},
		{
			name: "single node SinglyLinkList pop",
			fields: fields{
				length: 1,
				Head:   &ListNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "normal SinglyLinkList pop at begin",
			fields: fields{
				length: 3,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sll := &SinglyLinkedList{
				Length: tt.fields.length,
				Head:   tt.fields.Head,
			}
			if got := sll.PopAtBegin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SinglyLinkedList.PopAtBegin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_PopAtEnd(t *testing.T) {
	type fields struct {
		length int
		Head   *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "empty SinglyLinkedList pop",
			fields: fields{},
			want:   -1,
		},
		{
			name: "single node SinglyLinkList pop",
			fields: fields{
				length: 1,
				Head:   &ListNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "normal SinglyLinkList pop at end",
			fields: fields{
				length: 3,
				Head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sll := &SinglyLinkedList{
				Length: tt.fields.length,
				Head:   tt.fields.Head,
			}
			if got := sll.PopAtEnd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SinglyLinkedList.PopAtEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
