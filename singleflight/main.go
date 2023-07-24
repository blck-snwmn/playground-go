package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func print(g *singleflight.Group, key, value string) {
	v, _, shared := g.Do(key, func() (interface{}, error) {
		time.Sleep(time.Second)
		fmt.Printf("[key=%v] value=%v (in Do)\n", key, value)
		return value, nil
	})
	fmt.Printf("[key=%v] value=%v (shared=%v, input is %v)\n", key, v, shared, value)
}

func main() {
	var g singleflight.Group

	var sg sync.WaitGroup
	sg.Add(3)

	// 同時実行の場合は処理がまとめられる
	go func() {
		defer sg.Done()
		print(&g, "key1", "a")
	}()
	go func() {
		defer sg.Done()
		print(&g, "key1", "b")
	}()
	go func() {
		defer sg.Done()
		print(&g, "key2", "c")
	}()

	sg.Wait()

	fmt.Println("forget key1")
	g.Forget("key1")

	// Forget を実行すると指定したキーはリセット
	go print(&g, "key1", "a2")

	time.Sleep(time.Second * 5)

	// 同期処理の場合、順次実行されるので、その都度リクエストされる
	print(&g, "key3", "31")
	print(&g, "key3", "32")
}
