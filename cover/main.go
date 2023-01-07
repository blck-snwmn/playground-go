package main

import (
	"fmt"
	"os"

	"github.com/blck-snwmn/playground-go/cover/order"
	"github.com/blck-snwmn/playground-go/cover/pay"
)

func main() {
	switch len(os.Args) {
	case 0, 1:
		a := pay.Pay(10)
		fmt.Println(a)
	case 2:
		mode := os.Args[1]
		o := order.Order(mode)
		fmt.Println(o)
	default:
		fmt.Println("end")
	}
}
