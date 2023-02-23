package Array

import (
	"runtime"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	Array "AlgorithmGolang/Utils/array"
	"AlgorithmGolang/Utils/random"
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
			res := IntersectionOfTwoArraysMap(tt.args.nums1, tt.args.nums2)

			for i := 0; i < statisticsTimes; i++ {
				// time cost
				start := time.Now()
				_ = IntersectionOfTwoArraysMap(tt.args.nums1, tt.args.nums2)
				dur := time.Since(start)
				t.Logf("IntersectionOfTwoArraysMap | Actually function duration: %v", dur)
				timeTotalCount += uint64(dur)

				// memory cost
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				beforeAllocs := m.Mallocs
				beforeTotalAlloc := m.TotalAlloc

				// Call the function you want to test
				IntersectionOfTwoArraysMap(tt.args.nums1, tt.args.nums2)

				runtime.ReadMemStats(&m)
				afterAllocs := m.Mallocs
				afterTotalAlloc := m.TotalAlloc

				// Calculate the memory usage of the function
				allocs := afterAllocs - beforeAllocs
				allocsTotalCount += allocs
				totalAlloc := afterTotalAlloc - beforeTotalAlloc
				t.Logf("Allocs: %d, Total Alloc: %d", allocs, totalAlloc)
			}
			t.Logf("Average Allocs: %d", allocsTotalCount/uint64(statisticsTimes))
			t.Logf("Average times: %v", time.Duration(timeTotalCount/uint64(statisticsTimes)))

			tt.want = IntersectionOfTwoArraysBruteForce(tt.args.nums1, tt.args.nums2)
			sort.Ints(Array.Deduplicate(tt.want))
			sort.Ints(res)

			assert.Equalf(t, tt.want, res, "IntersectionOfTwoArraysMap(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
