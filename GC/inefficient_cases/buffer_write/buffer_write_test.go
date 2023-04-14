package buffer_write

import (
	"testing"
)

func BenchmarkProcessData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := generateRandomData(1024)
		processData(data)
	}
}

func BenchmarkProcessDataOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := generateRandomData(1024)
		processDataOptimized(data)
	}
}

// * go test -bench=. -benchmem
//Function							 Operations		Time per operation	Allocated memory per operation

//BenchmarkProcessData-16             	    42549	        29634 ns/op	    3248 B/op, 6 allocs/op
//BenchmarkProcessDataOptimized-16    	    40528	        27854 ns/op	    1024 B/op, 1 allocs/op
//PASS
//ok  	AlgorithmGolang/GC/inefficient_cases/buffer_write	3.560s
