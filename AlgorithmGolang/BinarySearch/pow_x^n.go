package BinarySearch

// * 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n ）。

// * -100.0 < x < 100.0
// * -2^31 <= n <= 2^31-1
// * n 是一个整数
// * -10^4 <= x^n <= 10^4

func PowXN(x float64, n int) float64 {
	// ? exclude 0^0
	if n < 0 {
		return 1 / PowXN(x, -n)
	}
	var res, tmp float64
	res, tmp = 1, x
	for n > 0 {
		if n%2 == 1 {
			res *= tmp
		}
		tmp *= tmp
		n >>= 1
	}
	return res
}
