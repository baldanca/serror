package serror

import (
	"errors"
	"fmt"
)

type (
	// Error model
	Error struct {
		errType string
		data    []interface{}
		err     error
	}
)

// New error function
func New(i ...interface{}) *Error {
	return &Error{data: i}
}

// Build the error object
func (e *Error) Error() *Error {
	var errorString string
	for _, d := range e.data {
		errorString = fmt.Sprintf("%s %v", errorString, d)
	}
	e.err = errors.New(errorString)
	return e
}

// Recover function
func Recover(i interface{}) *Error {
	if e, ok := i.(*Error); ok {
		e.errType = "error"
		return e
	}
	return &Error{"panic", []interface{}{i}, nil}
}

// Get a error from panic
func Get(f func()) (e error) {
	defer func() {
		r := recover()
		if r, ok := r.(*Error); ok {
			e = r.Error().err
		}
		e = fmt.Errorf("%v", r)
	}()
	f()
	return nil
}
