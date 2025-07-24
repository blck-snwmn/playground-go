package main

import (
	"encoding/json/v2"
	"fmt"
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

	{
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
		}

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("JSON output:", string(b))
	}
	{
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
		}

		// NOTE: Keys are not alphabetized even when using jsonv2.Deterministic
		b, err := json.Marshal(m, json.Deterministic(true))
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println("JSON output:", string(b))
	}
}
