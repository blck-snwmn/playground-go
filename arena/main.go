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
	// see: https://github.com/golang/go/blob/release-branch.go1.20/src/arena/arena.go
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
