package BinarySearch

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindPeakElement(t *testing.T) {
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
			name: "normal case",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {

		start := time.Now()
		res := BFForFPE(tt.args.nums)
		dur := time.Since(start)
		fmt.Println("BFForFPE|Actually function duration: ", dur)

		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, res, FindPeakElement(tt.args.nums), "FindPeakElement(%v)", tt.args.nums)
		})
	}
}
