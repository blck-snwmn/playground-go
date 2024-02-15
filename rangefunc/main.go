package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	for i := range rangeOverFunc {
		if i == 5 {
			fmt.Println("break in for loop")
			break
		}
		fmt.Printf("[main]: %d\n", i)
	}
	fmt.Println("-------------------")
	for i := range gen(3) {
		fmt.Println(i)
	}

	for i, v := range sortIter([]int{19, 100, 2, 7, 5, 50}) {
		fmt.Printf("%d: %d\n", i, v)
	}
}

func rangeOverFunc(yield func(int) bool) {
	for i := range 10 {
		fmt.Printf("[rangeOverFunc]: %d\n", i)
		if !yield(i) {
			//  the function must exit if yield returns false
			fmt.Println("break in f")
			return
		}
	}
}

func gen(end int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := range end {
			if !yield(i) {
				return
			}
		}
	}
}

func sortIter[T cmp.Ordered](input []T) func(func(int, T) bool) {
	slices.Sort(input)
	return func(yield func(int, T) bool) {
		for i, v := range input {
			if !yield(i, v) {
				return
			}
		}
	}
}

func invalid(yield func(int) bool) {
	for i := range 100 {
		if !yield(i) {
			fmt.Println("break in f")
			// return // panic
		}
	}
}
