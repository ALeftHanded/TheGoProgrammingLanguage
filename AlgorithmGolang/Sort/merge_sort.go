package Sort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) >> 1
	leftList, rightList := MergeSort(arr[:mid]), MergeSort(arr[mid:])
	// leftList, rightList := MergeSort(arr[:mid+1]), MergeSort(arr[mid+1:])
	// it will not make it...
	// because when mid is 1, mid+1 = 2
	// then there will be infinite loop in MergeSort(arr[:mid+1])
	return MergeSortedLists(leftList, rightList)
}

func MergeSortedLists(leftList, rightList []int) []int {
	var result []int
	leftIndex, rightIndex := 0, 0

	// Iterate until one of the lists is completely merged
	for leftIndex < len(leftList) && rightIndex < len(rightList) {
		// If the current element of leftList is less or equal, append it to the result
		if leftList[leftIndex] <= rightList[rightIndex] {
			result = append(result, leftList[leftIndex])
			leftIndex++
		} else {
			// Otherwise, append the current element of rightList to the result
			result = append(result, rightList[rightIndex])
			rightIndex++
		}
	}

	// Append any remaining elements from both lists (one of these will be empty)
	result = append(result, leftList[leftIndex:]...)
	result = append(result, rightList[rightIndex:]...)

	return result
}
