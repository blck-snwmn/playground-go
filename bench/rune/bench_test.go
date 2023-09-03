package rune

import "testing"

func Benchmark_bytes(b *testing.B) {
	input := "一二三四五六七八九十"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, r := range input {
			_ = r
		}
	}
}

func Benchmark_rune(b *testing.B) {
	input := "一二三四五六七八九十"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, r := range []rune(input) {
			_ = r
		}
	}
}
