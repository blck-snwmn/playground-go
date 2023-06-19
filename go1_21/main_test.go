package main

import (
	"testing"
)

func Test_isTest(t *testing.T) {
	if got := isTest(); got != true {
		t.Errorf("isTest() = %v, want %v", got, true)
	}
}
