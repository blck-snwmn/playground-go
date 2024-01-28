package foo_test

import (
	"testing"

	"github.com/blck-snwmn/playground-go/exporttest/exporttest/foo"
)

func TestExport(t *testing.T) {
	if foo.Export() != "internal" {
		t.Fatal("wrong value")
	}
}
