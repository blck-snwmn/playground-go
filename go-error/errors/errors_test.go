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

func TestErrorDescription(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 string
	}{
		{
			name: "nil is return 0 & no description",
			args: args{
				err: nil,
			},
			want:  0,
			want1: "",
		},
		{
			name: "InvalidInputError is return 12 & own's description",
			args: args{
				err: &InvalidInputError{input: "xxx", errMsg: "unsuport"},
			},
			want:  12,
			want1: "input=xxx: unsuport",
		},
		{
			name: "InvalidInputError is return 55 & own's description",
			args: args{
				err: &SystemError{code: 1002},
			},
			want:  55,
			want1: "1002: system error",
		},
		{
			name: "wrapped InvalidInputError is return 55 & own's description",
			args: args{
				err: xerrors.Errorf("wrap: %w", &SystemError{code: 1002}),
			},
			want:  55,
			want1: "1002: system error",
		},
		{
			name: "undefined ParentError is return 99 & common description",
			args: args{
				err: &testError{},
			},
			want:  99,
			want1: "unexpected error",
		},
		{
			name: "undefined erro is return 99 & common description",
			args: args{
				err: xerrors.New("test"),
			},
			want:  99,
			want1: "unexpected error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ErrorDescription(tt.args.err)
			if got != tt.want {
				t.Errorf("ErrorDescription() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ErrorDescription() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

type testError struct {
}

var _ ParentError = (*testError)(nil)

func (*testError) Error() string { return "" }

// Description ...
func (*testError) Description() string { return "" }
