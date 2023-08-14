package Array

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	arrayUtil "AlgorithmGolang/libs/arrayUtil"
	"AlgorithmGolang/libs/measureUtil"
	"AlgorithmGolang/libs/random"
)

func TestIntersectionOfTwoArraysBruteForce(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}

	randNum1 := random.GenRandomIncrIntList(2000, 2000, -1000, 20000000)
	randNum2 := random.GenRandomIncrIntList(3000, 3000, -1000, 20000000)
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal case",
			args: args{
				nums1: randNum2,
				nums2: randNum1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var allocsTotalCount, timeTotalCount uint64
			statisticsTimes := 100
			res := IntersectionOfTwoArrays(tt.args.nums1, tt.args.nums2)

			for i := 0; i < statisticsTimes; i++ {
				// time cost
				_, dur := measureUtil.ExecutionTime(IntersectionOfTwoArrays, tt.args.nums1, tt.args.nums2)
				timeTotalCount += uint64(dur)

				// memory cost
				_, allocs := measureUtil.MemoryCost(IntersectionOfTwoArrays, tt.args.nums1, tt.args.nums2)
				allocsTotalCount += allocs
			}
			t.Logf("Average Allocs: %d", allocsTotalCount/uint64(statisticsTimes))
			t.Logf("Average times: %v", time.Duration(timeTotalCount/uint64(statisticsTimes)))

			tt.want = IntersectionOfTwoArraysBruteForce(tt.args.nums1, tt.args.nums2)
			sort.Ints(arrayUtil.Deduplicate(tt.want))
			sort.Ints(res)

			assert.Equalf(t, tt.want, res, "IntersectionOfTwoArraysMap(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
