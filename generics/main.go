package main

import "fmt"

type curve[Point point[Point]] struct {
	newPoint func() Point
}

type point[T any] interface {
	Do(T, []byte) (T, error)
}

func NewFoo() *Foo {
	return &Foo{id: "1"}
}

type Foo struct {
	id string
}

func (f *Foo) Do(ff *Foo, b []byte) (*Foo, error) {
	return &Foo{id: ff.id + string(b)}, nil
}

var fooCurve = curve[*Foo]{
	newPoint: NewFoo,
}

func main() {
	f := fooCurve.newPoint()

	print := func(f *Foo) {
		fmt.Printf("%v\n", f)
	}

	f, _ = f.Do(f, []byte("aaaaa"))
	print(f)
	f, _ = f.Do(f, []byte("bbbbb"))
	print(f)
	f, _ = f.Do(f, []byte("ccccc"))
	print(f)
}
