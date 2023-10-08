package BinarySearch

// * 给定一个含有n个正整数的数组和一个正整数 target 。

// * 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr]
// * 返回其长度。如果不存在符合条件的子数组，返回 0 。

// * 1 <= target <= 10^9
// * 1 <= nums.length <= 10^5
// * 1 <= nums[i] <= 10^5

// * 二分查找？

func MinSubArrayLen(target int, nums []int) int {
	return -1
}

// * 滑动窗口
// * left, right为左右两端
// * 初次达到sum(nums[left:right+1]) >= target时进入下一阶段，否则0退出
// * left开始右移，若sum(nums[left:right+1]) < target，right右移，这个阶段去维护长度，取最小值

func MinSubArrayLen1stAC(target int, nums []int) int {
	left, right, sum, res := 0, 0, 0, 0
	// phase 1
	for ; right < len(nums); right++ {
		sum += nums[right]
		if sum >= target {
			res = right + 1
			break
		}
	}
	// phase 2
	if res > 0 {
		for {
			for sum >= target {
				sum -= nums[left]
				left++
				if sum >= target && right-left+1 < res {
					res = right - left + 1
				}
			}
			for sum < target {
				right++
				if right == len(nums) {
					return res
				}
				sum += nums[right]
			}
			if right-left+1 < res {
				res = right - left + 1
			}
		}
	}
	return res
}
