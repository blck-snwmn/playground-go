package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// strings
	fmt.Println(strings.TrimPrefix("aaabbbcccc", "aa"))
	fmt.Println(strings.CutPrefix("aaabbbcccc", "aa"))
	fmt.Println(strings.TrimSuffix("aaabbbcccc", "cc"))
	fmt.Println(strings.CutSuffix("aaabbbcccc", "cc"))
	// time
	n := time.Now()
	fmt.Println(n.Format(time.DateOnly))
	fmt.Println(n.Format(time.DateTime))
	fmt.Println(n.Format(time.TimeOnly))

	fmt.Println(time.Now().Compare(n)) // 1
}
