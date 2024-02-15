package main

import (
	"cmp"
	"encoding/hex"
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

func main() {
	// range integer
	for i, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		fmt.Printf("%d: %d\n", i, v)
	}
	for v := range 10 {
		fmt.Printf("%d\n", v)
	}
	now := time.Now()
	defer time.Since(now) // go vet!

	{
		dst := make([]byte, 0, 2)
		result := hex.AppendEncode(dst, []byte("hello world"))
		// hex.Encode return panic if dst is too small
		fmt.Printf("%q\n", dst)
		fmt.Printf("%q\n", result)
	}
	x := rand.IntN(10000)
	fmt.Println(x)

	l := []string{"a", "b", "c", "d", "e"}
	r := []string{"f", "g", "h", "i", "j"}
	xx := slices.Concat(l, r)
	fmt.Println(xx)

	xx = slices.Replace(xx, 2, 5, "x", "y", "z")
	fmt.Println(xx)

	{
		fmt.Println(cmp.Or(1, 2, 3))
		fmt.Println(cmp.Or(0, 0, 3))
		x := 100
		fmt.Println(*cmp.Or(nil, nil, &x))

		var input *int = nil
		output := cmp.Or(input)
		fmt.Println(output)
	}
}
