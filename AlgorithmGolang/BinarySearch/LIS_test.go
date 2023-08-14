package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AlgorithmGolang/libs/arrayUtil"
	"AlgorithmGolang/libs/random"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	randNums := random.GenRandomNums(1, 2500, -104, 104)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal ascending list test",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "normal descending list test",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: 1,
		},
		{
			name: "error case from leetcode",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
		{
			name: "random list test",
			args: args{
				nums: randNums,
			},
			// ! deliver the standard answer.
			want: arrayUtil.MaxInt(initDPForLIS(randNums)...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLIS(tt.args.nums), "lengthOfLIS(%v)", tt.args.nums)
		})
	}
}
