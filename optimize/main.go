package main

import (
	"bytes"
	"fmt"
)

//go:noinline
func greetNoInline() string {
	return "hello alice"
}

func greet() string {
	return "hello bob"
}

func main() {
	{
		s := greet()
		fmt.Println(s)
	}
	{
		s := greetNoInline()
		fmt.Println(s)
	}
	{
		sb := []byte("message")
		sbc := bytes.Clone(sb) // inlining call to bytes.Clone
		fmt.Println(string(sbc))
	}
}
