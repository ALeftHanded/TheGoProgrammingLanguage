package SinglyLinkedList

import (
	"reflect"
	"testing"

	"AlgorithmGolang/libs/arrayUtil"
	"AlgorithmGolang/libs/measureUtil"
	"AlgorithmGolang/libs/random"
)

func newCopy(lists []*ListNode) []*ListNode {
	tmp := make([]*ListNode, 0, len(lists))
	for i := range lists {
		tmp = append(tmp, lists[i].DeepCopyList())
	}
	return tmp
}

func TestMergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}

	// gen 5 random long *ListNode
	const (
		k         = 2000
		minLength = 30000
		maxLength = 30000
		minNum    = -100000
		maxNum    = 100000
	)

	var randomListNodes []*ListNode
	for i := 0; i < k; i++ {
		baseArr := random.GenRandomIncrIntList(minLength, maxLength, minNum, maxNum)
		var tmpListNode *ListNode
		if baseArr == nil {
			tmpListNode = nil
		} else {
			tmpListNode = ArrayToSinglyLinkedList(arrayUtil.ConvertIntToInterface(baseArr)).Head
		}
		randomListNodes = append(randomListNodes, tmpListNode)
	}
	tmp := newCopy(randomListNodes)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal test",
			args: args{lists: []*ListNode{
				ArrayToSinglyLinkedList([]interface{}{1, 2, 3}).Head,
				ArrayToSinglyLinkedList([]interface{}{4}).Head,
				ArrayToSinglyLinkedList([]interface{}{5, 6}).Head,
				ArrayToSinglyLinkedList([]interface{}{0, 7}).Head,
				ArrayToSinglyLinkedList([]interface{}{-1, 8}).Head,
			}},
			want: ArrayToSinglyLinkedList([]interface{}{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}).Head,
		},
		{
			name: "empty test",
			args: args{lists: []*ListNode{
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
			}},
			want: ArrayToSinglyLinkedList([]interface{}{}).Head,
		},
		{
			name: "nil test",
			args: args{lists: []*ListNode{
				nil,
				nil,
			}},
			want: nil,
		},
		{
			name: "empty and nil test",
			args: args{lists: []*ListNode{
				nil,
				ArrayToSinglyLinkedList([]interface{}{}).Head,
			}},
			want: ArrayToSinglyLinkedList([]interface{}{}).Head,
		},
		{
			name: "time cost test",
			args: args{lists: randomListNodes},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "time cost test" {
				measureUtil.ExecutionTime(MergeKLists, tmp)
				tmp = newCopy(tt.args.lists)
				measureUtil.ExecutionTime(MergeKListsWithMinHeap, tmp)
				tmp = newCopy(tt.args.lists)
				measureUtil.ExecutionTime(MergeKListsWithMinHeapOp, tmp)

			} else {
				if got := MergeKListsWithMinHeapOp(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("MergeKLists() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
