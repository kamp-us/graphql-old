package errors

import (
	"bytes"
	"errors"
	"fmt"
)

type Errors []error

func (e Errors) Err() error {
	if len(e) == 0 {
		return nil
	}

	return e
}

func (e Errors) Error() string {
	var buf bytes.Buffer

	if n := len(e); n == 1 {
		buf.WriteString("1 error: ")
	} else {
		fmt.Fprintf(&buf, "%d errors: ", n)
	}

	for i, err := range e {
		if i != 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(err.Error())
	}

	return buf.String()
}

func (e Errors) Slice() []error {
	return []error(e)
}

// This is convenience method so we don't have to fight with package imports.
func New(message string) error {
	return errors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
