package main

import (
	"fmt"
	"iter"
)

func main() {
	fmt.Println("======= mapf =======")
	for i := range mapf(gen(3), func(x int) int { return x * 2 }) {
		fmt.Printf("[map]%d\n", i)
	}

	fmt.Println("======= filter =======")
	for i := range filter(gen(10), func(x int) bool { return x%2 == 0 }) {
		fmt.Printf("[filter]%d\n", i)
	}

	fmt.Println("======= take =======")
	for i := range take(gen(10), 3) {
		fmt.Printf("[take]%d\n", i)
	}

	fmt.Println("======= skip =======")
	for i := range skip(gen(10), 3) {
		fmt.Printf("[skip]%d\n", i)
	}

	fmt.Println("======= anyf =======")
	fmt.Printf("[anyf]: %v\n", anyf(gen(10), func(x int) bool { return x > 4 }))

	fmt.Println("======= allf =======")
	fmt.Printf("[allf]: %v\n", allf(gen(10), func(x int) bool { return x < 4 }))
}

func mapf[T, S any](seq iter.Seq[T], f func(T) S) iter.Seq[S] {
	return func(yield func(S) bool) {
		seq(func(v T) bool {
			return yield(f(v))
		})
	}
	// return func(yield func(S) bool) {
	// 	for v := range seq {
	// 		yield(f(v))
	// 	}
	// }
}

func filter[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(v T) bool {
			if f(v) {
				return yield(v)
			}
			return true
		})
	}
}

func take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		seq(func(v T) bool {
			if count == n {
				return false
			}
			count++
			return yield(v)
		})
	}
}

func skip[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		seq(func(v T) bool {
			if count < n {
				count++
				return true
			}
			return yield(v)
		})
	}
}

func anyf[T any](seq iter.Seq[T], f func(T) bool) bool {
	for i := range seq {
		if f(i) {
			return true
		}
	}
	return false
}

func allf[T any](seq iter.Seq[T], f func(T) bool) bool {
	for i := range seq {
		if !f(i) {
			return false
		}
	}
	return true
}

func gen(end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range end {
			fmt.Printf("\t[gen]: %d\n", i)
			if !yield(i) {
				return
			}
		}
	}
}

func seqFromSlice[T any](input []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}
}
