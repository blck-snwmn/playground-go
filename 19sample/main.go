package main

import (
	"encoding"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// Add Duration.Abs method
func sampleTime() {
	l := time.Date(2022, 12, 1, 12, 13, 10, 0, time.Local)
	r := time.Date(2022, 12, 4, 12, 13, 10, 0, time.Local)
	if l.Sub(r).Abs() == r.Sub(l).Abs() {
		fmt.Println("|l-r|==|r-l|")
	}
}

func sampleAtmic() {
	var i uint64
	atomic.AddUint64(&i, 100)
	fmt.Println(atomic.LoadUint64(&i))

	var ii atomic.Uint64
	// use atomic.AddUint64 in Add method
	ii.Add(101)
	fmt.Println(ii.Load())
	ii.Add(2)
	fmt.Println(ii.Load())
}

func sampleSort() {
	input := []string{"aaa", "aaa", "aaa", "abc", "acb", "acb", "xxx"}

	for _, target := range []string{"abc", "xxa"} {
		fmt.Printf("-----\nsearch word is %q\n", target)
		fmt.Println("find start")
		fi, ok := sort.Find(len(input), func(i int) int {
			fmt.Printf("\tindex is %d\n", i)
			ii := input[i]
			if target == ii {
				return 0
			}
			if target < ii {
				return -1
			}
			return 1
		})
		fmt.Println("find:", fi, ok)

		fmt.Println("serch start")
		si := sort.Search(len(input), func(i int) bool {
			fmt.Printf("\tindex is %d\n", i)
			return input[i] >= target
		})
		fmt.Println("serch:", si)
	}
}

func sampleFmtAppend() {
	s := "aaaaa"
	b := []byte(s)
	b = fmt.Append(b, "bbb", "c")
	b = fmt.Appendf(b, "zzzz%d", 1)

	fmt.Println(string(b))
}

func sampleJoinPath() {
	endpoint := "https://example.com"
	endpoint, err := url.JoinPath(endpoint, "aaaa", "bbb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endpoint)

	endpoint, err = url.JoinPath(endpoint, "..", "..", "ccc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endpoint)

	endpoint, err = url.JoinPath(endpoint, "ddd", "..", "eee")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endpoint)
}

func wrapper(delimiter string, f func()) {
	const template = "[%s:%s]================\n"
	fmt.Printf(template, delimiter, "start")
	defer fmt.Printf(template, delimiter, "end")
	f()
}

type input struct {
	id   int
	name string
}

// MarshalText implements encoding.TextMarshaler
func (i *input) MarshalText() (text []byte, err error) {
	return []byte(fmt.Sprintf("name-%s:id=%d", i.name, i.id)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (i *input) UnmarshalText(text []byte) error {
	texts := strings.Split(string(text), ":")
	if len(texts) != 2 {
		return errors.New("invalid input text")
	}
	for _, t := range texts {
		kv := strings.Split(t, "-")
		if len(kv) != 2 {
			return fmt.Errorf("unexpected format: %s", t)
		}
		switch kv[0] {
		case "id":
			var err error
			i.id, err = strconv.Atoi(kv[1])
			if err != nil {
				return err
			}
		case "name":
			i.name = kv[1]
		default:
			return fmt.Errorf("unexpected key name: %s", kv[0])
		}
	}
	return nil
}

var _ encoding.TextUnmarshaler = (*input)(nil)
var _ encoding.TextMarshaler = (*input)(nil)

func main() {
	var i input
	flag.TextVar(&i, "input", &input{10, "x"}, "test")
	flag.Parse()
	wrapper("text marshal", func() {
		fmt.Printf("%#v\n", i)
	})
	wrapper("time", sampleTime)
	wrapper("atomic", sampleAtmic)
	wrapper("sort", sampleSort)
	wrapper("fmt.append", sampleFmtAppend)
	wrapper("url.join_path", sampleJoinPath)

}
