package main

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"runtime/secret"
)

type MyError int

func (e MyError) Error() string {
	return fmt.Sprintf("MyError code: %d", e)
}

const (
	MyErrorOne MyError = iota + 1
)

func doSomething() error {
	return fmt.Errorf("an error occurred: %w", MyErrorOne)
}

func generateID() int {
	id := 42
	return id
}

type MyStruct struct {
	ID *int
}

func main() {
	fmt.Println("Go 1.26 features demo")
	fmt.Println("======================= new =======================")
	x := []MyStruct{
		{
			ID: new(10),
		},
		{
			ID: new(generateID()),
		},
	}
	for _, v := range x {
		fmt.Println("MyStruct ID:", *v.ID)
	}

	fmt.Println("======================= errors =======================")
	// errors.As, errors.AsType
	var myErr MyError
	err := doSomething()
	if err != nil && errors.As(err, &myErr) {
		fmt.Println("Caught MyError:", myErr)
	}
	if e, ok := errors.AsType[MyError](err); ok {
		fmt.Println("Caught MyError using AsType:", e)
	}

	fmt.Println("======================= bytes.Buffer.Peek =======================")
	b := bytes.NewBufferString("test,buffer,data")
	l, err := b.ReadString(',')
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
	} else {
		fmt.Println("Read from buffer:", l)
	}

	by, _ := b.Peek(2)
	fmt.Println("Peeked bytes:", string(by))

	by = make([]byte, 2)
	n, err := b.Read(by)
	if err != nil {
		fmt.Println("Error reading bytes from buffer:", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(by))
	}

	l, err = b.ReadString(',')
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
	} else {
		fmt.Println("Read from buffer:", l)
	}
	fmt.Println("======================= slog.NewMultiHandler =======================")
	h1 := slog.NewJSONHandler(os.Stdout, nil)
	h2 := slog.NewTextHandler(os.Stdout, nil)
	mh := slog.NewMultiHandler(h1, h2)
	logger := slog.New(mh)
	logger.Info("This is a test log message", "key1", "value1", "key2", 42)

	fmt.Println("======================= runtime/secret =======================")
	// supported on linux/amd64, linux/arm64 only.
	secret.Do(func() {
		fmt.Println("Inside secret.Do")
	})
}
