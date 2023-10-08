package DoublePointers

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/submissions/?envType=study-plan-v2&envId=top-interview-150


// 111111 22222 333
// 112111 22222 333
// 112231 22222 333
// 112233 22222 333

// 1234567
func RemoveDuplicatesII(nums []int) int {
	slow, fast,count := 0,0,0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			if count == 1 {
				slow +=1
				nums[slow] = nums[fast]
			} else {
				nums[slow+1] = nums[slow]
				slow +=2
				nums[slow] = nums[fast]
			}
			count = 0
			continue
		}
		fast++
		count++
	}
	if count >= 2 {
		nums[slow+1] = nums[slow]
		slow++
	}
	return slow+1
}