package main

import "testing"

func BenchmarkStringIndexWay(b *testing.B) {
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		stringIndexWay()
	}
}
