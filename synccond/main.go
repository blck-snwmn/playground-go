package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	cond  *sync.Cond
}

func main() {
	c := Counter{cond: sync.NewCond(&sync.Mutex{})}

	f := func(limit int) {
		c.cond.L.Lock()
		defer c.cond.L.Unlock()
		for c.value < limit {
			fmt.Printf("\twaiting. limit=%d\n", limit)
			c.cond.Wait()
		}
		fmt.Printf("value: %d. end\n", c.value)
	}

	go f(10)
	go f(20)
	go f(30)

	for i := 0; i < 33; i++ {
		time.Sleep(100 * time.Millisecond)
		c.cond.L.Lock()
		fmt.Printf("incrementing: loop count: %d\n", i)
		c.value++
		c.cond.L.Unlock()
		c.cond.Broadcast()
	}
}
