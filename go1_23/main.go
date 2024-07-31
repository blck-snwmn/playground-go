package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	iterSample()
	slicesSample()
	mapSample()
}

func slicesSample() {
	data := []string{
		"hello", "world", "!",
	}
	fmt.Println("======= slices.All =======")
	for _, x := range slices.All(data) {
		fmt.Printf(x)
	}

	fmt.Println("======= slices.Values =======")
	for x := range slices.Values(data) {
		fmt.Println(x)
	}

	fmt.Println("======= slices.Backward =======")
	for _, x := range slices.Backward(data) {
		fmt.Println(x)
	}

	fmt.Println("======= slices.AppendSeq =======")
	for _, x := range slices.AppendSeq(data, slices.Values([]string{"good", "bye"})) {
		fmt.Println(x)
	}

	fmt.Println("======= slices.Chunk =======")
	for chunk := range slices.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3) {
		fmt.Println(chunk)
	}

	fmt.Println("======= slices.Collect =======")
	fmt.Println(slices.Collect(slices.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)))
}

func mapSample() {
	data := map[string]string{
		"hello": "world",
		"good":  "bye",
	}
	fmt.Println("======= maps.All =======")
	for k, v := range maps.All(data) {
		fmt.Println(k, v)
	}

	fmt.Println("======= maps.Keys =======")
	for k := range maps.Keys(data) {
		fmt.Println(k)
	}

	fmt.Println("======= maps.Values =======")
	for v := range maps.Values(data) {
		fmt.Println(v)
	}

	fmt.Println("======= maps.Insert =======")

	maps.Insert(data, func(yield func(string, string) bool) {
		yield("like", "this")
	})
	for k, v := range maps.All(data) {
		fmt.Println(k, v)
	}

	fmt.Println("======= maps.Collect =======")
	m := maps.Collect(func(yield func(string, int) bool) {
		f := func(k string) {
			yield(k, len(k))
		}
		f("hello world!")
		f("i like this. how about you?")
	})
	for k, v := range m {
		fmt.Println(k, v)
	}
}
