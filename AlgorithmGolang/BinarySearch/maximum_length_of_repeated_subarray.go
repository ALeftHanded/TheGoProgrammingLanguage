package BinarySearch

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ? 给两个整数数组 nums1 和 nums2 ，返回两个数组中公共的、长度最长的子数组的长度。

// * 1 <= nums1.length, nums2.length <= 1000
// * 0 <= nums1[i], nums2[i] <= 100

// * 滑动窗口
// * nums1的窗口长度一点点增大来搜索nums2

func MaximumLengthOfRepeatedSubarray(nums1 []int, nums2 []int) int {
	if !haveCommonElements(nums1, nums2) {
		return 0
	}
	slideLength := 1
	for i := 0; i+slideLength <= len(nums1); i++ {
		start := time.Now()
		tmpSlice := nums1[i : i+slideLength]
		for searchSliceInNums2(tmpSlice, nums2) {
			fmt.Println("tmpSlice:", tmpSlice)
			slideLength += 1
			if i+slideLength > len(nums1) {
				break
			}
			tmpSlice = nums1[i : i+slideLength]

		}
		dur := time.Since(start)
		fmt.Printf("dur is %s, i is %d\n", dur, i)
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
