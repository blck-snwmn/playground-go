package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	// q := strings.Repeat("x", 2)
	i := &Inner{
		InnerNumber: 10,
	}
	s := &SearchRequest{
		// Query:         q,
		// PageNumber:    3,
		// ResultPerPage: 5234,
		// PackedSample:  []int64{1, 2, 1000, 4, 5},
		// Test:          1,
		Type:       SearchRequest_BBBB,
		InnerValue: i,
	}
	r, err := proto.Marshal(s)
	fmt.Println(err)
	fmt.Printf("%d\n", r)
	// [16 1 34 1 120 192 62 2]
	// 192=01000000
	// 62=0111110
	// 1000=1111101000
	// 011111001000000

	// 1
	// 000001111101000
}
