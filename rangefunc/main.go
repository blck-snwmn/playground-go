package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		fmt.Println(i)
	}
	fmt.Println("----")
	for i := range f {
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

func f(yield func(int) bool) bool {
	yield(51)
	yield(51)
	yield(51)
	yield(51)
	yield(51)
	for i := range 100 {
		if !yield(i) {
			fmt.Println("break in f")
			return false // break
		}
	}
	return false
}
