package main

import (
	"context"
	"fmt"
	"time"

	"github.com/canonical/ctxtime"
)

func main() {
	ctx := context.Background()
	now := time.Now()
	fmt.Println(now)

	ctx = ctxtime.ContextWithTime(ctx, now)
	time.Sleep(time.Second)

	now2 := ctxtime.Now(ctx)
	fmt.Println(now2)
}
