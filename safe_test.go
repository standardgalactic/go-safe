package safe

import (
	"errors"
	"fmt"
	"testing"
)

type customError struct {
	msg string
}

func (s *customError) Error() string {
	return s.msg
}

type throwData struct {
	value string
}

func TestDo(t *testing.T) {
	var (
		errReturn      = errors.New("return error")
		errPanic       = errors.New("panic error")
		errPanicString = errors.New("panic string error")
		errPanicCustom = &customError{msg: "panic custom error"}
		errPanicStruct = fmt.Errorf("%#v", &throwData{value: "message"})
	)
	testCases := []struct {
		name  string
		fn    func() error
		equal func(got error) bool
	}{
		{
			name: "ReturnError",
			fn: func() error {
				return errReturn
			},
			equal: func(got error) bool {
				return errors.Is(got, errReturn)
			},
		},
		{
			name: "PanicError",
			fn: func() error {
				panic(errPanic)
			},
			equal: func(got error) bool {
				return errors.Is(got, errPanic)
			},
		},
		{
			name: "PanicCustomError",
			fn: func() error {
				panic(errPanicCustom)
			},
			equal: func(got error) bool {
				return errors.Is(got, errPanicCustom)
			},
		},
		{
			name: "PanicString",
			fn: func() error {
				panic(errPanicString.Error())
			},
			equal: func(got error) bool {
				return got.Error() == errPanicString.Error()
			},
		},
		{
			name: "PanicStruct",
			fn: func() error {
				panic(&throwData{value: "message"})
			},
			equal: func(got error) bool {
				return got.Error() == errPanicStruct.Error()
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := Func(tc.fn)(); !tc.equal(err) {
				t.Error("the error must match")
			}
		})
	}
}
