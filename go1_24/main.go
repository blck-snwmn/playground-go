package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

// generics type alias
type F[T any] = func(t T, s string)

var ft F[string] = func(t string, s string) {
	fmt.Println(t, s)
}

func do[T any](f F[T], t T, s string) {
	f(t, s)
}

type innerWithOmitempty struct {
	A int    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

type inner struct {
	A int    `json:"a"`
	B string `json:"b"`
}

type outer struct {
	IWOEmpty  innerWithOmitempty `json:"iwoempty,omitempty"`
	IWOZero   innerWithOmitempty `json:"iwozero,omitzero"`
	IEmpty    inner              `json:"iempty,omitempty"`
	IZero     inner              `json:"izero,omitzero"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdateAt  time.Time          `json:"updated_at,omitzero"`
}

func main() {
	ft("hello", "world")
	do(ft, "xxxx", "yyyyy")

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

	sl := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	sl.Info("info")

	sl = slog.New(slog.DiscardHandler)
	sl.Info("info") //discard

	o := outer{}
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
