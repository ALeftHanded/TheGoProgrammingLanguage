package DoublePointers

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/submissions/?envType=study-plan-v2&envId=top-interview-150


// 111111 22222 333
// 112111 22222 333
// 112231 22222 333
// 112233 22222 333

// simplify
func RemoveDuplicatesII(nums []int) int{
    write := 1
    count := 1
	// Start read pointer at 1, as the first element is always considered unique.
    for read := 1; read < len(nums); read++ {
        if nums[read] == nums[read - 1] {
            count++
        } else {
            count = 1
        }
        if count <= 2 {
            nums[write] = nums[read]
            write++
        }
    }
    return write
}

// 链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/solutions/702970/gong-shui-san-xie-guan-yu-shan-chu-you-x-glnq/
// trunk-ignore(golangci-lint/unused)
func removeDuplicatesCommonSolve(nums []int) int {
    // innerFunc 负责去除数组中出现次数超过 maxOccurrence 的元素
    innerFunc := func(maxOccurrence int) int {
        writeIndex := 0  // 用于写入元素到新数组的索引
        for _, currentVal := range nums {
            // 判断条件：确保每个元素至多被保留 maxOccurrence 次
            // 1. 当新数组长度（writeIndex）小于 maxOccurrence 时，直接添加元素。
            // 2. 当新数组中的第 writeIndex-maxOccurrence 个元素与当前元素不同，
            //    则表示当前元素出现次数未超过 maxOccurrence。
            if writeIndex < maxOccurrence || nums[writeIndex-maxOccurrence] != currentVal {
                // 把当前元素添加到新数组
                nums[writeIndex] = currentVal
                writeIndex++
            }
        }
        return writeIndex  // 返回新数组长度
    }
    return innerFunc(2)
}
