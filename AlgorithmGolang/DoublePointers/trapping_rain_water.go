package DoublePointers

func Trap(height []int) int {
	l,r := 0, len(height)-1
	l_max, r_max := 0,0
	res := 0
	for l<r {
		if height[l] < height[r] {
			if height[l] < l_max {
				// height[l] < height[r] will ensure we can store l_max - height[l] water
				res += l_max - height[l]
			} else {
				l_max = height[l]
			}
			l++

		} else {
			if height[r] < r_max {
				res += r_max - height[r]
			} else {
				r_max = height[r]
			}
			r--
		}
	}
	return res
}