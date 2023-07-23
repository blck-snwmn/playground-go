package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
	"testing/slogtest"
)

func Test_isTest(t *testing.T) {
	if got := isTest(); got != true {
		t.Errorf("isTest() = %v, want %v", got, true)
	}
}

func Test_success_slog_w_json(t *testing.T) {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}
			ms = append(ms, m)
		}
		return ms
	}

	err := slogtest.TestHandler(h, results)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_fail_slog_w_txt(t *testing.T) {
	var buf bytes.Buffer
	h := slog.NewTextHandler(&buf, nil) // Use texthandler as a failing test.

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}
			ms = append(ms, m)
		}
		return ms
	}

	err := slogtest.TestHandler(h, results)
	if err != nil {
		t.Fatal(err)
	}
}
