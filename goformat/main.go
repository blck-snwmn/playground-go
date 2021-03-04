package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d, %9.2f", 1, 2.0)
	fmt.Println()
	// インデックス指定
	// 制度指定する場合は最後のやつ参照
	fmt.Printf("%[1]d, %[1]b, %9.5[2]f", 1, 2.0)
	fmt.Println()
	fmt.Printf("%[1]*.[3]*[2]f", 3, 123456789.987654321, 5)
	fmt.Println()
	// インデックス指定をすることで、複数指定しなくていい
	fmt.Printf(template, "xaaaax")
	fmt.Println()
}

const template = `
func (%[1]s)fun1(){}
func (%[1]s)fun2(){}
func (%[1]s)fun3(){}
func (%[1]s)fun4(){}
func (%[1]s)fun5(){}
func (%[1]s)fun6(){}
func (%[1]s)fun7(){}
`
