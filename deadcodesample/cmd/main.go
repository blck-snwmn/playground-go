package main

import (
	"fmt"

	"github.com/blck-snwmn/playground-go/deadcodesample/module"
)

func main() {
	str := module.ExportedFunc3()
	fmt.Println(str)
}
