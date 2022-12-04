package BinarySearch

// Core function of Binary Search
func condition(num, n int) bool {
	// ? Why can not be num > n
	return num >= n
}

// BinarySearch Template for Binary Search
// Find index of number from sorted array, return -1 if the number we find is not exist
// todo recite it
func BinarySearch(arr []int, n int) int {
	// ? How to initialize the boundary variable `left` and `right`?
	left, right := 0, len(arr)-1
	// ? When to exit the loop?
	// ? When we use `left < right` and `left <= right` as the while loop condition?
	for left < right {
		// ? why not mid := (left + right) >> 1
		// * overflow, golang int range -2^63 - 2^63-1,
		// * when left + right > 2^63-1, overflow happened
		mid := left + (right-left)>>1
		if condition(arr[mid], n) {
			// ? What happened if right = mid-1 or right = mid+1
			right = mid
		} else {
			// ? How to update the boundary
			// ? How to choose the appropriate combination from `left = mid`, `left = mid + 1` and `right = mid`, `right = mid - 1`?
			// * 取决于循环中止条件，
			// ? What happened if left = mid-1 or right = mid+1
			// * left = mid-1, then range [new_left(mid-1), right] and [old_left, mid] have element new_left(mid-1) in common which is unnecessary.
			// * Vice versa right = mid+1
			left = mid + 1
		}
	}
	if arr[left] == n {
		return left
	}
	return -1
}
