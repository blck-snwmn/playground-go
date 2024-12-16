package exp

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestExp(t *testing.T) {
	// This test is set to wait for one day using sleep, but the test execution time is less than one second.
	synctest.Run(func() {
		start := time.Now()
		time.Sleep(24 * time.Hour)

		since := time.Since(start)
		if since != 24*time.Hour {
			t.Errorf("time.Since(start) = %v, want >= 1s", since)
		}
	})
}
