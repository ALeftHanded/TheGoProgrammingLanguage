package maxHeapUtil

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		mh  *MaxHeap
		num int
	}
	tests := []struct {
		name string
		args args
		want *MaxHeap
	}{
		{
			name: "normal test",
			args: args{
				mh:  &MaxHeap{3, &MaxHeap{2, nil, nil}, &MaxHeap{1, nil, nil}},
				num: 4,
			},
			want: &MaxHeap{
				4, &MaxHeap{3, &MaxHeap{2, nil, nil}, nil}, &MaxHeap{1, nil, nil},
			},
		},
		{
			name: "inserting into an empty heap",
			args: args{
				mh:  nil,
				num: 5,
			},
			want: &MaxHeap{Val: 5},
		},
		{
			name: "inserting smaller value",
			args: args{
				mh:  &MaxHeap{6, &MaxHeap{4, nil, nil}, &MaxHeap{5, nil, nil}},
				num: 3,
			},
			want: &MaxHeap{6, &MaxHeap{4, &MaxHeap{3, nil, nil}, nil}, &MaxHeap{5, nil, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.mh, tt.args.num); !got.Equal(tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMaxHeapFromIntArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *MaxHeap
	}{
		{
			name: "normal test",
			args: args{arr: []int{5, 3, 2, 1, 4}},
			want: &MaxHeap{
				5, &MaxHeap{4, &MaxHeap{1, nil, nil}, &MaxHeap{3, nil, nil}}, &MaxHeap{2, nil, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitMaxHeapFromIntArrayBad(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMaxHeapFromIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapifyInit(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name       string
		args       args
		want       *MaxHeap
		validCheck bool
	}{
		{
			name: "normal test",
			args: args{arr: []int{3, 2, 5, 1, 4}},
			want: &MaxHeap{
				Val: 5,
				Left: &MaxHeap{
					3,
					&MaxHeap{1, nil, nil},
					&MaxHeap{2, nil, nil},
				},
				Right: &MaxHeap{4, nil, nil},
			},
		},
		{
			name: "even length test",
			args: args{arr: []int{1, 3, 5, 4, 2, 6}},
			want: &MaxHeap{
				Val: 6,
				Left: &MaxHeap{
					3,
					&MaxHeap{1, nil, nil},
					&MaxHeap{2, nil, nil},
				},
				Right: &MaxHeap{
					5,
					&MaxHeap{4, nil, nil},
					nil,
				}, // The value here should have been swapped with root, but the root is skipped
			},
			validCheck: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HeapifyInit(tt.args.arr)
			if !got.ValidCheck() {
				t.Errorf("[-] Heap is not valid!")
			}
			if !got.Equal(tt.want) {
				t.Errorf("HeapifyInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
