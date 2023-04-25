package graph

// 链接：https://leetcode.cn/problems/number-of-islands

func NumIslands(grid [][]byte) int {
	var count int
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j, m, n)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
		return
	}
	dfs(grid, i-1, j, m, n)
	dfs(grid, i+1, j, m, n)
	dfs(grid, i, j-1, m, n)
	dfs(grid, i, j+1, m, n)
}
