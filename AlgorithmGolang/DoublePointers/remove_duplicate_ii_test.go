package DoublePointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		wantNums []int
	}{
		{
			name: "normal test",
			args: args{
				nums: []int{0,0,1,1,2,2,2,2,3,3,3,4,4,4,4},
			},
			wantNums: []int{0,0,1,1,2,2,3,3,4,4},
		},
		{
			name: "single num test",
			args: args{
				nums: []int{0,1,2,3,4,5,6,7},
			},
			wantNums: []int{0,1,2,3,4,5,6,7},
		},
		{
			name: "one num test",
			args: args{
				nums: []int{0,0,0},
			},
			wantNums: []int{0,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicatesII(tt.args.nums)
			want := len(tt.wantNums)
			assert.Equal(t, want, got, "RemoveDuplicatesII() = %v, want %v", got, want)
			assert.Equal(t, want, got, "RemoveDuplicatesII() nums = %v, want nums %v", tt.args.nums[:want], tt.wantNums)
		})
	}
}
