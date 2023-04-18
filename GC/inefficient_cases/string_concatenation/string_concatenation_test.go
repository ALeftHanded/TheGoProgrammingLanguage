package string_concatenation

import (
	"testing"
)

var benchmarkStrs = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func BenchmarkConcatenateStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStrings(benchmarkStrs)
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
		ConcatenateStringsWithSprintf()
	}
}

func BenchmarkConcatenateStringsWithBuilderAndGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsWithBuilderAndGrow(benchmarkStrs)
	}
}

// * go test -bench=. -benchmem

//Function										       Operations	   Time per operation	Allocated memory per operation
//BenchmarkConcatenateStrings-16                      	 3739914	       289 ns/op	      64 B/op	       9 allocs/op
//BenchmarkConcatenateStringsWithBuilder-16           	12547957	        85.5 ns/op	      24 B/op	       2 allocs/op
//BenchmarkConcatenateStringsWithJoin-16              	12004039	        96.5 ns/op	      16 B/op	       1 allocs/op
//BenchmarkConcatenateStringsWithSprintf-16           	 4429815	       239 ns/op	      16 B/op	       1 allocs/op
//BenchmarkConcatenateStringsWithBuilderAndGrow-16    	16140784	        62.1 ns/op	      16 B/op	       1 allocs/op
