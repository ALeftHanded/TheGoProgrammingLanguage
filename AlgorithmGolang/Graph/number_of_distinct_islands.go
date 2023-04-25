package graph

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func NumberOfDistinctIslands(grid [][]byte) int {
	res := make(map[string]bool)
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				tmpShape := make([][]int, 0, m*n)
				tmpShape = dfsForNumberOfDistinctIslands(grid, i, j, m, n, tmpShape)
				if len(tmpShape) > 0 {
					res[hashShapeOp(tmpShape)] = true
				}
			}
		}
	}
	return len(res)
}

func dfsForNumberOfDistinctIslands(grid [][]byte, i, j, m, n int, shape [][]int) [][]int {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
		return shape
	}
	if grid[i][j] == '1' {
		shape = append(shape, []int{i, j})
		grid[i][j] = '0'
	}

	shape = dfsForNumberOfDistinctIslands(grid, i-1, j, m, n, shape)
	shape = dfsForNumberOfDistinctIslands(grid, i+1, j, m, n, shape)
	shape = dfsForNumberOfDistinctIslands(grid, i, j-1, m, n, shape)
	shape = dfsForNumberOfDistinctIslands(grid, i, j+1, m, n, shape)
	return shape
}

func NumberOfDistinctIslandsOp(grid [][]byte) int {
	res := make(map[string]bool)
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				tmpShape := make([][]int, 0, m*n)
				dfsForNumberOfDistinctIslandsOp(grid, i, j, m, n, &tmpShape)
				if len(tmpShape) > 0 {
					res[hashShapeOp(tmpShape)] = true
				}
			}
		}
	}
	return len(res)
}

func dfsForNumberOfDistinctIslandsOp(grid [][]byte, i, j, m, n int, shape *[][]int) {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
		return
	}

	*shape = append(*shape, []int{i, j})
	grid[i][j] = '0'

	for _, direction := range directions {
		dfsForNumberOfDistinctIslandsOp(grid, i+direction[0], j+direction[1], m, n, shape)
	}
}

func hashShapeOp(shape [][]int) string {
	baseIndex := []int{shape[0][0], shape[0][1]}
	for i := range shape {
		shape[i][0] -= baseIndex[0]
		shape[i][1] -= baseIndex[1]
	}
	sort.Slice(shape, func(i, j int) bool {
		if shape[i][0] == shape[j][0] {
			return shape[i][1] < shape[j][1]
		}
		return shape[i][0] < shape[j][0]
	})

	var builder strings.Builder
	totalLength := 2 * len(shape)
	builder.Grow(totalLength)
	for i := range shape {
		for j := range shape[i] {
			builder.WriteString(strconv.Itoa(shape[i][j]))
		}
	}
	return builder.String()
}

func hashShape(shape [][]int) string {
	baseIndex := []int{shape[0][0], shape[0][1]}
	for i := range shape {
		shape[i][0] -= baseIndex[0]
		shape[i][1] -= baseIndex[1]
	}
	sort.Slice(shape, func(i, j int) bool {
		if shape[i][0] == shape[j][0] {
			return shape[i][1] < shape[j][1]
		}
		return shape[i][0] < shape[j][0]
	})
	return fmt.Sprint(shape)
}

// NumberOfDistinctIslandsWithCustomHash for benchmark
func NumberOfDistinctIslandsWithCustomHash(grid [][]byte, hashFunc func([][]int) string) int {
	res := make(map[string]bool)
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				tmpShape := make([][]int, 0, m*n)
				tmpShape = dfsForNumberOfDistinctIslands(grid, i, j, m, n, tmpShape)
				if len(tmpShape) > 0 {
					res[hashFunc(tmpShape)] = true
				}
			}
		}
	}
	return len(res)
}
