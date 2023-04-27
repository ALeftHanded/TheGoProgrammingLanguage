package graph

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func NumberOfDistinctIslands2(grid [][]byte) int {
	hashSet := make(map[string]struct{})
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				tmpShape := make([][]int, 0, m*n)
				dfsForNumberOfDistinctIslandsOp(grid, i, j, m, n, &tmpShape)
				if len(tmpShape) > 0 {
					hashSet[genHashStr(tmpShape)] = struct{}{}
				}
			}
		}
	}
	return len(hashSet)
}

// ? if left reflection and top reflection are same, only add once
// * I think it is...
// M[i,j] => left reflection => M[i, n-j] =>
// rotate 90 => M[j, i] =>
// rotate 180 => M[m-i, j] =>
// rotate 270 => M[n-j, m-i]

// M[i,j] => top reflection => M[m-i, j] =>
// rotate 90 => M[n-j, m-i] =>
// rotate 180 => M[i, n-j] =>
// rotate 270 => M[j, i] =>
func genHashStr(shape [][]int) string {
	uniqueHashStrList := make(map[string]bool)
	baseShape := genBaseShape(shape)
	uniqueHashStrList[buildSingleHashStr(baseShape)] = true

	if !reflect.DeepEqual(baseShape, leftRotate90(baseShape)) {
		for i := 0; i < 3; i++ {
			baseShape = leftRotate90(baseShape)
			uniqueHashStrList[buildSingleHashStr(baseShape)] = true
		}

		leftRefectionShape := leftReflection(baseShape)
		uniqueHashStrList[buildSingleHashStr(leftRefectionShape)] = true
		for i := 0; i < 3; i++ {
			leftRefectionShape = leftRotate90(leftRefectionShape)
			uniqueHashStrList[buildSingleHashStr(leftRefectionShape)] = true
		}
	}

	hashStrList := make([]string, 0, len(uniqueHashStrList))
	for hashStr := range uniqueHashStrList {
		hashStrList = append(hashStrList, hashStr)
	}
	sort.Strings(hashStrList)
	return hashStrList[0]
}

func genBaseShape(shape [][]int) [][]int {
	baseIndex := []int{shape[0][0], shape[0][1]}
	offsetX, offsetY := 0, 0
	for i := range shape {
		shape[i][0] -= baseIndex[0]
		if shape[i][0] < offsetX {
			offsetX = shape[i][0]
		}
		shape[i][1] -= baseIndex[1]
		if shape[i][1] < offsetY {
			offsetY = shape[i][1]
		}
	}

	if offsetY < 0 || offsetX < 0 {
		for i := range shape {
			shape[i][0] -= offsetX
			shape[i][1] -= offsetY
		}
	}
	sortShape(shape)
	return shape
}

func leftRotate90(shape [][]int) [][]int {
	// rotate shape left 90 angle
	n := 0
	for _, index := range shape {
		if index[1] > n {
			n = index[1]
		}
	}
	// 0, n => 0, 0
	// 0, 0 => n, 0
	// m, 0 => n, m
	// m, n => 0, m
	// x, y => n-y, x
	res := make([][]int, len(shape))
	for i := range res {
		res[i] = []int{n - shape[i][1], shape[i][0]}
	}
	sortShape(res)
	return res
}

// leftReflection shape left reflection
func leftReflection(shape [][]int) [][]int {
	offsetY := 0
	for i := range shape {
		shape[i][1] = 0 - shape[i][1]
		if shape[i][1] < offsetY {
			offsetY = shape[i][1]
		}
	}
	for i := range shape {
		shape[i][1] -= offsetY
	}
	sortShape(shape)
	return shape
}

func buildSingleHashStr(shape [][]int) string {
	var builder strings.Builder
	totalLength := 2 * len(shape)
	builder.Grow(totalLength)
	for i := range shape {
		for j := range shape[i] {
			// index to string
			builder.WriteString(strconv.Itoa(shape[i][j]))
		}
	}
	return builder.String()
}

func sortShape(shape [][]int) {
	sort.Slice(shape, func(i, j int) bool {
		if shape[i][0] == shape[j][0] {
			return shape[i][1] < shape[j][1]
		}
		return shape[i][0] < shape[j][0]
	})
}
