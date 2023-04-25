package graph

import (
	"reflect"
	"testing"
)

func TestSurroundedRegions(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "all surrounded",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
		},
		// ! bug
		// record bug fix
		// ! error code
		//	if board[i][j] == '-' {
		//	board[i][j] = 'O'
		//	}
		//	if board[i][j] == 'O' {
		//	board[i][j] = 'X'
		//	}
		// * correct code
		// 			if board[i][j] == '-' {
		//				board[i][j] = 'O'
		//			} else if board[i][j] == 'O' {
		//				board[i][j] = 'X'
		//			}

		{
			name: "all no surrounded",
			args: args{
				board: [][]byte{
					{'O', 'O', 'O', 'O'},
					{'O', 'X', 'X', 'O'},
					{'O', 'X', 'X', 'O'},
					{'O', 'O', 'O', 'O'},
				},
			},
			want: [][]byte{
				{'O', 'O', 'O', 'O'},
				{'O', 'X', 'X', 'O'},
				{'O', 'X', 'X', 'O'},
				{'O', 'O', 'O', 'O'},
			},
		},
		{
			name: "normal",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'X', 'X'},
					{'X', 'O', 'X', 'O'},
					{'X', 'X', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'O'},
				{'X', 'X', 'X', 'X'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SurroundedRegions(tt.args.board)
		})
		if !reflect.DeepEqual(tt.args.board, tt.want) {
			t.Errorf("SurroundedRegions() = %v, want %v", tt.args.board, tt.want)
		}
	}
}
