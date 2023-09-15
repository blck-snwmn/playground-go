package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	{
		f := func() {
			fmt.Println("once")
		}
		once.Do(f)
	}
	{
		f := func() {
			fmt.Println("once")
		}
		once.Do(f)
	}
}
