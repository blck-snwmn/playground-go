package errors

import "fmt"

// ParentError ...
type ParentError interface {
	error
	mustInpml()
}

var _ ParentError = (*InvalidInputError)(nil)

func (iie *InvalidInputError) Error() string {
	return fmt.Sprintf("input=%s: %s", iie.input, iie.errMsg)
}

func (*InvalidInputError) mustInpml() {}

// InvalidInputError ...
type InvalidInputError struct {
	input  string
	errMsg string
}

var _ ParentError = (*SystemError)(nil)

func (se *SystemError) Error() string {
	return fmt.Sprintf("%d: system error", se.code)
}

func (*SystemError) mustInpml() {}

// SystemError ...
type SystemError struct {
	code uint64
}
