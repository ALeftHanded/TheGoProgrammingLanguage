package DoublePointers

// * when there is no additional space for solving, always consider DoublePointer method

// RemoveDuplicates takes a sorted slice of integers `nums` and removes duplicate
// elements in-place, returning the length of the modified slice.
func RemoveDuplicates(nums []int) int {
    // `slow` and `fast` are pointers initialized to track the elements in the slice.
    // `slow` points to the last non-duplicate element, while `fast` scans through the slice.
    slow, fast := 0, 1
    
    // Iterate while `fast` hasn't reached the end of the slice.
    for fast < len(nums) {
        // If the element at `slow` is different from the element at `fast`,
        // then we have found a new unique element.
        if nums[slow] != nums[fast] {
            // Increment `slow` to point to the next position.
            slow++
            
            // Replace the element at the new `slow` index with the element at `fast`.
            nums[slow] = nums[fast]
        }
        
        // Increment `fast` to continue scanning the slice.
        fast++
    }
    
    // Return the length of the slice after duplicates have been removed.
    // Since `slow` is 0-based, we add 1.
    return slow + 1
}