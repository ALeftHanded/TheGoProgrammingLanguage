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
