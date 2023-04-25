package graph

func MaxAreaOfIsland(grid [][]int) int {
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				tmp := dfsMaxArea(grid, i, j, len(grid), len(grid[0]), 0)
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

func dfsMaxArea(grid [][]int, i, j, m, n, count int) int {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return 0
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
		count++
	}

	count += dfsMaxArea(grid, i-1, j, m, n, 0)
	count += dfsMaxArea(grid, i, j-1, m, n, 0)
	count += dfsMaxArea(grid, i, j+1, m, n, 0)
	count += dfsMaxArea(grid, i+1, j, m, n, 0)
	return count
}
