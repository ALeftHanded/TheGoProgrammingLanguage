package BinarySearch

// * 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n ）。

// * -100.0 < x < 100.0
// * -2^31 <= n <= 2^31-1
// * n 是一个整数
// * -10^4 <= x^n <= 10^4

func PowXN(x float64, n int) float64 {
	// ? exclude 0^0
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / PowXN(x, -n)
	}
	var res, tmp float64
	binN := BinInt(n)
	res, tmp = 1, x
	for i := 0; i < len(binN); i++ {
		if binN[i] {
			res *= tmp
		}
		tmp *= tmp
	}
	return res
}

func BinInt(n int) []bool {
	res := make([]bool, 0, 32)
	for n > 0 {
		res = append(res, n%2 == 1)
		n >>= 1
	}
	return res
}
