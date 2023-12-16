package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("America/Los_Angeles")
	la := location

	locationTokyo, _ := time.LoadLocation("Asia/Tokyo")
	jst := locationTokyo

	p := func(t time.Time) {
		fmt.Printf("%v, %v, %v\n", t, t.UTC(), t.Unix())
	}

	{
		fmt.Println("=====start summer time=====")
		d := time.Date(2023, 3, 12, 1, 59, 59, 0, la)
		p(d) // 2023-03-12 01:59:59 -0800 PST

		d = d.Add(time.Second)
		p(d) // 2023-03-12 03:00:00 -0700 PDT
	}
	{
		fmt.Println("=====end summer time=====")
		d := time.Date(2023, 11, 5, 1, 59, 59, 0, la)
		p(d) // 2023-11-05 01:59:59 -0700 PDT

		d = d.Add(time.Second)
		p(d) // 2023-11-05 01:00:00 -0800 PST
	}
	{
		fmt.Println("==========")
		var t time.Time
		t = time.Date(2023, 9, 1, 8, 0, 0, 0, la)
		fmt.Println(t)
		fmt.Println(t.In(jst))

		fmt.Println("==========")
		t = time.Date(2023, 12, 1, 8, 0, 0, 0, la)
		fmt.Println(t)
		fmt.Println(t.In(jst))
	}
}
