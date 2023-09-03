package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"runtime"
)

func main() {
	PrintMemUsage()
	n := 1000000

	var up string
	flag.StringVar(&up, "m", "p", "")
	flag.Parse()

	switch up {
	case "p":
		do(n, func(n int) map[int]*[128]byte {
			m := make(map[int]*[128]byte)
			for j := 0; j < n; j++ {
				var rb [128]byte
				randBytes(rb[:])
				m[j] = &rb
			}
			return m
		})
	case "s":
		do(n, func(n int) map[int][128]byte {
			m := make(map[int][128]byte)
			for j := 0; j < n; j++ {
				var rb [128]byte
				randBytes(rb[:])
				m[j] = rb
			}
			return m
		})
	case "big":
		do(n, func(n int) map[int][129]byte {
			m := make(map[int][129]byte)
			for j := 0; j < n; j++ {
				// size が 128 を超える場合
				// map[int]*[128]byte と 同様の挙動をする
				var rb [129]byte
				randBytes(rb[:])
				m[j] = rb
			}
			return m
		})
	case "pbig":
		do(n, func(n int) map[int]*[129]byte {
			m := make(map[int]*[129]byte)
			for j := 0; j < n; j++ {
				// size が 128 を超える場合
				// map[int]*[128]byte と 同様の挙動をする
				var rb [129]byte
				randBytes(rb[:])
				m[j] = &rb
			}
			return m
		})
	default:
		panic("unknown")
	}
}

func do[T any](n int, gen func(n int) map[int]T) {
	m := gen(n)
	runtime.GC()
	PrintMemUsage()

	clear(m)

	runtime.GC()
	PrintMemUsage()
	runtime.KeepAlive(m)
}

func randBytes(b []byte) {
	rand.Read(b[:])
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
}
