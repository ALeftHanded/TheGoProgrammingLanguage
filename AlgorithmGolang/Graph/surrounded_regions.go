package graph

func SurroundedRegions(board [][]byte) {
	m, n := len(board), len(board[0])
	for i := range board {
		for j := range board[i] {
			if inGraphEdge(i, j, m, n) {
				if board[i][j] == 'O' {
					dfsSurroundedRegions(board, i, j, m, n)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '-' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsSurroundedRegions(board [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n ||
		board[i][j] == 'X' || board[i][j] == '-' {
		return
	}
	board[i][j] = '-'
	dfsSurroundedRegions(board, i-1, j, m, n)
	dfsSurroundedRegions(board, i+1, j, m, n)
	dfsSurroundedRegions(board, i, j-1, m, n)
	dfsSurroundedRegions(board, i, j+1, m, n)
}

func inGraphEdge(i, j, m, n int) bool {
	return i == 0 || i == m-1 || j == 0 || j == n-1
}
