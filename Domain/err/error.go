package err

import (
	"fmt"
)

const (
	Unauthorized = "Unauthorized"
	Validation   = "Validation"  // Validation error
	Conflict     = "Conflict"    // Conflict error
	Unexpected   = "ServerError" // Unexpected server error
	NotFound     = "NotFound"    // Resource not found error
)

// Error represents a custom domain error with a type and message.
type Error struct {
	kind    string
	Message string
}

// New creates a new Error with the given type and message.
func new(errType string, message string) *Error {
	return &Error{kind: errType, Message: message}
}

// Error returns the string representation of the Error.
func (e *Error) Error() string {
	return fmt.Sprintf(e.Message)
}

// Type returns the string of the Error.
func (e *Error) Type() string {
	return e.kind
}

// NewValidation creates a new validation error with the given message.
func NewValidation(message string) *Error {
	return new(Validation, message)
}

// NewConflict creates a new conflict error with the given message.
func NewConflict(message string) *Error {
	return new(Conflict, message)
}
func NewUnauthorized(message string) *Error {
	return new(Unauthorized, message)
}

// NewUnexpected creates a new unexpected server error with the given message.
func NewUnexpected(message string) *Error {
	return new(Unexpected, message)
}

// NewNotFound creates a new not found error with the given message.
func NewNotFound(message string) *Error {
	return new(NotFound, message)
}
