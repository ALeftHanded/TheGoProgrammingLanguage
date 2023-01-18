package BinarySearch

// ? 给两个整数数组 nums1 和 nums2 ，返回两个数组中公共的、长度最长的子数组的长度。

// * 1 <= nums1.length, nums2.length <= 1000
// * 0 <= nums1[i], nums2[i] <= 100

// * dp[i] 以nums1[i]为开头的和nums2的MaximumLengthOfRepeatedSubarray

// l1 = len(nums1) - 1
// dp[l1]
// dp[l1 - 1] = dp[l1] + 1 if [nums1[l1 - 1], nums1[l1]] can be find in nums2 else 1

func MaximumLengthOfRepeatedSubarray(nums1 []int, nums2 []int) int {

}
