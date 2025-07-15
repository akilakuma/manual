package main

import (
	"testing"
)

func BenchmarkCheckVgaPlayer1(b *testing.B) {
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		CheckVgaPlayer1(code, pattern)
	}
}

func BenchmarkCheckVgaPlayer2(b *testing.B) {
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		CheckVgaPlayer2(code, pattern)
	}
}

func BenchmarkCheckVgaPlayer1Miss(b *testing.B) {
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		CheckVgaPlayer1(code2, pattern)
	}
}

func BenchmarkCheckVgaPlayer2Miss(b *testing.B) {
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		CheckVgaPlayer2(code2, pattern)
	}
}
