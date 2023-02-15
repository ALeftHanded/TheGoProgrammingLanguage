package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrixII(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "normal case",
			args: args{
				matrix: [][]int{
					{1, 7, 8},
					{2, 9, 10},
					{3, 11, 12},
				},
				target: 2,
			},
			want: true,
		},
		{
			name: "edge case 1",
			args: args{
				matrix: [][]int{
					{1, 7, 8},
					{2, 9, 10},
					{3, 11, 12},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "edge case 2",
			args: args{
				matrix: [][]int{
					{1, 7, 8},
					{2, 9, 10},
					{3, 11, 12},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "edge case 3",
			args: args{
				matrix: [][]int{
					{1, 7, 8},
					{2, 9, 10},
					{3, 11, 12},
				},
				target: 8,
			},
			want: true,
		},
		{
			name: "edge case 4",
			args: args{
				matrix: [][]int{
					{1, 7, 8},
					{2, 9, 10},
					{3, 11, 12},
				},
				target: 12,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SearchMatrixII(tt.args.matrix, tt.args.target), "SearchMatrixII(%v, %v)", tt.args.matrix, tt.args.target)
		})
	}
}
