package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed test.json
var test []byte

//go:embed test2.json
var test2 []byte

//go:embed test.json
//go:embed data/*
var data embed.FS

//go:embed template/template.tmpl
var tmp string

type Foo struct {
	Name string
}

func main() {
	fmt.Println(string(test))
	fmt.Println(string(test2))
	o, _ := data.ReadFile("data/data1.json")
	fmt.Println(string(o))
	o, _ = data.ReadFile("test.json")
	fmt.Println(string(o))
	// read no exsit file
	_, err := data.ReadFile("testx.json")
	fmt.Println(err)

	// using text/template
	t, _ := template.New("foo").Parse(tmp)
	t.Execute(os.Stdout, Foo{Name: "template"})
}
