package main

import (
	"fmt"
	"reflect"
)

func assignableValue(src interface{}, dest interface{}) bool {
	// ポインタ型の場合、Elem() がsrcの型と同じ場合、AssignableTo が true になる
	val := reflect.ValueOf(dest)
	fmt.Printf("\t\t[assignableValue] val := reflect.ValueOf(dest): `%v`\n", val)
	typ := val.Type()
	fmt.Printf("\t\t[assignableValue] typ := val.Type(): `%v`\n", typ)
	destType := typ.Elem()
	fmt.Printf("\t\t[assignableValue] destType := typ.Elem(): `%v`\n", destType)

	return reflect.TypeOf(src).AssignableTo(destType)
}

func assignableType(src interface{}, dest interface{}) bool {
	return reflect.TypeOf(src).AssignableTo(reflect.TypeOf(dest))
}

func printAssignable(src interface{}, dest interface{}, assignable func(src interface{}, dest interface{}) bool) {
	ownType := reflect.TypeOf(src)
	targetType := reflect.TypeOf(dest)

	fmt.Printf("\t`%v` assign to `%v` =  %t\n", ownType, targetType, assignable(src, dest))
}

func printAssignableType(src interface{}, dest interface{}) {
	fmt.Println("printAssignableType")
	printAssignable(src, dest, assignableType)
}

func printAssignableValue(src interface{}, dest interface{}) {
	fmt.Println("printAssignableValue")
	printAssignable(src, dest, assignableValue)
}

type greeter interface {
	greet() string
}

var _ greeter = helloGreeter{}

type helloGreeter struct{}

func (helloGreeter) greet() string {
	return "hello world"
}

var _ greeter = (*goodByeGreeter)(nil)

type goodByeGreeter struct{}

func (*goodByeGreeter) greet() string {
	return "good bye"
}

func assignableExperiment() {
	{
		type A interface{}
		type AA struct{}
		var a A
		printAssignableValue(AA{}, &a)
		printAssignableType(AA{}, AA{})
	}
	{
		var g greeter
		hg := helloGreeter{}
		gg := &goodByeGreeter{}
		printAssignableValue(hg, &g)
		printAssignableValue(gg, &g)
	}
}

func main() {
	assignableExperiment()
}
