package oblate

import (
	"bytes"
	"errors"
)

type Error struct {
	errs []error
}

// New is the same as Join(errors.New(text), errs...).
func New(text string, errs ...error) error {
	arg := make([]error, 0, len(errs)+1)
	arg = append(arg, errors.New(text))
	arg = append(arg, errs...)
	return Join(arg...)
}

// Join returns an error that wraps the given errors.
// Any nil error values are discarded. Join returns nil if every value in errs is nil.
// The error formats as the strings obtained by calling the Error method of the first
// non-nil error in errs.
func Join(errs ...error) error {
	n := 0
	for _, v := range errs {
		if v != nil {
			n++
		}
	}
	if n == 0 {
		return nil
	}
	es := make([]error, 0, n)
	for _, v := range errs {
		if v != nil {
			es = append(es, v)
		}
	}
	return &Error{errs: es}
}

// Error returns the first non-nil of the error list.
func (e *Error) Error() string {
	return e.errs[0].Error()
}

// Cause returns the error formats as the concatenation of the strings obtained by calling
// the Error method of each element of errs except the first error, with a newline between
// each string.
func (e *Error) Cause() string {
	var b bytes.Buffer
	for i, v := range e.errs[1:] {
		if i != 0 {
			b.WriteString("\n")
		}
		b.WriteString(v.Error())
	}
	return b.String()
}

// Unwrap implements Wrapper interface.
func (e *Error) Unwrap() []error {
	return e.errs
}
