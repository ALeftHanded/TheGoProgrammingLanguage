package BinarySearch

// * 峰值元素是指其值严格大于左右相邻值的元素。
// * 给你一个整数数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// * 你可以假设nums[-1] = nums[n] = -∞ 。
// * 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

// * 1 <= nums.length <= 1000
// * -2^31 <= nums[i] <= 2^31 - 1
// * 对于所有有效的 i 都有 nums[i] != nums[i + 1]

// * 序列是单增的，则peak是数组末尾的值
// * 序列非单增，找到第一个拐点即为peak值
// * 如何寻找拐点？

// * 二分查找
// * mid点有如下情形
// * 单增情形右半区间必然有
// * 满足峰值条件，直接返回
// * 单减、波谷情形，左半区间

func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	if len(nums) == 1 {
		return 0
	}
	for left <= right {
		mid := (right-left)>>1 + left
		// ! 这里有可能区间还剩两个元素
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				return right
			}
		} else {
			if nums[mid] > nums[mid-1] {
				if mid == len(nums)-1 {
					return mid
				}
				if nums[mid] > nums[mid+1] {
					return mid
				} else {

					left = mid + 1
				}
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// * 暴力寻找

func FindPeakElement1stAC(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i-1] >= nums[i] {
			return i - 1
		}
	}
	return len(nums) - 1
}
