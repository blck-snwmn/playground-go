package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Printf("time: %v\n", time.Since(now))
	}()
	// bucket: 10
	// rate: 1/2s
	l := rate.NewLimiter(2, 10)

	ctx := context.Background()

	for i := 0; i < 11; i++ {
		if err := l.Wait(ctx); err != nil {
			panic(err)
		}

		fmt.Printf("%d\n", i)
	}
}
