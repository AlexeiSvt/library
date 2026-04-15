package errors

import "errors"

var (
	ErrNegativeFactorial = errors.New("factorial of a negative number is undefined")
	ErrDivisionByZero = errors.New("division by zero")

)

var (
	ErrEmptySlice = errors.New("slice is empty")
)

var (
    ErrEmptyString = errors.New("string is empty")
    ErrEmptyInput  = errors.New("one or both strings are empty")
)