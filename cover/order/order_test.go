package order

import "testing"

func TestOrder(t *testing.T) {
	if got := Order("test"); got != "world" {
		t.Errorf("got=%v, want=%v", got, "world")
	}
}
