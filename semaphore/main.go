package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func sleeper(s *semaphore.Weighted, w int64, d time.Duration) {
	if err := s.Acquire(context.Background(), w); err != nil {
		fmt.Println("error1")
		return
	}
	defer s.Release(w)
	fmt.Printf("[now:%v]weight=%d\n", time.Now(), w)
	time.Sleep(d)
}

func main() {
	sem := semaphore.NewWeighted(10)
	go sleeper(sem, 1, time.Second*5)
	go sleeper(sem, 3, time.Second*5)
	time.Sleep(time.Second)
	go sleeper(sem, 7, time.Second*5) // 先に確保した側がReleaseするまで遅延される
	time.Sleep(time.Second)
	go sleeper(sem, 2, time.Second*5)
	time.Sleep(time.Second * 15)
}
