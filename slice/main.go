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
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		oldlen := len(s)
		fmt.Println(s)
		s = append(s[:2], s[4:]...)
		fmt.Println(s)
		fmt.Println(s[len(s):oldlen]) // show deleted elements
	}
}
