package main

import (
	"cmp"
	"fmt"
	"log/slog"
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
}
