package Sort

import (
	"math"

	"AlgorithmGolang/libs/maxHeapUtil"
)

func HeapSortBadInit(arr []int) []int {
	maxHeap := maxHeapUtil.InitMaxHeapFromIntArrayBad(arr)
	return maxHeap.ToOrderedArray()
}

func HeapSortV2(arr []int) []int {
	maxHeap := maxHeapUtil.HeapifyInit(arr)
	return maxHeap.ToOrderedArray()
}

func HeapSortV3(arr []int) []int {
	// heapify arr
	lastNonLeafNodeIndex := (len(arr) / 2) - 1
	for i := lastNonLeafNodeIndex; i >= 0; i-- {
		maxHeapUtil.CheckAndSwap(arr, i)
	}
	var res []int
	for arr[0] != math.MinInt {
		res = append(res, arr[0])
		arr[0] = math.MinInt
		swapMinInt(arr, 0)
	}
	return res
}

func swapMinInt(arr []int, index int) {
	if index > (len(arr)/2)-1 {
		return
	}
	largestIndex := maxHeapUtil.GetLargestIndex(arr, index)
	if largestIndex != index {
		arr[index], arr[largestIndex] = arr[largestIndex], arr[index]
		swapMinInt(arr, largestIndex)
	}
}

func HeapSort(arr []int) []int {
	// heapify arr
	lastNonLeafNodeIndex := (len(arr) / 2) - 1
	for i := lastNonLeafNodeIndex; i >= 0; i-- {
		maxHeapUtil.CheckAndSwap(arr, i)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapUtil.CheckAndSwap(arr[:i], 0)
	}
	return arr
}
