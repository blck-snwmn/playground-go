package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a int
	b uint
	c int32
}

type B struct {
	b uint
	c int32
}

func (*B) do() {}

type C struct {
	f func()
}

func (c *C) do() {
	c.f()
}

type I interface {
	do()
}

func main() {
	a := A{10, 20, 30}
	fmt.Printf("[A]struct adress is %p\n", &a)
	fmt.Printf("[A]field a adress = %p, value = %d\n", &a.a, a.a)
	fmt.Printf("[A]field b adress = %p, value = %d\n", &a.b, a.b)
	fmt.Printf("[A]field c adress = %p, value = %d\n", &a.c, a.c)

	// 別のstructのアドレスを別のstructへ代入する
	var b *B
	b = (*B)(unsafe.Pointer(&a.b)) // A.b 以降のstructの構造が同じなので、A.bのアドレスをBのポインタ変数は渡してみる
	fmt.Printf("[B]struct adress is %p\n", &b)
	fmt.Printf("[B]field b adress = %p, value = %d\n", &b.b, b.b)
	fmt.Printf("[B]field c adress = %p, value = %d\n", &b.c, b.c)

	// size
	fmt.Printf("[B]field a size is %d\n", uintptr(unsafe.Pointer(&a.b))-uintptr(unsafe.Pointer(&a.a)))
	fmt.Printf("[B]field b size is %d\n", uintptr(unsafe.Pointer(&a.c))-uintptr(unsafe.Pointer(&a.b)))

	var i I = b
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", i.do)

	var f = i.do
	fmt.Printf("%p\n", f)

	fmt.Println("====")
	var c = &C{func() {}}
	var ii I = c
	fmt.Printf("c.f=%p\n", c.f)
	fmt.Printf("c.do=%p\n", c.do)
	fmt.Printf("ii.do%p\n", ii.do)

	c.f = func() {
		fmt.Println("a")
	}
	fmt.Printf("c.f==\t%p\n", c.f)
	fmt.Printf("c.do==\t%p\n", c.do)
	fmt.Printf("ii.do=\t%p\n", ii.do)
	fmt.Printf("ii.do(*C)=\t%p\n", ii.(*C).do)

	fmt.Printf("c=\t%p\n", c)
	fmt.Printf("ii=\t%p\n", ii)
	fmt.Printf("ii.*C=\t%p\n", ii.(*C))
}
