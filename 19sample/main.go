package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Add Duration.Abs method
func sampleTime() {
	l := time.Date(2022, 12, 1, 12, 13, 10, 0, time.Local)
	r := time.Date(2022, 12, 4, 12, 13, 10, 0, time.Local)
	if l.Sub(r).Abs() == r.Sub(l).Abs() {
		fmt.Println("|l-r|==|r-l|")
	}
}

func sampleAtmic() {
	var i uint64
	atomic.AddUint64(&i, 100)
	fmt.Println(atomic.LoadUint64(&i))

	var ii atomic.Uint64
	// use atomic.AddUint64 in Add method
	ii.Add(101)
	fmt.Println(ii.Load())
	ii.Add(2)
	fmt.Println(ii.Load())
}

func main() {
	sampleTime()
	sampleAtmic()
}
