package BinarySearch

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
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
		//{
		//	name: "random list test",
		//	args: args{
		//		nums: genRandomNums(1, 2500, -104, 104),
		//	},
		//	// ! deliver the standard answer.
		//	want: 0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLIS(tt.args.nums), "lengthOfLIS(%v)", tt.args.nums)
		})
	}
}

func genRandomNums(minLen, maxLen, minElement, maxElement int) []int {
	var res []int
	ranLen := rand.Intn(maxLen-minLen) + minLen
	for i := 0; i < ranLen; i++ {
		randNum := rand.Intn(maxElement-minElement) + minElement
		res = append(res, randNum)
	}
	return res
}
