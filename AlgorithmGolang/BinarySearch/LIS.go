package BinarySearch

// * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// * 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// * 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

// * 1 <= nums.length <= 2500
// * -104 <= nums[i] <= 104

// * --------

// * 分析
// * DP问题
// * n为数组下标，dp[n]为包含nums[n]的最长严格递增子序列的长度
// * 有dp[0] = 1
// * dp[n] = max([(dp[j] + 1) ** int(nums[n]>nums[j]) for j in range(len(dp[:n]))])

func lengthOfLIS(nums []int) int {
	return max(initDP(nums))
}

func initDP(nums []int) []int {
	// init dp
	dp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		tmpDpRes := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > tmpDpRes {
					tmpDpRes = dp[j] + 1
				}
			}
		}
		dp = append(dp, tmpDpRes)
	}
	return dp
}

func max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
