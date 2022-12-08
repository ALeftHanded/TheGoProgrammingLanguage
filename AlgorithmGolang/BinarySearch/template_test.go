package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal test",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				n:   5,
			},
			want: 2,
		},
		{
			name: "cannot find/tle test",
			args: args{
				arr: []int{-1, 0, 3, 5, 9, 12},
				n:   2,
			},
			want: -1,
		},
		{
			name: "single element array test",
			args: args{
				arr: []int{1},
				n:   1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.args.arr, tt.args.n)
			assert.Equalf(t, tt.want, got, "BinarySearch() = %v, want %v", got, tt.want)
		})
	}
}
