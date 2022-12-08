## Binary search

Binary search conditions that can apply to:
1. a sorted array
2. search a given element x in array

### Binary search template

> Ref: 
>
> [Binary Search 101](https://leetcode.com/problems/binary-search/discuss/423162/Binary-Search-101)
> 
> [详解二分查找算法](https://www.cnblogs.com/kyoner/p/11080078.html)
>
> [二分查找、二分边界查找算法的模板代码总结](https://segmentfault.com/a/1190000016825704)
>
> [[Python] Powerful Ultimate Binary Search Template. Solved many problems](https://leetcode.com/discuss/general-discussion/786126/python-powerful-ultimate-binary-search-template-solved-many-problems)

#### Some of the most common problems include:

1. When to exit the loop? Should we use `left < right` or `left <= right` as the while loop condition?
2. How to initialize the boundary variable `left` and `right`?
3. How to update the boundary? How to choose the appropriate combination from `left = mid`, `left = mid + 1` and `right = mid`, `right = mid - 1`?

#### template in golang

模版有很多种，这里只展示最通用的一种

其中`left = 0`、`mid = left + (right - left)>>1`总是不会变的



```go
func BinarySearch(arr []int, n int) int {
	// ? How to initialize the boundary variable `left` and `right`?
	// * We just use closed intervals for binary search without a specific reason.
	left, right := 0, len(arr)-1
	// ? Why we use the comparison left <= right rather than left < right?
	// * We just use closed intervals for binary search
	// * Using left <= right ensures that the search will continue until every element in the search space has been considered
	// * While using left < right would cause the search to terminate early if there are an even number of elements in the search space.
	// ! Wrong case example: arr = [5], n = 0, using left < right will result in -1
	for left <= right {
		// ? Why not mid := (left + right) >> 1
		// * overflow, golang int range -2^63 - 2^63-1,
		// * when left + right > 2^63-1, overflow happened
		mid := left + (right-left)>>1
		if arr[mid] == n {
			return mid
		}
		if arr[mid] < n {
			// ? Why not left = mid
			// * We just use closed intervals for binary, so exclude `mid`
			left = mid + 1
		} else {
			// ? Why not right = mid
			// * We just use closed intervals for binary, so exclude `mid`
			right = mid - 1
		}
	}
	return -1
}
```





#### 最左查找模版

python

```python
# search_space: could be an array, a range, etc. Usually it's sorted in ascending order.
# Minimize k , s.t. condition(k) is True
def binary_search(array) -> int:
    def condition(value) -> bool:
        pass

    left, right = min(search_space), max(search_space) # could be [0, n], [1, n] etc. Depends on problem
    while left < right:
        mid = left + (right - left) >> 1
        if condition(mid):
            right = mid
        else:
            left = mid + 1
    return left
```



### Binary search相关的CodeTop频度大于等于20的习题

#### Easy

1. [704. 二分查找](https://leetcode-cn.com/problems/binary-search/)
2. [69. Sqrt(x)](https://leetcode-cn.com/problems/sqrtx/)

#### Medium

1. [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
2. [300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
3. [718. 最长重复子数组](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)
4. [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)
5. [162. 寻找峰值](https://leetcode-cn.com/problems/find-peak-element/)

#### Hard

1. [4. 寻找两个正序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)