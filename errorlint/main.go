package main

import (
	"errors"
	"fmt"
)

type sampleError struct{}

func (sampleError) Error() string {
	return "sample error"
}

var errSample = errors.New("sample error")

func main() {
	err := do()
	if err == errSample {
		panic("err != errSample")
	}
	if err.Error() != "sample error" { // no check
		panic("err.Error() != errSample.Error()")
	}
	switch err {
	case errSample:
		panic("err != errSample")
	}
	switch err.(type) {
	case sampleError:
		panic("err.(type) == sampleError")
	}
	err = fmt.Errorf("sample error: %v", err)
}

func do() error {
	return nil
}
