package main

import (
	"fmt"
	"sort"
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

func sampleSort() {
	input := []string{"aaa", "aaa", "aaa", "abc", "acb", "acb", "xxx"}

	for _, target := range []string{"abc", "xxa"} {
		fmt.Printf("-----\nsearch word is %q\n", target)
		fmt.Println("find start")
		fi, ok := sort.Find(len(input), func(i int) int {
			fmt.Printf("\tindex is %d\n", i)
			ii := input[i]
			if target == ii {
				return 0
			}
			if target < ii {
				return -1
			}
			return 1
		})
		fmt.Println("find:", fi, ok)

		fmt.Println("serch start")
		si := sort.Search(len(input), func(i int) bool {
			fmt.Printf("\tindex is %d\n", i)
			return input[i] >= target
		})
		fmt.Println("serch:", si)
	}
}

func sampleFmtAppend() {
	s := "aaaaa"
	b := []byte(s)
	b = fmt.Append(b, "bbb", "c")
	b = fmt.Appendf(b, "zzzz%d", 1)

	fmt.Println(string(b))
}

func wrapper(delimiter string, f func()) {
	const template = "[%s:%s]================\n"
	fmt.Printf(template, delimiter, "start")
	defer fmt.Printf(template, delimiter, "end")
	f()
}

func main() {
	wrapper("time", sampleTime)
	wrapper("atomic", sampleAtmic)
	wrapper("sort", sampleSort)
	wrapper("fmt.append", sampleFmtAppend)
}
