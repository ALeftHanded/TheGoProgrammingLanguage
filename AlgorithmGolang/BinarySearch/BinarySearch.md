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