package main

import (
	"bytes"
	"cmp"
	"context"
	"embed"
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"os"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"log/slog"
)

//go:embed **.**
var fs embed.FS

func main() {
	{
		// log/slog
		fmt.Println("===========log/slog===========")
		l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

		l.Info("use json handler", slog.Bool("boolkey", true))
		slog.Info("before SetDefault", slog.Bool("boolkey", true))

		slog.SetDefault(l.With(slog.String("withkey", "withvalue")))

		slog.Warn("warn")
	}
	{
		// cmp
		fmt.Println("===========cmp===========")
		fmt.Printf("1 cmp 2%d\n", cmp.Compare(1, 2))
		fmt.Printf("1 cmp 1%d\n", cmp.Compare(1, 1))
		fmt.Printf("3 cmp 2%d\n", cmp.Compare(3, 2))
	}
	{
		// slices
		fmt.Println("===========slices===========")
		fmt.Printf("max=%d\n", slices.Max([]int{1, 2, 100, 3}))
		fmt.Printf("max=%d\n", max(1, 2, 100, 3))
		fmt.Printf("max=%d\n", slices.Min([]int{1, 2, 0, 3}))
		fmt.Printf("max=%d\n", min(1, 2, 0, 3))
		var (
			index int
			ok    bool
		)
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 15, 18, 20, 21, 22, 23, 24, 25, 26, 27, 40, 50, 60, 70, 80, 90, 100}
		sort.Ints(input)
		fmt.Println(input)
		index, ok = slices.BinarySearch(input, 25)
		fmt.Printf("exists?`%v`, search=%v(index=%d)\n", ok, input[index], index)

		index, ok = slices.BinarySearch(input, 11)
		fmt.Printf("exists?`%v`, search=%v(index=%d)\n", ok, input[index], index)

		input = []int{1, 2, 100, 3}
		fmt.Printf("[before]sort=%v\n", input)
		slices.Sort(input)
		fmt.Printf("[after]sort=%v\n", input)
		slices.Reverse(input)
		fmt.Printf("reverse=%v\n", input)
		input = slices.Replace(input, 1, 3, 10, 20, 30, 40, 50)
		fmt.Printf("replace=%v\n", input) // [100 10 20 30 40 50 1]

		input = slices.Insert(input, 1, 11, 12, 13)
		fmt.Printf("insert=%v\n", input) // [100 11 12 13 10 20 30 40 50 1]

		input = slices.Delete(input, 1, 5)
		fmt.Printf("delete=%v\n", input) // [100 20 30 40 50 1]
	}
	{
		// maps
		fmt.Println("===========maps===========")
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		mc := maps.Clone(m)
		fmt.Printf("origin=%p, clone=%p\n", m, mc)
		fmt.Printf("equal=%v\n", maps.Equal(m, mc))
		m["a"] = 10
		m["d"] = 123
		m["e"] = 321
		fmt.Printf("this=%v, other=%v\n", m, mc)
		maps.Copy(mc, m)
		fmt.Printf("this=%v, other=%v(copied)\n", m, mc)

		fmt.Printf("keys=%v\n", maps.Keys(m))     // indeterminate order
		fmt.Printf("values=%v\n", maps.Values(m)) // indeterminate order
		maps.DeleteFunc(mc, func(key string, value int) bool {
			return value%2 == 0
		})
		fmt.Printf("delete=%v\n", mc)
	}
	{
		// clear
		fmt.Println("===========clear===========")
		slice := []int{1, 2, 3, 4, 5}
		fmt.Printf("input=%v\n", slice)
		clear(slice)
		fmt.Printf("input=%v(cleared)\n", slice)

		m := map[string]int{"a": 1, "b": 2}
		fmt.Printf("input=%v\n", m)
		clear(m)
		fmt.Printf("input=%v(cleared)\n", m)

		mNaN := map[float64]int{math.NaN(): 10, math.Inf(0): 100, math.Inf(-1): 20}
		fmt.Printf("input=%v\n", mNaN)
		mNaN[math.NaN()] = 12
		mNaN[math.Inf(0)] = 13
		mNaN[math.Inf(-1)] = 14
		fmt.Printf("input=%v(changed)\n", mNaN) // NaN is duplicated
		clear(mNaN)
		fmt.Printf("input=%v(cleared)\n", mNaN) // clear NaN key
	}
	{
		// strings/bytes
		fmt.Println("===========strings/bytes===========")
		fmt.Printf("ContainsFunc=%v\n", strings.ContainsFunc("abcd", func(r rune) bool {
			return r == 'a'
		}))
		fmt.Printf("ContainsFunc=%v\n", bytes.ContainsFunc([]byte{0x00, 0x02, 0x05}, func(r rune) bool {
			return r%2 == 1
		}))
		fmt.Printf("ContainsFunc=%v\n", bytes.ContainsFunc([]byte{0x00, 0x02, 0x04}, func(r rune) bool {
			return r%2 == 1
		}))
	}
	{
		// context
		fmt.Println("===========context===========")
		ctx := context.Background()
		fmt.Printf("ctx.error=%v\n", ctx.Err())
		ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond))
		defer cancel()
		<-ctx.Done()

		fmt.Printf("ctx.error=%v\n", ctx.Err())

		key := struct{}{}
		ctx = context.WithValue(ctx, key, "value")
		fmt.Printf("ctx.error=%v, value=%v\n", ctx.Err(), ctx.Value(key))

		ctx, cancel = context.WithDeadlineCause(ctx, time.Now().Add(-time.Millisecond), errors.New("text string")) // ctx.error=context deadline exceeded, cause=context deadline exceeded
		defer cancel()
		fmt.Printf("ctx.error=%+v, cause=%+v\n", ctx.Err(), context.Cause(ctx))

		ctx = context.WithoutCancel(ctx)
		fmt.Printf("ctx.error=%v, value=%v\n", ctx.Err(), ctx.Value(key)) // ctx.error=context deadline exceeded, cause=text string
	}
	{
		// context.WithDeadlineCause
		fmt.Println("===========context.WithDeadlineCause===========")
		ctx := context.Background()
		ctx, cancel := context.WithDeadlineCause(ctx, time.Now().Add(-time.Millisecond), errors.New("text string"))
		defer cancel()
		fmt.Printf("ctx.error=%+v, cause=%+v\n", ctx.Err(), context.Cause(ctx))
	}
	{
		// context.AfterFunc
		fmt.Println("===========context.AfterFunc===========")
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		_ = context.AfterFunc(ctx, func() {
			fmt.Println("stop in AfterFunc callback")
		})

		cancel()
		time.Sleep(time.Second)
	}
	{
		// sync
		fmt.Println("===========sync===========")
		var defaultValue int
		defaultValue = 10
		ot := sync.OnceValue(func() int { return defaultValue })
		defaultValue = 20
		fmt.Printf("value=%v\n", ot())
		defaultValue = 30
		fmt.Printf("value=%v\n", ot())

		ovs := sync.OnceValues(func() (int, int) {
			return defaultValue / 2, defaultValue / 3
		})
		defaultValue = 40
		l, r := ovs()
		fmt.Printf("values=(%v, %v)\n", l, r)
		l, r = ovs()
		fmt.Printf("values=(%v, %v)\n", l, r)
		defaultValue = r

		of := sync.OnceFunc(func() {
			defaultValue = defaultValue + 40
		})
		fmt.Printf("value=%v\n", defaultValue)
		of()
		fmt.Printf("value=%v\n", defaultValue)
		of()
		fmt.Printf("value=%v\n", defaultValue)
	}
	{
		// sync.OnceValue
		fmt.Println("===========sync.OnceValue===========")
		ot := newot()
		ot2 := ot

		fmt.Printf("now=%v\n", time.Now())
		fmt.Printf("time=%v\n", ot.Time())
		time.Sleep(time.Second)
		fmt.Printf("time=%v\n", ot2.Time())
	}
	{
		// embed
		// embed's openFile implements io.ReaderAt
		fmt.Println("===========embed===========")
		f, err := fs.Open("main.go")
		if err != nil {
			panic(err)
		}
		ra, ok := f.(io.ReaderAt)
		if !ok {
			panic("not io.ReaderAt")
		}
		buf := make([]byte, 10)
		n, err := ra.ReadAt(buf, 2)
		if err != nil {
			panic(err)
		}
		fmt.Printf("n=%v, buf=%v\n", n, string(buf))
	}
	{
		// testing
		fmt.Println("===========testing===========")
		fmt.Printf("isTest=%v\n", isTest())
	}
}

// See: https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/sync/waitgroup.go;l=24
type noCopy struct{}

var _ sync.Locker = (*noCopy)(nil)

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type onceTime struct {
	noCopy noCopy
	ot     func() time.Time
}

func newot() *onceTime {
	return &onceTime{ot: sync.OnceValue(func() time.Time { return time.Now() })}
}

func (o *onceTime) Time() time.Time {
	return o.ot()
}

func isTest() bool {
	return testing.Testing()
}
