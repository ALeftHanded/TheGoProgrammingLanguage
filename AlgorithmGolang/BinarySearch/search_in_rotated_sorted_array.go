package BinarySearch

// * There is an integer arrayUtil nums sorted in ascending order (with distinct values).
// *
// * Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
// * such that the resulting arrayUtil is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// * For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
// *
// * Given the arrayUtil nums after the possible rotation and an integer target,
// * return the index of target if it is in nums, or -1 if it is not in nums.
//
// * You must write an algorithm with O(log n) runtime complexity.

// * Constraints:
// * 1 <= nums.length <= 5000
// * -104 <= nums[i] <= 104
// * All values of nums are unique.
// * nums is an ascending arrayUtil that is possibly rotated.
// * -104 <= target <= 104

// * 分析，两种情形，单调递增或者先增再后增
// * 1. mid小于target
// * 1.1. 单调递增情形，取右半区间
// * 1.2. mid命中先增区间，取右半区间
// * 1.3. mid命中后增区间，target在后增区间，取右半区间
// * 1.4. mid命中后增区间，target在先增区间，取左半区间
// * 2. mid大于target
// * 2.1. 单调递增情形，取左半区间
// * 2.2. mid命中先增区间，target在后增区间，取右半区间
// * 2.3. mid命中先增区间，target在先增区间，取左半区间
// * 2.4. mid命中后增区间，取左半区间

func SearchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			if nums[mid] < nums[left] && nums[right] < target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > nums[right] && nums[left] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
