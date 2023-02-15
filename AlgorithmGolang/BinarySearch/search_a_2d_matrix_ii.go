package BinarySearch

// * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
// * 每行的元素从左到右升序排列。
// * 每列的元素从上到下升序排列。

// * m == matrix.length
// * n == matrix[i].length
// * 1 <= n, m <= 300
// * -10^9<= matrix[i][j] <= 10^9
// * -10^9<= target <= 10^9

// * 如果大于target，则左上肯定不在搜索范围，在右上，和下部（不包含mid所在行）使用递归搜索
// * 如果小于target，则搜索范围，在右上，和下部使用递归搜索

func SearchMatrixII(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	rowLen, colLen := len(matrix), len(matrix[0])
	left, right := 0, rowLen*colLen-1
	for left <= right {
		mid := left + (right-left)>>1
		midRow, midCol := mid/colLen, mid%colLen
		if matrix[midRow][midCol] == target {
			return true
		} else if matrix[midRow][midCol] < target {
			if SearchMatrixII(genSubMatrix(matrix, midRow, midCol, GenLowerMatrixForBiggerTarget), target) ||
				SearchMatrixII(genSubMatrix(matrix, midRow, midCol, GenUpperRightMatrixForBiggerTarget), target) {
				return true
			}
			return false
		} else {
			if SearchMatrixII(genSubMatrix(matrix, midRow, midCol, GenUpperMatrixForSmallerTarget), target) ||
				SearchMatrixII(genSubMatrix(matrix, midRow, midCol, GenLowerLeftMatrixForSmallerTarget), target) {
				return true
			}
			return false
		}
	}
	return false
}

type genSubMatrixMode int

const (
	// GenUpperRightMatrixForBiggerTarget 截取右上部分矩阵，不包含坐标所在列，但包含坐标所在行
	GenUpperRightMatrixForBiggerTarget = 1
	// GenLowerMatrixForBiggerTarget 截取下半部分矩阵，不包含坐标所在行
	GenLowerMatrixForBiggerTarget = 2
	// GenUpperMatrixForSmallerTarget 截取上半部分矩阵，不包含坐标所在行
	GenUpperMatrixForSmallerTarget = 3
	// GenLowerLeftMatrixForSmallerTarget 截取左下部分矩阵，不包含坐标所在列，但包含坐标所在行
	GenLowerLeftMatrixForSmallerTarget = 4
)

func genSubMatrix(matrix [][]int, iRow, iCol int, mode genSubMatrixMode) [][]int {
	var res [][]int
	switch mode {
	case GenUpperRightMatrixForBiggerTarget:
		if iCol < len(matrix[0])-1 {
			for i := 0; i <= iRow; i++ {
				res = append(res, matrix[i][iCol+1:])
			}
		}
	case GenLowerMatrixForBiggerTarget:
		if iRow < len(matrix)-1 {
			res = matrix[iRow+1:]
		}
	case GenUpperMatrixForSmallerTarget:
		if iRow > 0 {
			res = matrix[:iRow]
		}
	case GenLowerLeftMatrixForSmallerTarget:
		if iCol > 0 {
			for i := iRow; i < len(matrix); i++ {
				res = append(res, matrix[i][:iCol])
			}
		}
	}
	return res
}
