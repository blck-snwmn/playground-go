package main

import "fmt"

type AA[T any] interface {
	New() T
}

type BB[Aa AA[Aa]] struct {
	elm Aa
}

func (b BB[Aa]) Dd() Aa {
	return b.elm.New()
}

type C struct{}

func (c C) New() C {
	return C{}
}

func main() {
	bc := BB[C]{}
	fmt.Printf("%#v\n", bc.Dd().New())
}
