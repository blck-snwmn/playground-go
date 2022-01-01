package main

import (
	"fmt"
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

	// 同時実行は制限される
	go print(&g, "key1", "a")
	go print(&g, "key1", "b")
	go print(&g, "key2", "c")

	time.Sleep(time.Second * 5)

	g.Forget("key1")

	// Forget を実行すると指定したキーはリセット
	go print(&g, "key1", "a2")

	time.Sleep(time.Second * 5)

	// 同期処理の場合、順次実行されるので、キャッシュされた値が変えることはない
	print(&g, "key3", "31")
	print(&g, "key3", "32")
}
