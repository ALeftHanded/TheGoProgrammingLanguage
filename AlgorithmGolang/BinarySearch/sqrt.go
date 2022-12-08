package BinarySearch

import "math"

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留整数部分 ，小数部分将被舍去 。
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		// * mid > sqrt(x)
		if mid*mid > x {
			right = mid - 1
			// * mid <= sqrt(x) < mid + 1
		} else if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else {
			left = mid + 1
		}
	}
	return 0
}

func MathSqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func mySqrtOptimization(x int) int {
	left, right := 0, x
	var res int
	for left <= right {
		mid := (left + right) >> 1
		// * mid > sqrt(x)
		if mid*mid > x {
			right = mid - 1
		} else {
			// ? why use res = mid rather than res = mid + 1
			// * The res should fulfill the condition that res^2 <= x < (res+1)^2
			// * If we set res = mid + 1, it is possible that (mid+1)^2 <= x
			// ! wrong case x = 0 will result in 1
			res = mid
			left = mid + 1

		}
	}
	return res
}
