package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limit := rate.NewLimiter(rate.Every(time.Second), 10)

	ctx := context.Background()
	prev := time.Now()
	start := prev
	for i := 0; i < 20; i++ {
		if err := limit.Wait(ctx); err != nil {
			fmt.Println(err)
			return
		}
		now := time.Now()
		fmt.Println(i, now.Sub(prev), now.Sub(start))
		prev = now
	}
}
