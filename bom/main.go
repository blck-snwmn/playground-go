package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spkg/bom"
)

func main() {
	for _, path := range []string{"./bom.csv", "./non-bom.csv"} {
		fmt.Printf("----\nread file:%s\n", path)

		file, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("output:%q\n", string(file))

		reader := csv.NewReader(bom.NewReader(bytes.NewReader(file)))
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, r := range records {
			fmt.Println(r)
		}
	}
}
