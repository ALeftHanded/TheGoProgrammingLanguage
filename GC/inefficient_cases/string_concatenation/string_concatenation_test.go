package string_concatenation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AlgorithmGolang/libs/arrayUtil"
	"AlgorithmGolang/libs/random"
)

var benchmarkStrs = random.GenerateRandomStringSlice(1000, 2000, 100, 200)
var iStrs = arrayUtil.ConvertStringSliceToInterface(benchmarkStrs)

func TestConcatenateStrings(t *testing.T) {
	baseLine := ConcatenateStringsWithPlus(benchmarkStrs)
	assert.Equal(t, baseLine, ConcatenateStringsWithBuilder(benchmarkStrs))
	assert.Equal(t, baseLine, ConcatenateStringsWithSprintf(iStrs))
	assert.Equal(t, baseLine, ConcatenateStringsWithJoin(benchmarkStrs))
	assert.Equal(t, baseLine, ConcatenateStringsWithBuilderAndGrow(benchmarkStrs))
}

func BenchmarkConcatenateStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsWithPlus(benchmarkStrs)
	}
}

func BenchmarkConcatenateStringsWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsWithBuilder(benchmarkStrs)
	}
}

func BenchmarkConcatenateStringsWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsWithJoin(benchmarkStrs)
	}
}

func BenchmarkConcatenateStringsWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsWithSprintf(iStrs)
	}
}

func BenchmarkConcatenateStringsWithBuilderAndGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsWithBuilderAndGrow(benchmarkStrs)
	}
}

// * go test -bench=. -benchmem
//BenchmarkConcatenateStrings-16                              6445            173310 ns/op         1379144 B/op        129 allocs/op
//BenchmarkConcatenateStringsWithBuilder-16                 101217             11828 ns/op           85216 B/op         14 allocs/op
//BenchmarkConcatenateStringsWithJoin-16                    314462              3839 ns/op           20480 B/op          1 allocs/op
//BenchmarkConcatenateStringsWithSprintf-16                 215265              5731 ns/op           20887 B/op          2 allocs/op
//BenchmarkConcatenateStringsWithBuilderAndGrow-16          350841              3303 ns/op           20480 B/op          1 allocs/op
