package error_custom

import (
	"errors"
	"fmt"
)

//https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f
type CustomError struct {
	errorType    ErrorType
	wrappedError error
	errorContext errorContext
}

func (c CustomError) Error() string {
	return c.wrappedError.Error()
}

func (c CustomError) Stacktrace() string {
	return fmt.Sprint("%+v\n", c.wrappedError)
}
func (c CustomError) Type() int {
	return int(c.errorType)
}

func New(mess string) error {
	return CustomError{errorType: Error, wrappedError: errors.New(mess)}
}
