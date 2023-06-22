package testdir

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_tmpdir(t *testing.T) {
	dir := t.TempDir()
	wf, err := os.Create(filepath.Join(dir, "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	wf.WriteString("test")
	wf.Close()

	rf, err := os.Open(filepath.Join(dir, "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 4)
	rf.Read(buf)
	rf.Close()
	if string(buf) != "test" {
		t.Fatal("invalid data")
	}
}
