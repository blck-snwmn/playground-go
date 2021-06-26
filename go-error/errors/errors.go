package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

// ParentError ...
type ParentError interface {
	error
	Description() string
}

var _ ParentError = (*InvalidInputError)(nil)

func (iie *InvalidInputError) Error() string {
	return iie.Description()
}

// Description ...
func (iie *InvalidInputError) Description() string {
	return fmt.Sprintf("input=%s: %s", iie.input, iie.errMsg)
}

// InvalidInputError ...
type InvalidInputError struct {
	input  string
	errMsg string
}

var _ ParentError = (*SystemError)(nil)

func (se *SystemError) Error() string {
	return se.Description()
}

// Description ...
func (se *SystemError) Description() string {
	return fmt.Sprintf("%d: system error", se.code)
}

// SystemError ...
type SystemError struct {
	code uint64
}

type ValidateError struct {
	input          string
	validationName string
	errMsg         string
}

// ErrorDescription return error code and description
func ErrorDescription(err error) (uint64, string) {
	if err == nil {
		return 0, ""
	}
	var pe ParentError
	if xerrors.As(err, &pe) {
		switch e := pe.(type) {
		case *InvalidInputError:
			return 12, e.Description()
		case *SystemError:
			return 55, e.Description()
		}
	}
	return 99, "unexpected error"
}

// ErrorDescription return error code and description
func ErrorDescription2(err error) (uint64, string) {
	if err == nil {
		return 0, ""
	}
	var (
		iie = &InvalidInputError{}
		se  = &SystemError{}
	)
	if xerrors.As(err, &iie) {
		return 12, iie.Description()
	}
	if xerrors.As(err, &se) {
		return 55, se.Description()
	}
	return 99, "unexpected error"
}
