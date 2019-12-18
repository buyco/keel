package helper

import (
	"errors"
	"fmt"
)

// ErrorPrintf returns go error from formater
func ErrorPrintf(str string, args ...interface{}) error {
	return fmt.Errorf(str, args...)
}

// ErrorPrint returns error from string
func ErrorPrint(str string) error {
	return errors.New(str)
}
