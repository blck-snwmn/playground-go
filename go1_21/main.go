package main

import (
	"cmp"
	"fmt"
	"log/slog"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	{ // log/slog
		l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

		l.Info("use json handler", slog.Bool("boolkey", true))
		slog.Info("before SetDefault", slog.Bool("boolkey", true))

		slog.SetDefault(l.With(slog.String("withkey", "withvalue")))

		slog.Warn("warn")
	}
	{
		// cmp
		fmt.Printf("1 cmp 2%d\n", cmp.Compare(1, 2))
		fmt.Printf("1 cmp 1%d\n", cmp.Compare(1, 1))
		fmt.Printf("3 cmp 2%d\n", cmp.Compare(3, 2))
	}
	{
		// slices
		fmt.Printf("max=%d\n", slices.Max([]int{1, 2, 100, 3}))
		fmt.Printf("max=%d\n", slices.Min([]int{1, 2, 0, 3}))
		var (
			index int
			ok    bool
		)
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 15, 18, 20, 21, 22, 23, 24, 25, 26, 27, 40, 50, 60, 70, 80, 90, 100}
		sort.Ints(input)
		fmt.Println(input)
		index, ok = slices.BinarySearch(input, 25)
		fmt.Printf("exists?`%v`, search=%v(index=%d)\n", ok, input[index], index)

		index, ok = slices.BinarySearch(input, 11)
		fmt.Printf("exists?`%v`, search=%v(index=%d)\n", ok, input[index], index)
	}
	{
		// clear
		slice := []int{1, 2, 3, 4, 5}
		fmt.Printf("input=%v\n", slice)
		clear(slice)
		fmt.Printf("input=%v(cleared)\n", slice)

		m := map[string]int{"a": 1, "b": 2}
		fmt.Printf("input=%v\n", m)
		clear(m)
		fmt.Printf("input=%v(cleared)\n", m)

		mNaN := map[float64]int{math.NaN(): 10, math.Inf(0): 100, math.Inf(-1): 20}
		fmt.Printf("input=%v\n", mNaN)
		mNaN[math.NaN()] = 12
		mNaN[math.Inf(0)] = 13
		mNaN[math.Inf(-1)] = 14
		fmt.Printf("input=%v(changed)\n", mNaN) // NaN is duplicated
		clear(mNaN)
		fmt.Printf("input=%v(cleared)\n", mNaN) // clear NaN key
	}
	{
		// strings/bytes
		fmt.Println("===========strings/bytes===========")
		fmt.Printf("ContainsFunc=%v\n", strings.ContainsFunc("abcd", func(r rune) bool {
			return r == 'a'
		}))
		fmt.Printf("ContainsFunc=%v\n", bytes.ContainsFunc([]byte{0x00, 0x02, 0x05}, func(r rune) bool {
			return r%2 == 1
		}))
		fmt.Printf("ContainsFunc=%v\n", bytes.ContainsFunc([]byte{0x00, 0x02, 0x04}, func(r rune) bool {
			return r%2 == 1
		}))
	}
}
