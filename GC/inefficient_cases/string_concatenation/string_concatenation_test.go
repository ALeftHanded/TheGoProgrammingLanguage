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

func BenchmarkConcatenateStringsOp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateStringsOp(benchmarkStrs)
	}
}

// * go test -bench=. -benchmem

//Function						Operations	Time per operation	Allocated memory per operation
//BenchmarkConcatenateStrings	4172774		291 ns/op			64 B/op, 9 allocs/op
//BenchmarkConcatenateStringsOp	13136090	86.9 ns/op			24 B/op, 2 allocs/op
