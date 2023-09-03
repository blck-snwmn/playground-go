package testbuildtag

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	if os.Getenv("CLASS") == "" {
		t.Skip("Skipping no CLASS environment variable set")
	}
}
