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

func f(yield func(int64) bool) bool {
	for i := range 100 {
		i := int64(i)
		if !yield(i) {
			fmt.Println("break in f")
			return false // break
		}
	}
	return false
}
