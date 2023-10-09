package MonotonicStack

func Trap(height []int) int {
	var water int
	// 保持这个栈的单调非严格递减，即里面放了一堆高度持续降低的柱子
	decreasingStack := []int{}

	for i, h := range height {
		if len(decreasingStack) > 0 {
			// 取出栈顶柱子的高度
			stackTopHeight := height[decreasingStack[len(decreasingStack)-1]]
			// 当遇到新高度的柱子，说明有坑可以接雨水
			for h > stackTopHeight {
				// ? 坑底出栈
				decreasingStack = decreasingStack[:len(decreasingStack)-1]

				// 左边没柱子了，没坑，退出
				if len(decreasingStack) == 0 {
					break
				}
				// 取出坑左侧柱子和高度
				newStackTopIndex := decreasingStack[len(decreasingStack)-1]
				newStackTopHeight := height[newStackTopIndex]

				// 计算高度差
				heightLess := (min(newStackTopHeight, h) - stackTopHeight)
				// 计算两根柱子的距离
				distance := i - newStackTopIndex - 1
				water += distance * heightLess

				// 更新栈顶柱子的高度
				stackTopHeight = newStackTopHeight
			}
		}

		decreasingStack = append(decreasingStack, i)
	}
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
