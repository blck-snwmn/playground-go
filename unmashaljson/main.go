package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string `json:"name"`
}

type Nop struct{}

func main() {
	{
		content := `{"name": "John"}`
		var data Data
		err := json.Unmarshal([]byte(content), &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
	{
		content := `{"name": "John"}`
		var nop Nop
		err := json.Unmarshal([]byte(content), &nop)
		if err != nil {
			panic(err)
		}
		fmt.Println(nop)
	}
	{
		// occurs panic
		// content := ""
		// var data Data
		// err := json.Unmarshal([]byte(content), &data)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(data)
	}
	{
		content := ""
		var nop Nop
		err := json.Unmarshal([]byte(content), &nop)
		if err != nil {
			panic(err)
		}
		fmt.Println(nop)
	}
}
