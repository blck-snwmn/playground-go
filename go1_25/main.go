package main

import (
	"bytes"
	"context"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"os"
	"slices"
	"strings"
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
	sampleJsonV2Read()
	sampleJsonV2Transform()

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"Sample log message",
		slog.GroupAttrs("req", slog.String("method", "GET"),
			slog.String("url", "https://example.com"),
			slog.Int("status", 200),
		),
	)

	// panicSample causes the panic
	panicSample()
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

func sampleJsonV2Read() {
	type item struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	{
		fmt.Println("\nReading JSON with jsontext decoder:")

		data := `[{"field1":"value0","field2":0},{"field1":"value1","field2":10}, {"field1":"value2","field2":20}, {"field1":"value3","field2":30}, {"field1":"value4","field2":40}]`
		fmt.Printf("Reading JSON: %s\n", data)
		decoder := jsontext.NewDecoder(strings.NewReader(data))
		t, err := decoder.ReadToken()
		if err != nil {
			fmt.Println("Error reading token:", err)
			return
		}
		if t.Kind() != '[' {
			fmt.Println("Expected BeginArray token, got:", t)
			return
		}
		for decoder.PeekKind() != 0 {
			if decoder.PeekKind() == ']' {
				break
			}

			var itm item
			err := json.UnmarshalDecode(decoder, &itm)
			if err != nil {
				fmt.Println("Error unmarshaling item:", err)
				return
			}
			fmt.Printf("Item: %+v\n", itm)
		}
	}
	{
		fmt.Println("\nReading JSON with jsontext decoder from multiple lines:")

		data := `{"field1": "value1", "field2": 10}
		{"field1": "value2", "field2": 20}
		{"field1": "value3", "field2": 30}
		{"field1": "value4", "field2": 40}`

		fmt.Printf("Reading JSON: %s\n", data)

		decoder := jsontext.NewDecoder(strings.NewReader(data))

		for {
			var itm item
			err := json.UnmarshalDecode(decoder, &itm)
			if err != nil {
				if err == io.EOF {
					fmt.Println("End of input reached")
					break
				}

				fmt.Println("Error unmarshaling item:", err)
			}
			fmt.Printf("Item: %+v\n", itm)
		}
	}
	{
		fmt.Println("\nReading JSON with jsontext decoder each token:")
		data := `{"field1": "value1", "field2": 10}`
		fmt.Printf("Reading JSON: %s\n", data)
		decoder := jsontext.NewDecoder(strings.NewReader(data))
		for {
			token, err := decoder.ReadToken()
			if err != nil {
				if err == io.EOF {
					fmt.Println("End of input reached")
					break
				}
				fmt.Println("Error reading token:", err)
				break
			}
			fmt.Printf("\tKind:%-6s Token:%+v\n", token.Kind(), token)
		}
	}
}

func sampleJsonV2Transform() {
	type user struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address struct {
			City string `json:"city"`
			Zip  string `json:"zip"`
		} `json:"address"`
	}
	data := `[
	{"id": 1, "name": "alice", "age": 30, "address": {"city": "Wonderland", "zip": "12345"}},
	{"id": 2, "name": "bob", "age": 25, "address": {"city": "Builderland", "zip": "67890"}},
	{"id": 3, "name": "charlie", "age": 35, "address": {"city": "Chocolate Factory", "zip": "54321"}}
	]`

	decoder := jsontext.NewDecoder(strings.NewReader(data))

	out := new(bytes.Buffer)
	enc := jsontext.NewEncoder(out)

	for {
		fmt.Println("Reading next token...")
		pt := decoder.PeekKind()
		switch pt {
		case 0:
			// EndOfInput
			fmt.Println("End of input reached")
			break
		case '[':
			// BeginArray
			err := enc.WriteToken(jsontext.BeginArray)
			if err != nil {
				fmt.Println("Error writing begin array:", err)
				break
			}
			fmt.Println("BeginArray token found, starting to read array elements...")
			decoder.ReadToken() // Consume the BeginArray token
			continue
		case ']':
			// EndArray
			err := enc.WriteToken(jsontext.EndArray)
			if err != nil {
				fmt.Println("Error writing end array:", err)
				break
			}
			fmt.Println("EndArray token found, finishing reading array elements...")
			decoder.ReadToken() // Consume the EndArray token
			break
		}
		var u user
		err := json.UnmarshalDecode(decoder, &u)
		if err != nil {
			if err == io.EOF {
				fmt.Println("End of input reached")
				break
			}
			fmt.Println("Error unmarshaling user:", err)
			break
		}
		fmt.Printf("User: %+v\n", u)

		maskedUser := user{
			ID:   u.ID,
			Name: u.Name,
			Age:  u.Age,
			Address: struct {
				City string `json:"city"`
				Zip  string `json:"zip"`
			}{
				City: "REDACTED",
				Zip:  "REDACTED",
			},
		}
		err = json.MarshalEncode(enc, maskedUser)
		if err != nil {
			fmt.Println("Error marshaling masked user:", err)
			break
		}
	}
	fmt.Println("Masked JSON output:", out.String())
}

func panicSample() {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	panic("This is a sample panic")
}
