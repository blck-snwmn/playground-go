package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	var (
		ctxTime   = time.Second * 3
		afterTime = time.Second * 5
	)

	ctx, cancel := context.WithTimeout(context.Background(), ctxTime)
	defer cancel()

	now := time.Now()
	select {
	case <-time.After(afterTime):
		fmt.Printf("hello(after)-since:%v\n", time.Since(now))
	case <-ctx.Done():
		fmt.Printf("hello(ctx.Done)-since:%v\n", time.Since(now))
	}
}
