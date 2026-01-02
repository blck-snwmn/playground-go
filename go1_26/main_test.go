package main

import (
	"fmt"
	"os"
	"testing"
)

func TestXxx1(t *testing.T) {
	d := t.ArtifactDir()
	fmt.Println("ArtifactDir:", d)

	d = t.ArtifactDir()
	fmt.Println("ArtifactDir:", d)
}

func TestXxx2(t *testing.T) {
	d := t.ArtifactDir()
	fmt.Println("ArtifactDir:", d)

	f, err := os.Create(d + "/tmp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "hello, world")
}
