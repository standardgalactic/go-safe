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
		errThrow     = errors.New("throw error")
		errPanic     = errors.New("panic error")
		errString    = errors.New("string error")
		errCustom    = &customError{msg: "custom error"}
		data         = &throwData{value: "message"}
		errThrowData = fmt.Errorf("%#v", data)
	)
	var (
		errThrowFunc = Func(func() error {
			return errThrow
		})
		errPanicFunc = Func(func() error {
			panic(errPanic)
		})
		errCustomFunc = Func(func() error {
			panic(errCustom)
		})
		errStringFunc = Func(func() error {
			panic(errString.Error())
		})
		throwDataFunc = Func(func() error {
			panic(data)
		})
	)

	t.Run("case=throw-error", func(t *testing.T) {
		if err := errThrowFunc(); err != errThrow {
			t.Errorf("got: '%v', want: '%v'", err, errThrow)
		}
	})

	t.Run("case=panic-error", func(t *testing.T) {
		if err := errPanicFunc(); err != errPanic {
			t.Errorf("got: '%v', want: '%v'", err, errPanic)
		}
	})

	t.Run("case=custom-error", func(t *testing.T) {
		if err := errCustomFunc(); err != errCustom {
			t.Errorf("got: '%v', want: '%v'", err, errCustom)
		}
	})

	t.Run("case=string-error", func(t *testing.T) {
		if err := errStringFunc(); err.Error() != errString.Error() {
			t.Errorf("got '%v', want '%v'", err.Error(), errString.Error())
		}
	})

	t.Run("case=throw-data", func(t *testing.T) {
		if err := throwDataFunc(); err.Error() != errThrowData.Error() {
			t.Errorf("got: '%#v', want: '%#v'", err.Error(), errThrowData.Error())
		}
	})
}
