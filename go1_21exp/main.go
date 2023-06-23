package main

import (
	"fmt"
	"sync"
)

func main() {
	var sg sync.WaitGroup
	for i := 0; i < 10; i++ {
		sg.Add(1)
		go func() {
			defer sg.Done()
			fmt.Printf("num=%d\n", i)
		}()
	}
	sg.Wait()
}
