package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "normal case",
			args: args{
				target: 77,
				nums:   []int{2, 3, 77, 1, 2, 4, 3, 7},
			},
			want: 1,
		},
		{
			name: "empty case",
			args: args{
				target: 77,
				nums:   []int{1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			// ! 没有初始化sum
			name: "error case 1",
			args: args{
				target: 15,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "error case 2",
			args: args{
				target: 6,
				nums:   []int{10, 2, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MinSubArrayLen(tt.args.target, tt.args.nums), "MinSubArrayLen(%v, %v)", tt.args.target, tt.args.nums)
		})
	}
}
