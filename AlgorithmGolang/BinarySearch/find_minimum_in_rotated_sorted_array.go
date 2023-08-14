package BinarySearch

// Suppose an arrayUtil of length n sorted in ascending order is rotated between 1 and n times.
// For example, the arrayUtil nums = [0,1,2,4,5,6,7] might become:
//
// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
//
// Notice that rotating an arrayUtil [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the arrayUtil [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
//
// Given the sorted rotated arrayUtil nums of unique elements,
// return the minimum element of this arrayUtil.
//
// You must write an algorithm that runs in O(log n) time.

// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// All the integers of nums are unique.
// nums is sorted and rotated between 1 and n times.

// * 分析：单增或先增后增
// * 单增或单元素数组直接返回nums[left]
// * 1. nums[mid] < nums[left] 必然先增后增，且命中后增区域，此时在左闭区间, 需要包括mid
// * 2. nums[mid] > nums[left]
// * 2.1. 单增，此时在nums[left]，直接返回nums[left]
// * 2.2. 先增后增且命中先增区域，此时在右闭区间，闭区间不需要包括mid
// * 3. nums[mid] == nums[left] 只可能是单元素数组，返回单元素即可
// ! 错误点1，第3条分析有错误，忽略了双元素数组的情况，还可能是[2, 1]
// ! 错误点2，没有严格按照二分模版写，导致分支写错
// ! 错误点3，忽略了循环中出现查询区间单增的情况，即用初始情形覆盖循环中情形

func FindMinimumInRotatedSortedArray(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == nums[left] {
			if nums[right] < nums[left] {
				return nums[right]
			}
			return nums[left]
		}
		if nums[mid] > nums[left] {
			if nums[left] < nums[right] {
				return nums[left]
			}
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// * 分析：单增或先增后增
// * 单增或单元素数组直接返回nums[left]
// * 1. nums[mid] == nums[right] 必然为单元素数组
// * (left + right) >> 1 <= right
// * if left + right = 2k+1 (k是任意整数), left < right, 有right >= k+1 > k 即 (left + right) >>1
// * if left + right = 2k, left < right, 有 right > k 即 (left + right) >>1
// * if left == right, right == (left+right) >> 1
// * 1. nums[mid] > nums[right] 必然先增后增，且命中先增区域，此时最小值在右闭区间, 不需要包括mid
// * 2. nums[mid] < nums[right]
// * 2.1. 单增，此时最小值在nums[left]，直接返回nums[left]，也在左闭区间
// * 2.2. 先增后增且命中后增区域，此时最小值在左闭区间, 需要带上mid，mid有可能是min值

func Optimization_FindMinimumInRotatedSortedArray_v1(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == nums[right] {
			break
		}
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right]
}

// * nums[mid] == nums[right]时break和left<=right中的等于可以抵消掉

func Optimization_FindMinimumInRotatedSortedArray_final(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right]
}
