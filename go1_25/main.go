package main

import (
	"bytes"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"maps"
	"slices"
	"sync"
	"time"
)

func main() {
	{
		start := time.Now()

		var wg sync.WaitGroup
		wg.Go(func() {
			// Simulate some work
			time.Sleep(1 * time.Second)
		})

		wg.Go(func() {
			time.Sleep(2 * time.Second)
		})

		wg.Wait()

		since := time.Since(start)
		if since >= 3*time.Second {
			fmt.Println("Total execution time:", since)
		}
	}

	newMap := func(count int) map[string]string {
		m := map[string]string{}
		for i := 0; i < count; i++ {
			m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
		}
		return m
	}

	{
		m := newMap(10)

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("JSON output:", string(b))
	}
	{
		m := newMap(10)

		b, err := json.Marshal(m, json.Deterministic(true))
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("JSON output:", string(b))
	}
	{
		m := newMap(10)

		keysSeq := maps.Keys(m)
		keys := slices.Sorted(keysSeq)

		out := new(bytes.Buffer)
		enc := jsontext.NewEncoder(out)

		err := enc.WriteToken(jsontext.BeginObject)
		if err != nil {
			fmt.Println("Error writing begin object:", err)
			return
		}
		for _, k := range keys {
			err := enc.WriteToken(jsontext.String(k))
			if err != nil {
				fmt.Println("Error writing key:", err)
				return
			}
			err = enc.WriteToken(jsontext.String(m[k]))
			if err != nil {
				fmt.Println("Error writing value:", err)
				return
			}
		}
		err = enc.WriteToken(jsontext.EndObject)
		if err != nil {
			fmt.Println("Error writing end object:", err)
			return
		}

		fmt.Printf("JSON output: %s\n", out.String())
	}
}
