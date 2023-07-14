package errors

import "errors"

// These variables are used to give access to underlying standard library
// errors functions. We can wrap in custom functionality as needed or
// mock them during testing.
var (
	As = errors.As
	Is = errors.Is
)
