package test

import (
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Run("sub1", func(t *testing.T) {

		t.Setenv("key1", "value1")
		v := os.Getenv("key1")
		if v != "value1" {
			t.Errorf("got %s, want value1", v)
		}
	})
	t.Run("sub2", func(t *testing.T) {
		t.Setenv("key1", "value2")
		v := os.Getenv("key1")
		if v != "value2" {
			t.Errorf("got %s, want value2", v)
		}
	})
}
