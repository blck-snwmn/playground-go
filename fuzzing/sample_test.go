package fuzzing

import (
	"errors"
	"fmt"
	"testing"
)

func Foo(s string, i int) error {
	if i > 100 {
		return errors.New("error")
	}
	return nil
}

func FuzzFoo(f *testing.F) {
	f.Add("aaa", 10)
	f.Fuzz(func(t *testing.T, s string, i int) {
		fmt.Println(s, i)
		if err := Foo(s, i); err != nil {
			t.Errorf("failed input is = (%v, %v)", s, i)
		}
	})
}
