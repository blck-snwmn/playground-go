package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		fmt.Println(i)
	}
	fmt.Println("----")
	for i := range rangeOverFunc {
		if i < 40 {
			continue
		}
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
			fmt.Println("break in f")
			return
		}
	}
}
