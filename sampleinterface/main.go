package main

import (
	"fmt"

	"github.com/blck-snwmn/playground-go/sampleinterface/module"
	"github.com/blck-snwmn/playground-go/sampleinterface/user"
)

func main() {
	m := module.Module{}
	fmt.Println(user.Greet(&m))
}
