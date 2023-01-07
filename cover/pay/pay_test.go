package pay

import "testing"

func TestPay(t *testing.T) {
	if got := Pay(10); got != 0 {
		t.Errorf("unexpected value. got=%d, want=10", got)
	}
}
