package main

import (
	"encoding/json"
	"testing"
	"testing/synctest"
	"time"
)

func TestSynctest(t *testing.T) {
	w := t.Output()
	enc := json.NewEncoder(w)
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
