package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// strings
	for s := range strings.Lines("abc\ndef\nghi\n") {
		fmt.Print(s)
	}
	for s := range strings.SplitSeq("abc,def,ghi", ",") {
		fmt.Println(s)
	}
	for s := range strings.SplitAfterSeq("abc,def,ghi", ",") {
		fmt.Println(s)
	}
	for i, s := range strings.Fields("abc\ndef\tghi jkl") {
		fmt.Printf("%d: %s\n", i, s)
	}
	for i, s := range strings.FieldsFunc("ddddaxxxxxbyyyy", func(r rune) bool {
		return r == 'a' || r == 'b'
	}) {
		fmt.Printf("%d: %s\n", i, s)
	}
	// bytes
	for bs := range bytes.Lines([]byte("abc\ndef\nghi\n")) {
		fmt.Print(string(bs))
	}
	for bs := range bytes.SplitSeq([]byte("abc,def,ghi"), []byte(",")) {
		fmt.Println(string(bs))
	}
	for bs := range bytes.SplitAfterSeq([]byte("abc,def,ghi"), []byte(",")) {
		fmt.Println(string(bs))
	}
	for i, bs := range bytes.Fields([]byte("abc\ndef\tghi jkl")) {
		fmt.Printf("%d: %s\n", i, string(bs))
	}
	for i, bs := range bytes.FieldsFunc([]byte("ddddaxxxxxbyyyy"), func(r rune) bool {
		return r == 'a' || r == 'b'
	}) {
		fmt.Printf("%d: %s\n", i, string(bs))
	}
	// Root
	r, err := os.OpenRoot(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()
	f, err := r.Open("exp/main_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("```go\n%s```\n", string(bs))

	_, err = r.Open("../README.md")
	fmt.Println(err)

	//
	now := time.Now()
	fmt.Println(now)
	txt, err := now.AppendText([]byte("time: "))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txt)

	txt, err = now.AppendBinary([]byte("time: "))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txt)
}
