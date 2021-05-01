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

// func do() {
// 	// q := strings.Repeat("xあ", 3)
// 	// i := &Inner{
// 	// 	InnerNumber: 10,
// 	// }
// 	s := &SearchRequest{
// 		// Query:      q,
// 		// PageNumber: 3,
// 		ResultPerPage: -5234,
// 		// PackedSample:  []int64{1, 2, 1000, 4, 5},
// 		// Test: 1,
// 		// Test32: 2,
// 		// Type:       SearchRequest_BBBB,
// 		// InnerValue: i,
// 	}
// 	r, err := proto.Marshal(s)
// 	fmt.Println(err)
// 	show(r)
// 	// [16 1 34 1 120 192 62 2]
// 	// 192=01000000
// 	// 62=0111110
// 	// 1000=1111101000
// 	// 011111001000000

// 	// 1
// 	// 000001111101000
// }

// func doDuplicateField() {
// 	var buf []byte

// 	{
// 		i := &Inner{
// 			InnerNumber: 10,
// 			Tttttt:      "x",
// 		}
// 		s := &SearchRequest{
// 			PackedSample: []int64{1, 2, 1000, 4, 5},
// 			InnerValue:   i,
// 			PageNumber:   1,
// 			Test:         1,
// 		}
// 		r, err := proto.Marshal(s)
// 		if err != nil {
// 			return
// 		}
// 		buf = append(buf, r...)
// 	}
// 	{
// 		i := &Inner{
// 			InnerStr: "testsssss",
// 			Tttttt:   "z",
// 		}
// 		s := &SearchRequest{
// 			PackedSample: []int64{3, 5, 6},
// 			InnerValue:   i,
// 			Query:        "a",
// 			Test:         2,
// 		}
// 		r, err := proto.Marshal(s)
// 		if err != nil {
// 			return
// 		}
// 		buf = append(buf, r...)
// 	}
// 	// SearchRequestについて
// 	// - 片方だけ定義されているものはマージされている
// 	// - 両方で定義されているfloat のフィールドは後勝ち
// 	// - Repeat (PackedSample) は連結されている
// 	// Innerも同様に片方だけならマージされ、両方で定義されているものは後勝ち
// 	s := &SearchRequest{}
// 	err := proto.Unmarshal(buf, s)
// 	fmt.Println(err)
// 	fmt.Printf("PackedSample={%+v}\n", s.PackedSample)
// 	fmt.Printf("InnerValue={%+v}\n", s.InnerValue)
// 	fmt.Printf("PageNumber={%+v}\n", s.PageNumber)
// 	fmt.Printf("Query={%+v}\n", s.Query)
// 	fmt.Printf("Test={%+v}\n", s.Test)
// }

func main() {
	// do()
	// doDuplicateField()
	s := &SearchRequest{
		UInt32:    1,
		UInt64:    100,
		SInt32:    122,
		SInt64:    56789,
		IInt32:    6423,
		IInt64:    234242,
		BBool:     true,
		FFixed64:  443,
		SSfixed64: 677,
		DDouble:   1.25,
	}
	b, err := proto.Marshal(s)
	if err != nil {
		return
	}
	show(b)
	// bool の値チェック。0じゃなければtrue
	// err = proto.Unmarshal([]byte{0b01101000, 0b00000011}, s)
	// if err != nil {
	// 	fmt.Println("error", err)
	// }
	// fmt.Println(s)
}
