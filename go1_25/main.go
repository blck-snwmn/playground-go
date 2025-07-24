package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	{
		start := time.Now()

		var wg sync.WaitGroup
		wg.Go(func() {
			// Simulate some work
			time.Sleep(1 * time.Second)
		})

		wg.Go(func() {
			time.Sleep(2 * time.Second)
		})

		wg.Wait()

		since := time.Since(start)
		if since >= 3*time.Second {
			fmt.Println("Total execution time:", since)
		}
	}
}
