package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args {
		fmt.Printf("%q\n", arg)
	}
}
