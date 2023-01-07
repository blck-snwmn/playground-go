//go:build goexperiment.arenas
// +build goexperiment.arenas

package main

import (
	"arena"
	"fmt"
)

type Elm[T any] struct {
	id    string
	value T
}

func main() {
	a := arena.NewArena()
	defer a.Free()
	{
		e := arena.New[Elm[int]](a)
		*e = Elm[int]{
			id:    "id-123456",
			value: 100,
		}
		fmt.Println(e)
	}

	fmt.Println("end")
}
