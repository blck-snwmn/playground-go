package copybyte

import (
	"bytes"
	"testing"
)

var src = bytes.Repeat([]byte{0x00, 0xA0}, 20)

func Benchmark_Copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := make([]byte, len(src))
		copy(dst, src)
	}
}

func Benchmark_BytesCone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.Clone(src)
	}
}

func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append([]byte{}, src...)
	}
}
