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
	// return max.Int(initDPForLIS(nums)...)
	return len(initTails(nums))
}

func initDPForLIS(nums []int) []int {
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

// * Binary search + DP

// * tails[k]存储长度为k+1的所有LIS的最小末位数字
// * 让末位数字尽可能小，这样找到下一个递增数字的可能性会更大

// * tails的最终长度即为LIS的最大长度，因为我们要找到最长的LIS才能完成这个数组的初始化

// * tails 单增，易证
// * tails[k]对应LIS: {a1, a2, ... , ak}, tails[k] = ak
// * tails[k - 1]对应LIS: {b1, b2, ... , bk-1}, tails[k-1] = bk-1
// * 根据LIS的单增特性，必然有ak > ak-1
// * 根据tails的定义，必然有ak-1 >= bk-1
// * ak > ak-1 >= bk-1 => ak > bk-1 => tails[k] > tails[k-1]

// ? tails的单增属性有什么用
// * 用来二分搜索nums[k]，以更新tails现存元素或者计算新的tails

func initTails(nums []int) []int {
	tails := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		left, right := 0, len(tails)-1
		for left <= right {
			mid := (left + right) >> 1
			if tails[mid] == nums[i] {
				left = mid
				break
			}
			if tails[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if len(tails) == left {
			// calculate tails[len(tails)]
			tails = append(tails, nums[i])
		} else {
			// update tails[:len(tails)-1]
			tails[left] = nums[i]
		}
	}
	return tails
}
