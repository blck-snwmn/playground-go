package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("America/Los_Angeles")
	la := location

	p := func(t time.Time) {
		fmt.Printf("%v, %v, %v\n", t, t.UTC(), t.Unix())
	}

	{
		fmt.Println("=====start summer time=====")
		d := time.Date(2023, 3, 12, 1, 59, 59, 0, la)
		p(d) // 2023-03-12 01:59:59 -0500 EST

		d = d.Add(time.Second)
		p(d) // 2023-03-12 03:00:00 -0400 EDT
	}
	{
		fmt.Println("=====end summer time=====")
		d := time.Date(2023, 11, 5, 1, 59, 59, 0, la)
		p(d) // 2023-11-05 01:59:59 -0400 EDT

		d = d.Add(time.Second)
		p(d) // 2023-11-05 01:00:00 -0500 EST
	}
}
