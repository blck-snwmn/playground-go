package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("data: %X\n", data)
	data2 := data
	fmt.Printf("data2: %X\n", data2)

	fmt.Println("=====")

	data2[0] = 0x14
	fmt.Printf("data: %X\n", data)
	fmt.Printf("data2: %X\n", data2)

	fmt.Println("=====")

	data3 := data[2:6]
	fmt.Printf("data3: %X\n", data3)
	data3[1] = 0x15
	fmt.Printf("data: %X\n", data)
	fmt.Printf("data2: %X\n", data2)
	fmt.Printf("data3: %X\n", data3)
	{
		fmt.Println("=====")
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		oldlen := len(s)
		fmt.Printf("%[1]p: %[1]v\n", s)
		fmt.Printf("cap=%d, %d\n", cap(s), cap(s[:2]))
		s = append(s[:2], s[4:]...)
		fmt.Printf("%[1]p: %[1]v\n", s)
		fmt.Printf("%[1]p: %[1]v\n", s[len(s):oldlen]) // show deleted elements
		fmt.Printf("cap=%d, %d, %d, %d\n", cap(s), cap(s[:2]), cap(s[4:]), cap(s[len(s):oldlen]))
	}
	{
		fmt.Println("=====")
		var x []greeter = []greeter{
			english{},
			japanese{},
		}

		for _, v := range x {
			v.greet()
		}
	}
}

type greeter interface {
	greet()
}

type english struct{}

func (e english) greet() {
	fmt.Println("Good morning!")
}

type japanese struct{}

func (j japanese) greet() {
	fmt.Println("おはようございます！")
}
