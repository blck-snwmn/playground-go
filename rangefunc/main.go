package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
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
		fmt.Printf("[gen]: %d\n", i)
	}

	for i, v := range sortIter([]int{19, 100, 2, 7, 5, 50}) {
		fmt.Printf("[sort]%d: %d\n", i, v)
	}

	{
		count := 0
		for i, v := range randRange2 {
			fmt.Printf("[rand:1]%d: %d\n", i, v)
			count++
			if count == 3 {
				break
			}
		}
	}
	{
		count := 0
		for i, v := range randRange2 {
			fmt.Printf("[rand:2]%d: %d\n", i, v)
			count++
			if count == 3 {
				break
			}
		}
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

func randRange2(yield func(int, int) bool) {
	for {
		if !yield(rand.Int(), rand.Int()) {
			break
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
