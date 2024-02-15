package main

import (
	"fmt"
)

func main() {
	for i := range rangeOverFunc {
		if i == 50 {
			fmt.Println("break in for loop")
			break
		}
		fmt.Println(i)
	}
}

func rangeOverFunc(yield func(int) bool) {
	for i := range 100 {
		if !yield(i) {
			//  the function must exit if yield returns false
			fmt.Println("break in f")
			return
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
