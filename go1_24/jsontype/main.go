package main

import (
	"encoding/json"
	"fmt"
)

type emb struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type myStruct struct {
	emb
	FieldX string `json:"fieldX"`
}

func main() {
	data := `{"field1": 1,"field2":"value2","fieldX":"valueX"}`
	var s myStruct
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			// Go1.24: UnmarshalTypeError: Field: emb.field1
			// Go1.23: UnmarshalTypeError: Field: field1
			fmt.Printf("UnmarshalTypeError: Field: %s\n", e.Field)
		} else {
			fmt.Printf("Error: %s\n", err)
		}
		return
	}

	fmt.Println(s)
}
