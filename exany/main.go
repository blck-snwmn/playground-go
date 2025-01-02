package main

import (
	"fmt"
	"reflect"
)

type foo struct{}

func main() {
	f := foo{}
	do1(f)
	do1(&f)
	do1p(f)
	do1p(&f)

	// do1: main.foo
	// do2: main.foo
	// do1: *main.foo
	// do2: *main.foo
	// do1p: main.foo
	// do2: *interface {}
	// do1p: *main.foo
	// do2: *interface {}

	fmt.Println("----")
	showReflectW(f)
	// Type: main.foo
	// type: main.foo
	// Final Type: main.foo

	fmt.Println("----")
	showReflectW(&f)
	// Type: *main.foo
	// type: *main.foo in pointer
	// type: main.foo
	// Final Type: main.foo

	fmt.Println("----")
	showReflectWP(f)
	// Type: *interface {}
	// type: *interface {} in pointer
	// type: interface {}
	// type: interface {} in interface1
	// type: interface {} in interface2
	// Final Type: main.foo
}

func do1(a any) {
	fmt.Printf("do1: %T\n", a)
	do2(a)
}

func do1p(a any) {
	fmt.Printf("do1p: %T\n", a)
	do2(&a)
}

func do2(a any) {
	fmt.Printf("do2: %T\n", a)
}

func showReflectW(a any) {
	showReflect(a)
}

func showReflectWP(a any) {
	showReflect(&a)
}
func showReflect(a any) {
	typ := reflect.TypeOf(a)
	fmt.Printf("Type: %v\n", typ)
	// if type is pointer, extract type
	if typ.Kind() == reflect.Ptr {
		fmt.Printf("type: %v in pointer\n", typ)
		typ = typ.Elem()
	}
	fmt.Printf("type: %v\n", typ)
	// if type is interface, extract type
	if typ.Kind() == reflect.Interface {
		fmt.Printf("type: %v in interface1\n", typ)
		val := reflect.ValueOf(a).Elem()
		if val.Kind() == reflect.Interface {
			fmt.Printf("type: %v in interface2\n", typ)
			typ = val.Elem().Type()
		}
	}
	fmt.Printf("Final Type: %v\n", typ)
}
