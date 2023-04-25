package graph

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all 1",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1},
				},
			},
			want: 15,
		},
		{
			name: "all 0",
			args: args{
				grid: [][]int{
					{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "normal",
			args: args{
				grid: [][]int{
					{1, 1, 1, 0, 1},
					{1, 1, 1, 0, 1},
					{0, 0, 0, 0, 1},
					{1, 1, 0, 1, 1},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAreaOfIsland(tt.args.grid); got != tt.want {
				t.Errorf("MaxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
