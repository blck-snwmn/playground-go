package main

import (
	"fmt"
	"time"
)

// Add Duration.Abs method
func sampleTime() {
	l := time.Date(2022, 12, 1, 12, 13, 10, 0, time.Local)
	r := time.Date(2022, 12, 4, 12, 13, 10, 0, time.Local)
	if l.Sub(r).Abs() == r.Sub(l).Abs() {
		fmt.Println("|l-r|==|r-l|")
	}
}

func main() {
	sampleTime()
}
