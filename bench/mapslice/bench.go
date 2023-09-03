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

	up := flag.Bool("p", false, "")
	flag.Parse()

	if *up {
		doUsePointer(n)
	} else {
		doNoPointer(n)
	}
}

func doUsePointer(n int) {
	do(n, func(n int) map[int]*[128]byte {
		m := make(map[int]*[128]byte)
		for j := 0; j < n; j++ {
			rb := randBytes()
			m[j] = &rb
		}
		return m
	})
}

func doNoPointer(n int) {
	do(n, func(n int) map[int][128]byte {
		m := make(map[int][128]byte)
		for j := 0; j < n; j++ {
			rb := randBytes()
			m[j] = rb
		}
		return m
	})
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

func randBytes() [128]byte {
	var b [128]byte
	rand.Read(b[:])
	return b
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
}
