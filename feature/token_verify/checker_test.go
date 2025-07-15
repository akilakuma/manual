package main

import (
	"testing"
)

var (
	validCode   = "ddea02631215"
	invalidCode = "ddea02631315"
	pattern     = "67d8309e4026e15a7a9fa79a0a33ed01e490ff15bea67e8dd474e21e93c5d7cd"
)

func TestCheck(t *testing.T) {
	if !Check(validCode, pattern) {
		t.Errorf("Expected true for valid code")
	}
	if Check(invalidCode, pattern) {
		t.Errorf("Expected false for invalid code")
	}
}

func BenchmarkCheckValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Check(validCode, pattern)
	}
}

func BenchmarkCheckInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Check(invalidCode, pattern)
	}
}
