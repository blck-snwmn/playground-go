package main

import (
	"fmt"
	"time"
)

func main() {
	chl := make(chan int)
	chr := make(chan int)

	go func() {
		chl <- 1
		chl <- 2
		time.Sleep(time.Millisecond * 100)
		chl <- 3
		close(chl)
	}()

	go func() {
		chr <- 4
		chr <- 5
		chr <- 6
		close(chr)
	}()

	ch := merge(chl, chr)
	for v := range ch {
		fmt.Println(v)
	}
	time.Sleep(time.Second)
}

func merge(chl, chr chan int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		count := 0
		for chl != nil || chr != nil {
			count++
			select {
			case v, open := <-chl:
				fmt.Println("chl", count)
				if !open {
					fmt.Println("\tchl closed", count)
					chl = nil // use ni channel
					break
				}
				ch <- v
			case v, open := <-chr:
				fmt.Println("chr", count)
				if !open {
					fmt.Println("\tchr closed", count)
					chr = nil // use ni channel
					break
				}
				ch <- v
			}
		}
	}()
	return ch
}
