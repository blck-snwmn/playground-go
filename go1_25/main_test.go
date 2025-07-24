package main

import (
	jsonv1 "encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestSynctest(t *testing.T) {
	w := t.Output()
	enc := jsonv1.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(map[string]string{
		"test": "synctest",
	})

	t.Attr("test", "synctest")

	// This test is set to wait for one day using sleep, but the test execution time is less than one second.
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		time.Sleep(24 * time.Hour)

		since := time.Since(start)
		if since != 24*time.Hour {
			t.Errorf("time.Since(start) = %v, want >= 1s", since)
		}
	})
}

func TestJSONV2Deterministic(t *testing.T) {

	m := map[string]string{}
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("key%02d", i)] = fmt.Sprintf("value%02d", i)
	}

	var wg sync.WaitGroup
	jsonbs := make([]string, 100)
	for i := 0; i < 100; i++ {
		wg.Go(func() {
			b, err := jsonv2.Marshal(m, jsonv2.Deterministic(true))
			if err != nil {
				t.Error("Error marshaling JSON:", err)
				return
			}
			jsonbs[i] = string(b)
		})
	}
	wg.Wait()

	for _, b := range jsonbs[1:] {
		if b != jsonbs[0] {
			t.Error("JSON output is not deterministic")
			return
		}
	}
	t.Logf("JSON output is deterministic: %s", jsonbs[0])
}
