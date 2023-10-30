package main

import (
	_ "embed"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v3"
)

//go:embed valid.yml
var valid []byte

type Sample struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Rules       string `yaml:"rules"`
}

func main() {
	var s Sample
	err := yaml.Unmarshal(valid, &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", s)
	fmt.Println(reflect.TypeOf(s.Rules))
}
