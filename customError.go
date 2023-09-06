package errors

import "fmt"

// ErrorType is a custom error type
type ErrorType uint

// ValidationErrorMessage default message for validation errors
const ValidationErrorMessage = "several validation errors occurred"

const (
	// NoType error without type
	NoType = ErrorType(iota)
	// Internal error type
	Internal
	// Validation error type
	Validation
	// NotFound error type
	NotFound
	// AlreadyExist error type
	AlreadyExist
	// Unauthorized error type
	Unauthorized
	// Forbidden error type
	Forbidden
)

var errorStrings = map[ErrorType]string{
	NoType:       "NoType",
	Internal:     "Internal",
	Validation:   "Validation",
	NotFound:     "NotFound",
	AlreadyExist: "AlreadyExist",
	Unauthorized: "Unauthorized",
	Forbidden:    "Forbidden",
}

var errorDescriptions = map[ErrorType]string{
	NoType:       "unrecognized error occurred",
	Internal:     "internal error occurred",
	Validation:   ValidationErrorMessage,
	NotFound:     "not found error occurred",
	AlreadyExist: "already exist error occurred",
	Unauthorized: "unauthorized error occurred",
	Forbidden:    "forbidden error occurred",
}

func (t ErrorType) String() string {
	if name, ok := errorStrings[t]; ok {
		return name
	}
	return fmt.Sprintf("%d", t)
}

func (t ErrorType) Description() string {
	if description, ok := errorDescriptions[t]; ok {
		return description
	}
	return t.String()
}

type customError struct {
	file      string
	line      int
	msg       string
	wrapped   error
	errorType ErrorType
	context   map[string]string
}

// Error returns the error message
//
// Return
//   - string - error message
func (c customError) Error() string {
	return c.msg
}

// New creates a new error
//
// Params
//   - msg - error message
//
// Return
//   - error - new generated error
func (t ErrorType) New(msg string) error {
	return newErr(2, t, msg)
}

// Newf creates a new error with formatted message
//
// Params
//   - msg - format for error message
//   - args - arguments for message formatting
//
// Return
//   - error - new generated error
func (t ErrorType) Newf(msg string, args ...interface{}) error {
	return newErr(2, t, fmt.Sprintf(msg, args...))
}

// Wrapf wraps an error
//
// Params
//   - error - error to wrapping
//   - msg - format for error message
//   - args - arguments for message formatting
//
// Return
//   - error - new generated wrapped error
func (t ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return wrapErr(err, 2, t, fmt.Sprintf(msg, args...))
}

// Wrap wraps an error
//
// Params
//   - error - error to wrapping
//   - msg - error message
//
// Return
//   - error - new generated wrapped error
func (t ErrorType) Wrap(err error, msg string) error {
	return wrapErr(err, 2, t, msg)
}
