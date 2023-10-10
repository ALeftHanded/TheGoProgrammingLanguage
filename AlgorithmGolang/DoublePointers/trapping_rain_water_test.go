package DoublePointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Normal case",
			args: args{nums: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},		
		{
			name: "Empty array",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "Single element",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "Two elements",
			args: args{nums: []int{1, 2}},
			want: 0,
		},
		{
			name: "Test case with 0s and 4s",
			args: args{nums: []int{3, 0, 2, 0, 4}},
			want: 7,
		},
		{
			name: "All zeros",
			args: args{nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trap(tt.args.nums)
			assert.Equal(t, tt.want, got, "they should be equal")
		})
	}

}