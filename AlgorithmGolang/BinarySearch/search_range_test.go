package BinarySearch

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
			want: []int{0, 4},
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

			start := time.Now()
			res := SearchRange(tt.args.nums, tt.args.target)
			dur := time.Since(start)
			fmt.Println("SearchRange|Actually function duration: ", dur)

			assert.Equalf(t, tt.want, res, "SearchRange(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
