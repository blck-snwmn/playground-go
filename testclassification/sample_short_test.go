package testbuildtag

import "testing"

func TestShort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping short test")
	}
}
