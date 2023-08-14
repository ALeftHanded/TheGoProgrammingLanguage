package Sort

import (
	"reflect"
	"sort"
	"testing"

	"AlgorithmGolang/libs/measureUtil"
	"AlgorithmGolang/libs/random"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}

	randomList := random.GenRandomNums(10000000, 10000000, -100000123, 100000123)
	sortedInts := make([]int, len(randomList))
	tmpStore := make([]int, len(randomList))
	copy(sortedInts, randomList)
	copy(tmpStore, randomList)
	sort.Slice(sortedInts, func(i, j int) bool {
		return sortedInts[i] < sortedInts[j]
	})

	tests := []struct {
		name        string
		args        args
		want        []int
		timeMeasure bool
		memMeasure  bool
	}{
		{
			name:        "cost compare",
			args:        args{randomList},
			want:        sortedInts,
			timeMeasure: true,
			memMeasure:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.timeMeasure {
				measureUtil.ExecutionTime(MergeSort, tt.args.arr)
				copy(tt.args.arr, tmpStore)
				measureUtil.ExecutionTime(HeapSortV2, tt.args.arr)
				copy(tt.args.arr, tmpStore)
				//measureUtil.ExecutionTime(HeapSortBadInit, tt.args.arr)
				//copy(tt.args.arr, tmpStore)
				measureUtil.ExecutionTime(HeapSortV3, tt.args.arr)
				copy(tt.args.arr, tmpStore)
				measureUtil.ExecutionTime(HeapSort, tt.args.arr)
				copy(tt.args.arr, tmpStore)
			}

			if tt.memMeasure {
				measureUtil.MemoryCost(MergeSort, tt.args.arr)
				copy(tt.args.arr, tmpStore)
				measureUtil.MemoryCost(HeapSortV2, tt.args.arr)
				copy(tt.args.arr, tmpStore)
				//measureUtil.MemoryCost(HeapSortBadInit, tt.args.arr)
				//copy(tt.args.arr, tmpStore)
				measureUtil.MemoryCost(HeapSortV3, tt.args.arr)
				copy(tt.args.arr, tmpStore)
				measureUtil.MemoryCost(HeapSort, tt.args.arr)
				copy(tt.args.arr, tmpStore)
			}

			if got := HeapSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
