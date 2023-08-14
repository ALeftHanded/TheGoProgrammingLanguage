package Array

import arrayUtil "AlgorithmGolang/libs/arrayUtil"

// * 给定两个数组nums1和nums2 ，返回 它们的交集。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

// * 1 <= nums1.length, nums2.length <= 1000
// * 0 <= nums1[i], nums2[i] <= 1000

func IntersectionOfTwoArrays(nums1 []int, nums2 []int) []int {
	num1Map := make(map[int]bool)
	for _, num := range nums1 {
		num1Map[num] = true
	}
	for _, num := range nums2 {
		if ok := num1Map[num]; ok {
			num1Map[num] = false
		}
	}
	resList := make([]int, 0, len(num1Map))
	for num := range num1Map {
		if !num1Map[num] {
			resList = append(resList, num)
		}
	}
	return resList
}

func IntersectionOfTwoArraysBruteForce(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1))
	for _, tmp := range nums1 {
		if arrayUtil.InIntArray(tmp, nums2) {
			res = append(res, tmp)
		}
	}
	return res
}
