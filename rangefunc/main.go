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

	for i := range callDefer {
		fmt.Printf("[callDefer]: %d\n", i)
	}

	p := panicer{}
	for i := range p.iter {
		fmt.Printf("[panicer]: %d\n", i)
	}
	// p.yield(100) // panic because yield becomes nil when the loop ends
}

// rangeOverFunc is a function that return single value
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

// randRange2 is a function that return two values
func randRange2(yield func(int, int) bool) {
	for {
		if !yield(rand.Int(), rand.Int()) {
			break
		}
	}
}

// callDefer is a function that has defer
func callDefer(yield func(int) bool) {
	defer func() {
		fmt.Println("[callDefer in iter func] deferred")
		yield(100)
		fmt.Println("[callDefer in iter func] deferred end")
	}()
	for i := range 10 {
		if !yield(i) {
			return
		}
	}
	fmt.Println("[callDefer in iter func] end")
}

// gen returns a `range-over function` that generates values with arguments.
func gen(end int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := range end {
			if !yield(i) {
				return
			}
		}
	}
}

// sortIter returns a `range-over function` that sorts an array of arguments and returns their values.
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

// invalid is a panic sample
func invalid(yield func(int) bool) {
	for i := range 100 {
		if !yield(i) {
			fmt.Println("break in f")
			// return // panic
		}
	}
}

type panicer struct {
	yield func(int) bool
}

func (p panicer) iter(yield func(int) bool) {
	p.yield = yield
	if p.yield == nil {
		fmt.Println("yield is nil")
	}
	defer func() {
		if p.yield == nil {
			fmt.Println("yield is nil in defer")
		}
	}()
	for i := range 10 {
		if !yield(i) {
			return
		}
	}
}
