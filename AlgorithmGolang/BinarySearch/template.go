package BinarySearch

// BinarySearch Template for Binary Search
// Find index of number from sorted arrayUtil, return -1 if the number we find is not exist
// todo recite it
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
