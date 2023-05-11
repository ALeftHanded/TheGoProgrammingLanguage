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

func NumberOfDistinctIslands2Op(grid [][]byte) int {
	hash := func(shape [][2]int) string {
		encode := func(shape [][2]int) string {
			x, y := shape[0][0], shape[0][1]
			coords := make([]string, len(shape))

			for i, point := range shape {
				coords[i] = strconv.Itoa(point[0]-x) + ":" + strconv.Itoa(point[1]-y)
			}

			return strings.Join(coords, "")
		}

		var shapes [][][2]int
		for _, t := range [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
			var transformed [][2]int
			for _, point := range shape {
				transformed = append(transformed, [2]int{t[0] * point[0], t[1] * point[1]})
			}
			shapes = append(shapes, transformed)
		}

		for _, s := range shapes {
			var reflected [][2]int
			for _, point := range s {
				reflected = append(reflected, [2]int{point[1], point[0]})
			}
			shapes = append(shapes, reflected)
		}

		minStr := ""
		for _, s := range shapes {
			sort.Slice(s, func(i, j int) bool {
				return s[i][0] < s[j][0] || (s[i][0] == s[j][0] && s[i][1] < s[j][1])
			})
			encoded := encode(s)
			if minStr == "" || encoded < minStr {
				minStr = encoded
			}
		}

		return minStr
	}

	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int) [][2]int
	dfs = func(i, j int) [][2]int {
		if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == '0' {
			return nil
		}

		grid[i][j] = '0'
		shape := [][2]int{{i, j}}

		for _, d := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			shape = append(shape, dfs(i+d[0], j+d[1])...)
		}

		return shape
	}

	islands := make(map[string]bool)

	for i, row := range grid {
		for j, num := range row {
			if num == '1' {
				islands[hash(dfs(i, j))] = true
			}
		}
	}

	return len(islands)
}
