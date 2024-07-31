package main

import (
	"fmt"
	"iter"
)

func iterSample() {
	fmt.Println("======= mapf =======")
	for i := range mapf(gen(3), func(x int) int { return x * 2 }) {
		fmt.Printf("[map]: %d\n", i)
	}

	fmt.Println("======= filter =======")
	for i := range filter(gen(10), func(x int) bool { return x%2 == 0 }) {
		fmt.Printf("[filter]: %d\n", i)
	}

	fmt.Println("======= flatMap =======")
	for i := range flatMap(
		seqFromSlice([]string{"hello", "world", "!"}),
		func(s string) iter.Seq[rune] {
			return func(yield func(rune) bool) {
				for _, r := range []rune(s) {
					yield(r)
				}
			}
		},
	) {
		fmt.Printf("[flatMap]: %c\n", i)
	}

	fmt.Println("======= filterMap =======")
	for i := range filterMap(
		seqFromSlice([]string{"hello", "", "world", "", "!", "!!"}),
		func(s string) option[string] {
			if len(s) < 2 {
				return option[string]{ok: false}
			}
			return option[string]{s, true}
		},
	) {
		fmt.Printf("[filterMap]: %s\n", i)
	}

	fmt.Println("======= findMap =======")
	v, ok := findMap(
		seqFromSlice([]string{"hello", "world", "!"}),
		func(s string) (string, error) {
			if s == "world" {
				return s, nil
			}
			return "", fmt.Errorf("not found")
		},
	)
	fmt.Printf("[findMap]: %s, %v\n", v, ok)

	fmt.Println("======= take =======")
	for i := range take(gen(10), 3) {
		fmt.Printf("[take]: %d\n", i)
	}

	fmt.Println("======= skip =======")
	for i := range skip(gen(10), 3) {
		fmt.Printf("[skip]: %d\n", i)
	}

	fmt.Println("======= concat =======")
	for i := range concat(gen(100), seqFromSlice([]string{"x", "y", "z", "xx", "yy", "zz"})) {
		fmt.Printf("[concat]: %#v\n", i)
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

func flatMap[T, S any](seq iter.Seq[T], f func(T) iter.Seq[S]) iter.Seq[S] {
	return func(yield func(S) bool) {
		for vs := range mapf(seq, f) {
			for v := range vs {
				if !yield(v) {
					return
				}
			}
		}
	}
}

type option[T any] struct {
	value T
	ok    bool
}

func filterMap[T, S any](seq iter.Seq[T], f func(T) option[S]) iter.Seq[S] {
	return func(yield func(S) bool) {
		seq(func(v T) bool {
			opt := f(v)
			if opt.ok {
				return yield(opt.value)
			}
			return true
		})
	}
}

func findMap[T, S any](seq iter.Seq[T], f func(T) (S, error)) (S, bool) {
	for i := range seq {
		v, err := f(i)
		if err == nil {
			return v, true
		}
	}
	var s S
	return s, false
}

type concated[L, R any] struct {
	left  L
	right R
}

func concat[L, R any](lseq iter.Seq[L], rseq iter.Seq[R]) iter.Seq[concated[L, R]] {
	return func(yield func(concated[L, R]) bool) {
		lv, lstop := iter.Pull(lseq)
		defer lstop()

		rv, rstop := iter.Pull(rseq)
		defer rstop()

		for {
			l, lmore := lv()
			r, rmore := rv()

			if !lmore || !rmore {
				// if either of the sequence is done, then we are done
				return
			}

			if !yield(concated[L, R]{l, r}) {
				return
			}
		}
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
