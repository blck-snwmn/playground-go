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
	sampleJsonV2Deterministic()
}

func sampleJsonV2Deterministic() {
	fmt.Println("Sample JSON v2 deterministic output:")
	{
		fmt.Println("Deterministic JSON output with string keys:")
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		b, err := json.Marshal(m, json.Deterministic(true))
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("\tJSON output:", string(b))
	}
	{
		fmt.Println("Deterministic JSON output with integer keys:")
		m := map[int]string{
			1: "value1",
			2: "value2",
			3: "value3",
		}
		b, err := json.Marshal(m, json.Deterministic(true))
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("\tJSON output:", string(b))
	}
	{
		fmt.Println("Deterministic JSON output with struct keys:")

		m := map[x]string{
			{"s1", 1}: "value1",
			{"s2", 2}: "value2",
			{"s3", 3}: "value3",
		}
		b, err := json.Marshal(m, json.Deterministic(true))
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("\tJSON output:", string(b))
	}
}

type x struct {
	S string `json:"s"`
	I int    `json:"i"`
}

// MarshalJSON implements json.Marshaler interface.
// This is required when using struct as a map key because:
// 1. JSON spec requires object keys to be strings
// 2. Without this, marshaling struct produces JSON object like {"s":"s1","i":1}
// 3. This would result in invalid JSON: {{"s":"s1","i":1}: "value1"}
// 4. By implementing MarshalJSON, we convert struct to a string representation
func (x x) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s-%d"`, x.S, x.I)
	return []byte(s), nil
}
