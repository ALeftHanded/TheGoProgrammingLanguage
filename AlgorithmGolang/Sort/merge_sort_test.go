package Sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal test",
			args: args{arr: []int{5, 3, 2, 1, 4}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortedList(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal test",
			args: args{left: []int{1, 3, 5, 7, 9}, right: []int{2, 4, 6, 8}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "all number in left smaller than right",
			args: args{left: []int{1, 2, 3, 4, 5}, right: []int{6, 7, 8, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "all number in left bigger than right",
			args: args{right: []int{1, 2, 3, 4, 5}, left: []int{6, 7, 8, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "left nil",
			args: args{right: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "right nil",
			args: args{left: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "all nil",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedLists(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
