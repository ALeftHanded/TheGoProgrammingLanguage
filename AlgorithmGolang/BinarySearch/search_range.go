package BinarySearch

// * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

// * 0 <= nums.length <= 10^5
// * -10^9 <= nums[i] <= 10^9
// * nums是一个非递减数组
// * -10^9 <= target <= 10^9

func SearchRange(nums []int, target int) []int {
	resRight, resLeft := -1, -1
	// find right
	resRight = SearchRangeRight(nums, target)
	// find left
	if resRight != -1 {
		resLeft = SearchRangeLeft(nums, target)
	}
	return []int{resLeft, resRight}
}

func SearchRangeRight(nums []int, target int) int {
	res := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else {
			if nums[mid] == target {
				res = mid
			}
			left = mid + 1
		}
	}
	return res
}

func SearchRangeLeft(nums []int, target int) int {
	res := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			if nums[mid] == target {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func SearchRange1stAC(nums []int, target int) []int {
	res := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			res = mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if res == -1 {
		return []int{-1, -1}
	}
	return searchLeftAndRight(nums, res, target)
}

func searchLeftAndRight(nums []int, index, target int) []int {
	left, right := index, index
	for left >= 0 && nums[left] == target {
		left--
	}
	for right < len(nums) && nums[right] == target {
		right++
	}
	return []int{left + 1, right - 1}
}
