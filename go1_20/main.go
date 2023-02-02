package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/multierr"
	"golang.org/x/xerrors"
)

func main() {
	// context
	{
		ctx := context.Background()
		ctx, cancel := context.WithCancelCause(ctx)
		e := errors.New("context error")
		go func() {
			cancel(e)
		}()
	tset:
		for {
			select {
			case <-ctx.Done():
				break tset
			}
		}
		err := context.Cause(ctx)
		fmt.Println(err)
		fmt.Println(err == e) // true
	}
	// err
	{
		err1 := xerrors.New("error1")
		err2 := xerrors.New("error2")
		err3 := xerrors.New("error3")
		errj := errors.Join(err1, err2, err3)

		fmt.Println("errj: ", errj)
		fmt.Printf("errj: %+v\n", errj) // no stacktrace. join call error's Error method
		fmt.Println(errors.Is(errj, err1))
		fmt.Println(errors.Is(errj, err2))
		fmt.Println(errors.Is(errj, err3))

		errc := multierr.Combine(err1, err2, err3) // with stacktrace
		fmt.Printf("multierr: %+v\n", errc)
	}
	// fmt
	{
		err1 := errors.New("first_error")
		err2 := errors.New("second_error")
		merr := fmt.Errorf("first err is %w, second err is %w", err1, err2)
		fmt.Println("merr: ", merr)
		fmt.Println(errors.Is(merr, err1))
		fmt.Println(errors.Is(merr, err2))
	}

	// strings
	fmt.Println(strings.TrimPrefix("aaabbbcccc", "aa"))
	fmt.Println(strings.CutPrefix("aaabbbcccc", "aa"))
	fmt.Println(strings.TrimSuffix("aaabbbcccc", "cc"))
	fmt.Println(strings.CutSuffix("aaabbbcccc", "cc"))
	// time
	n := time.Now()
	fmt.Println(n.Format(time.DateOnly))
	fmt.Println(n.Format(time.DateTime))
	fmt.Println(n.Format(time.TimeOnly))

	fmt.Println(time.Now().Compare(n)) // 1
}
