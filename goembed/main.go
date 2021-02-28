package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed test.json
var test []byte

//go:embed test2.json
var test2 []byte

//go:embed test.json
//go:embed data/*
var data embed.FS

func main() {
	fmt.Println(string(test))
	fmt.Println(string(test2))
	o, _ := data.ReadFile("data/data1.json")
	fmt.Println(string(o))
	o, _ = data.ReadFile("test.json")
	fmt.Println(string(o))
	_, err := data.ReadFile("testx.json")
	fmt.Println(err)
}
