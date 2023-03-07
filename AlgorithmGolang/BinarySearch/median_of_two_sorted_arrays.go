package BinarySearch

import "sort"

// * 给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
// * 算法的时间复杂度应该为 O(log (m+n)) 。

// * nums1.length == m
// * nums2.length == n
// * 0 <= m <= 1000
// * 0 <= n <= 1000
// * 1 <= m + n <= 2000
// * -10^6 <= nums1[i], nums2[i] <= 10^6

// * 二分查找，取mid1，mid2，假设m+n为奇数，此时中位数唯一，令nums1长度为奇数，nums2长度为偶数，假设nums1长度2m+1，nums2长度2n
// * 中位数大于等于全部数字的数量需要至少为m+n
// * 情况1：nums1[mid1] < nums2[mid2]，此时中位数在nums1的右半侧或者nums2左半侧
// ? 证明1：令中位数为nm，坐标为im，如果在nums2右半侧
// ? 此时im大于等于一半的nums2，也必然大于一半的nums1，其坐标将要大于中位数本应该在的坐标，左侧同理
// ? 如果nums1[mid1]可能为中位数，则需要nums1[mid1]>=nums2[mid2-1]，左区间应该包含mid
// ? 如果nums2[mid2]可能为中位数，则需要nums1[mid1+1]>=nums2[mid2]，右区间应该包含mid
// * 情况2：nums1[mid1] > nums2[mid2]，此时中位数在nums1的左半侧或者nums2右半侧
// * 情况3：nums1[mid1] = nums2[mid2]，此时nums2[mid2]即为中位数
// ? 此时nm大于了一半的nums2，大于了一半的nums1，此时nm该中位数大于了nums1的m个数字(0,1,...,m-1)，并且大于了nums2的n个数字（0,1,...,n-1）

// * 进入到下一个循环，此时nums1的搜索长度为m+1，nums2的搜索长度为n，那么此时剩余两个区间的中位数就是答案的中位数吗？
// ? nums1被排除掉m个数字，nums2被排除掉n个数字，应该在剩余两个区间找到第n+1大的数字，此时才满足要求，那么将不构成循环

// ! 注意二分搜索常见的error case， 双元素数组

//func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
//	left1, right1, left2, right2 := 0, len(nums1)-1, 0, len(nums2)-1
//
//	if (len(nums1)+len(nums2))%2 == 1 {
//		// 统计nums1和nums2搜索区间一定大于数字的数量
//		nums1CountGreaterThan, nums2CountGreaterThan := 0, 0
//		targetCountGreaterThan := (len(nums1) + len(nums2) - 1) >> 1
//		for {
//			mid1, mid2 := left1+(right1-left1)>>1, left2+(right2-left2)>>1
//			if nums1[mid1] < nums2[mid2] {
//				nums1CountGreaterThan = mid1
//				left1, right2 = mid1, mid2
//			} else if nums1[mid1] > nums2[mid2] {
//				nums2CountGreaterThan = mid2
//				right1, left2 = mid1, mid2
//			} else {
//				return float64(nums1[mid1])
//			}
//		}
//	}
//
//	return -1
//}

func MedianOfTwoSortedArraysBruteForce(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	}
	return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2+1]) / 2
}
