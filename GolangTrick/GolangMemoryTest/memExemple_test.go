package main

import "testing"

func BenchmarkMemExample(b *testing.B) {
	b.ReportAllocs()
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MemExample()
	}
	//b.StopTimer()
}
