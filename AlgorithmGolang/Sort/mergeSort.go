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
	return MergeSortedList(leftList, rightList)
}

func MergeSortedList(leftList, rightList []int) []int {
	var res []int
	l, r := 0, 0
	for {
		if l == len(leftList) {
			res = append(res, rightList[r:]...)
			break
		}
		if r == len(rightList) {
			res = append(res, leftList[l:]...)
			break
		}
		if leftList[l] <= rightList[r] {
			res = append(res, leftList[l])
			if l < len(leftList) {
				l++
			}
		} else {
			res = append(res, rightList[r])
			if r < len(rightList) {
				r++
			}
		}
	}
	return res
}
