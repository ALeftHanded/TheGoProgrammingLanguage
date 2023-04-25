package graph

import "testing"

type args struct {
	grid [][]byte
}

func runTestCases(testCases []struct {
	name string
	args args
	want int
}, f func([][]byte) int) {
	for _, tt := range testCases {
		_ = f(tt.args.grid)
	}
}

func TestNumberOfDistinctIslands(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		// ! Issue 1: Appending elements in inner function might not change the original slice.
		// ! Issue 2: Assigning a slice to another variable may lead to changes in the original slice.

		{
			name: "normal",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '1', '1'},
					{'1', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '1'},
					{'1', '1', '0', '1', '1'},
					{'1', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '1'},
					{'1', '1', '0', '1', '1'},
					{'1', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '1'},
				},
			},
			want: 4,
		},
		{
			name: "all 1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
				},
			},
			want: 1,
		},
		{
			name: "all 0",
			args: args{
				grid: [][]byte{
					{'0', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfDistinctIslands(tt.args.grid); got != tt.want {
				t.Errorf("NumberOfDistinctIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

//var benchmarkShapes = [][][]int{
//	{
//		{0, 0}, {1, 1}, {2, 2}, {3, 3},
//	},
//	{
//		{0, 0}, {1, 0}, {2, 0}, {3, 0},
//	},
//	{
//		{0, 0}, {0, 1}, {0, 2}, {0, 3},
//	},
//	{
//		{0, 0}, {1, 1}, {2, 2}, {3, 3},
//		{4, 4}, {5, 5}, {6, 6}, {7, 7},
//	},
//}
//
//func BenchmarkHashShape(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for _, shape := range benchmarkShapes {
//			hashShape(shape)
//		}
//	}
//}
//
//func BenchmarkHashShape2(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for _, shape := range benchmarkShapes {
//			hashShapeOp(shape)
//		}
//	}
//}

func BenchmarkNumberOfDistinctIslandsWithCustomHash(b *testing.B) {

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '1', '1'},
					{'1', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '1'},
					{'1', '1', '0', '1', '1'},
					{'1', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '1'},
					{'1', '1', '0', '1', '1'},
					{'1', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '1'},
				},
			},
			want: 4,
		},
		{
			name: "all 1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
				},
			},
			want: 1,
		},
		{
			name: "all 0",
			args: args{
				grid: [][]byte{
					{'0', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 0,
		},
		{
			name: "large grid",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0', '1', '1', '0', '1', '1'},
					{'1', '1', '0', '0', '0', '0', '1', '0', '1', '0'},
					{'0', '0', '0', '1', '1', '0', '0', '0', '0', '1'},
					{'1', '1', '0', '1', '1', '0', '1', '1', '0', '0'},
					{'1', '0', '0', '0', '0', '1', '1', '0', '0', '0'},
					{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
					{'1', '1', '0', '1', '1', '0', '1', '1', '0', '0'},
					{'1', '0', '0', '0', '0', '1', '1', '0', '0', '0'},
					{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
					{'1', '1', '0', '1', '1', '0', '1', '1', '0', '0'},
				},
			},
			want: 6,
		},
	}
	b.ResetTimer()
	//b.Run("hashShape", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		for _, tt := range tests {
	//			_ = NumberOfDistinctIslandsWithCustomHash(tt.args.grid, hashShape)
	//		}
	//	}
	//})
	//b.Run("hashShapeOp", func(b *testing.B) {
	//
	//	for i := 0; i < b.N; i++ {
	//		for _, tt := range tests {
	//			_ = NumberOfDistinctIslandsWithCustomHash(tt.args.grid, hashShapeOp)
	//		}
	//	}
	//})
	b.Run("WithHashShapeOp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runTestCases(tests, NumberOfDistinctIslands)

		}
	})
	b.Run("dfsOp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runTestCases(tests, NumberOfDistinctIslandsOp)

		}
	})
}
