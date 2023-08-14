package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal test even elements",
			args: args{
				nums:   []int{4, 5, 6, 1, 2, 3},
				target: 1,
			},
			want: 3,
		},
		{
			name: "normal test odd elements",
			args: args{
				nums:   []int{4, 5, 1, 2, 3},
				target: 5,
			},
			want: 1,
		},
		{
			name: "test single element",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "test ascending arrayUtil",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 2,
			},
			want: 1,
		},
		{
			// ! 需要cover这个用例
			name: "test two elements and descending arrayUtil",
			args: args{
				nums:   []int{2, 1},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SearchInRotatedSortedArray(tt.args.nums, tt.args.target),
				"SearchInRotatedSortedArray(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
