package BinarySearch

import "testing"

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
			name: "tle test",
			args: args{
				arr: []int{-1, 0, 3, 5, 9, 12},
				n:   2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.arr, tt.args.n); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
