package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func show(r []byte) {
	for _, rr := range r {
		fmt.Printf("0b%08b, ", rr)
	}
	fmt.Println()
}

func do() {
	// q := strings.Repeat("xあ", 3)
	// i := &Inner{
	// 	InnerNumber: 10,
	// }
	s := &SearchRequest{
		// Query:      q,
		// PageNumber: 3,
		ResultPerPage: -5234,
		// PackedSample:  []int64{1, 2, 1000, 4, 5},
		// Test: 1,
		// Test32: 2,
		// Type:       SearchRequest_BBBB,
		// InnerValue: i,
	}
	r, err := proto.Marshal(s)
	fmt.Println(err)
	show(r)
	// [16 1 34 1 120 192 62 2]
	// 192=01000000
	// 62=0111110
	// 1000=1111101000
	// 011111001000000

	// 1
	// 000001111101000
}

func doDuplicateField() {
	var buf []byte

	{
		i := &Inner{
			InnerNumber: 10,
			Tttttt:      "x",
		}
		s := &SearchRequest{
			PackedSample: []int64{1, 2, 1000, 4, 5},
			InnerValue:   i,
			PageNumber:   1,
			Test:         1,
		}
		r, err := proto.Marshal(s)
		if err != nil {
			return
		}
		buf = append(buf, r...)
	}
	{
		i := &Inner{
			InnerStr: "testsssss",
			Tttttt:   "z",
		}
		s := &SearchRequest{
			PackedSample: []int64{3, 5, 6},
			InnerValue:   i,
			Query:        "a",
			Test:         2,
		}
		r, err := proto.Marshal(s)
		if err != nil {
			return
		}
		buf = append(buf, r...)
	}
	// SearchRequestについて
	// - 片方だけ定義されているものはマージされている
	// - 両方で定義されているfloat のフィールドは後勝ち
	// - Repeat (PackedSample) は連結されている
	// Innerも同様に片方だけならマージされ、両方で定義されているものは後勝ち
	s := &SearchRequest{}
	err := proto.Unmarshal(buf, s)
	fmt.Println(err)
	fmt.Printf("PackedSample={%+v}\n", s.PackedSample)
	fmt.Printf("InnerValue={%+v}\n", s.InnerValue)
	fmt.Printf("PageNumber={%+v}\n", s.PageNumber)
	fmt.Printf("Query={%+v}\n", s.Query)
	fmt.Printf("Test={%+v}\n", s.Test)
}

func main() {
	// do()
	// doDuplicateField()
	s := &SearchRequest{
		PackedSample: []int64{
			100000000,
			// 2,
			// 3,
		},
		// PackedSampleSigned: []int64{
		// 	1,
		// 	2,
		// 	3,
		// },
		// PackedSampleUnsigned: []uint32{
		// 	1,
		// 	2,
		// 	3,
		// },
		// Query: "あ",
		// PageNumber: -10,
	}
	b, err := proto.Marshal(s)
	if err != nil {
		return
	}
	show(b)
	// 0b00011010, 0b00000011, 0b00000001, 0b00000010, 0b00000011,
	// 0b00100010, 0b00000011, 0b01100001, 0b01100010, 0b01100011,
	// 0b01001010, 0b00000011, 0b00000010, 0b00000100, 0b00000110,
	// 0b01010010, 0b00000011, 0b00000001, 0b00000010, 0b00000011,

	//  , , ,
	// 00101111101011110000100000000
}
