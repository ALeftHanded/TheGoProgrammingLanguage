package BinarySearch

import (
	"fmt"
	"reflect"
	"strings"

	"AlgorithmGolang/Utils/max"
	"AlgorithmGolang/Utils/min"
)

// ? 给两个整数数组 nums1 和 nums2 ，返回两个数组中公共的、长度最长的子数组的长度。

// * 1 <= nums1.length, nums2.length <= 1000
// * 0 <= nums1[i], nums2[i] <= 100

// * DP
// * dp[i][j] 为nums1
// * dpIndex[n]为以long[n]元素为末位的最长公共子数组在short中的index list

func MaximumLengthOfRepeatedSubarray(nums1 []int, nums2 []int) int {
	return max.Int(initDPForMLORS(nums1, nums2)...)
}

func initDPForMLORS(nums1 []int, nums2 []int) (dp []int) {
	return dp
}

// * 滑动窗口
// * 较短的数组在较长的数组上面从左到右滑动
// * 优化版本

func MaximumLengthOfRepeatedSubarray3rdAC(nums1 []int, nums2 []int) int {
	var res int
	for i := 0; i < len(nums1); i++ {
		length := min.Int(len(nums1)-i, len(nums2))
		res = max.Int(getLongestCommonSliceLength(nums1, nums2, i, 0, length), res)
	}
	for i := 0; i < len(nums2); i++ {
		length := min.Int(len(nums2)-i, len(nums1))
		res = max.Int(getLongestCommonSliceLength(nums1, nums2, 0, i, length), res)
	}
	return res
}

func getLongestCommonSliceLength(nums1, nums2 []int, addr1, addr2, length int) int {
	var res, count int
	for i := 0; i < length; i++ {
		if nums1[addr1+i] == nums2[addr2+i] {
			count += 1
		} else {
			res = max.Int(res, count)
			count = 0
		}
	}
	res = max.Int(res, count)
	return res
}

// * 滑动窗口
// * 较短的数组在较长的数组上面从左到右滑动
// * 每次滑动都计算公共部分，并且返回完全匹配上的数组的最大长度
// ? 优化项

func MaximumLengthOfRepeatedSubarray2ndAC(nums1 []int, nums2 []int) int {
	short, long := findShortAndLongArray(nums1, nums2)
	res := 0
	for i := 0; i < len(long)+len(short)-1; i++ {
		// ? what is the definition of i
		// * pointer to long and tail of short
		// get intersection left and right index
		var tmp int
		if i < len(short)-1 {
			// * phase 1: short is entering
			tmp = getLongestSubarrayFromCommonSlice(short[len(short)-1-i:], long[:i+1])
		} else if i < len(long) {
			// * phase 2: short is on the long
			tmp = getLongestSubarrayFromCommonSlice(short, long[i-(len(short)-1):i+1])
		} else if len(short)-(i-(len(long)-1))+1 > res {
			// * phase 3: short is exiting

			// i => len(long) + k (len(short)>k>=0)
			// when i=len(long) + len(short) - 2
			// shortSlice should be short[0:1]
			// longSlice should be long[len(long)-1:]
			tmp = getLongestSubarrayFromCommonSlice(short[:len(short)-(i-(len(long)-1))], long[i-(len(short)-1):])

		}
		// upd res
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func findShortAndLongArray(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums1) > len(nums2) {
		return nums2, nums1
	}
	return nums1, nums2
}

func getLongestSubarrayFromCommonSlice(short, long []int) int {
	var res, count int
	for i := 0; i < len(short); i++ {
		if short[i] == long[i] {
			count++
			if count > res {
				res = count
			}
		} else if len(short)-1-i > res {
			count = 0
		} else {
			break
		}

	}
	return res
}

// * 滑动窗口
// * nums1的窗口长度一点点增大来搜索nums2

func MaximumLengthOfRepeatedSubarray1stAC(nums1 []int, nums2 []int) int {
	if !haveCommonElements(nums1, nums2) {
		return 0
	}
	slideLength := 1
	for i := 0; i+slideLength <= len(nums1); i++ {
		tmpSlice := nums1[i : i+slideLength]
		for searchSliceInNums2(tmpSlice, nums2) {
			slideLength += 1
			if i+slideLength > len(nums1) {
				break
			}
			tmpSlice = nums1[i : i+slideLength]

		}
	}
	return slideLength - 1
}

func searchSliceInNums2(s []int, nums2 []int) bool {
	num2Str := fmt.Sprint(nums2)
	sStr := fmt.Sprint(s)
	return strings.Contains(
		fmt.Sprintf(" %s ", num2Str[1:len(num2Str)-1]),
		fmt.Sprintf(" %s ", sStr[1:len(sStr)-1]))
	//// ! need optimize
	//for i := 0; i <= len(nums2)-len(s); i++ {
	//	if reflect.DeepEqual(s, nums2[i:i+len(s)]) {
	//		return true
	//	}
	//}
	//return false
}

func haveCommonElements(nums1 []int, nums2 []int) bool {
	nums1Map := make(map[int]bool)
	for _, num := range nums1 {
		nums1Map[num] = true
	}
	for _, num := range nums2 {
		if nums1Map[num] {
			return true
		}
	}
	return false
}

func BruceForceMLORS(nums1 []int, nums2 []int) int {
	res1 := genAllSubArrayList(nums1)
	res2 := genAllSubArrayList(nums2)
	ans := 0
	for i := 0; i < len(res1); i++ {
		for j := 0; j < len(res2); j++ {
			if reflect.DeepEqual(res1[i], res2[j]) {
				if len(res1[i]) > ans {
					fmt.Println("slice is ", res1[i])
					ans = len(res1[i])
				}
			}
		}
	}
	fmt.Println("nums1 is ", nums1)
	fmt.Println("nums2 is ", nums2)
	fmt.Println("ans is ", ans)
	return ans
}

func genAllSubArrayList(nums []int) [][]int {
	res := make([][]int, 0, len(nums)*(len(nums)-1)/2)
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= len(nums)-i; j++ {
			res = append(res, nums[j:j+i])
		}
	}
	fmt.Println(res)
	return res
}
