package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AlgorithmGolang/libs/measureUtil"
)

func TestSearchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "normal case",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 2, 3, 4, 4, 4, 4, 4, 4, 6, 7, 7, 7, 7},
				target: 6,
			},
			want: []int{13, 13},
		},
		{
			name: "error case",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, _ := measureUtil.ExecutionTime(SearchRange, tt.args.nums, tt.args.target)
			assert.Equalf(t, tt.want, res[0], "SearchRange(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
