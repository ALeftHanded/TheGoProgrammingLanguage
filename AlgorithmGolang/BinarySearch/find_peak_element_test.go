package BinarySearch

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"AlgorithmGolang/Utils/random"
)

func TestFindPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}

	randCase := random.GenRandomIncrIntList(10000000, 10000000, -112312312312, 10000000)
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "normal case",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 12, 9},
			},
			want: 5,
		},
		{
			name: "random increasing case",
			args: args{
				nums: randCase,
			},
			want: len(randCase) - 1,
		},
		{
			name: "single case",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			// ! find error reason
			name: "error case",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {

		start := time.Now()
		res := FindPeakElement(tt.args.nums)
		dur := time.Since(start)
		fmt.Println("FindPeakElement|Actually function duration: ", dur)

		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, res, "FindPeakElement(%v)", tt.args.nums)
		})
	}
}
