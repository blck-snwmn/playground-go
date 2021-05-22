package errors

import (
	"testing"

	"golang.org/x/xerrors"
)

func TestAs(t *testing.T) {
	var (
		me ParentError
		e  error
		// concrete
		iie = &InvalidInputError{input: "xxx", errMsg: "unsuport"}
		se  = &SystemError{code: 1002}
		ee  = xerrors.New("test")
	)
	if !xerrors.As(iie, &me) {
		t.Errorf("InvalidInputError does not assign to ParentError")
	} else if got := me.Error(); got != "input=xxx: unsuport" {
		t.Errorf("me.Error()=%s, want=%s", got, "input=xxx: unsuport")
	}
	if !xerrors.As(se, &me) {
		t.Errorf("SystemError does not assign to ParentError")
	} else if got := me.Error(); got != "1002: system error" {
		t.Errorf("me.Error()=%s, want=%s", got, "1002: system error")
	}
	if !xerrors.As(se, &e) {
		t.Errorf("SystemError does not assign to error")
	}
	if xerrors.As(ee, &me) {
		t.Errorf("xerrors's err assign to ParentError")
	}
}

func TestAsWithTypeSwithes(t *testing.T) {
	var (
		// concrete
		iie = &InvalidInputError{input: "xxx", errMsg: "unsuport"}
		se  = &SystemError{code: 1002}
	)
	var calledInvalid, calledSystem bool

	asFunc := func(err error) {
		var me ParentError
		if !xerrors.As(err, &me) {
			t.Errorf("input error does not assign to ParentError")
			return
		}
		switch e := me.(type) {
		case *InvalidInputError:
			if e.input != "xxx" || e.errMsg != "unsuport" {
				t.Errorf("got={input:%s, errMsg=%s}, want={input:%s, errMsg=%s}", e.input, e.errMsg, "xxx", "unsuport")
			}
			calledInvalid = true
		case *SystemError:
			if e.code != 1002 {
				t.Errorf("got=%d, want=%d", e.code, 1002)
			}
			calledSystem = true
		default:
			t.Errorf("input error assign to unexpected type")
		}
	}

	asFunc(iie)
	asFunc(se)

	if !calledInvalid {
		t.Errorf("InvalidInputError is not called")
	}
	if !calledSystem {
		t.Errorf("SystemError is not called")
	}
}
