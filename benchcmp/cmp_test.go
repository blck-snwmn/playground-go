package main

import (
	"crypto/rand"
	// "math/rand"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := make([]byte, 100)
		rand.Read(b)
	}
}

func BenchmarkSample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := make([]byte, 1000)
		rand.Read(b)
	}
}
