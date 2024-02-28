package fault

import (
	"fmt"
)

// RequiredFieldError is returned if a required field is not set.
type RequiredFieldError struct {
	Type  string
	Field string
}

// Error message of the RequiredFieldError
func (err *RequiredFieldError) Error() string {
	return fmt.Sprintf("field '%s' is required for %s", err.Field, err.Type)
}

// NewRequiredError creates a new RequiredFieldError.
// typeName is the name of the type / class, used to determine where the error was thrown.
// fieldName name of the field that was not set.
func NewRequiredError(typeName, fieldName string) *RequiredFieldError {
	return &RequiredFieldError{
		Type:  typeName,
		Field: fieldName,
	}
}

// SetFieldError is returned if a field is set but should not be.
type SetFieldError struct {
	Type  string
	Field string
}

// Error message of the SetFieldError
func (err *SetFieldError) Error() string {
	return fmt.Sprintf("field '%s' is set but should not be for %s", err.Field, err.Type)
}

// NewSetFieldError creates a new SetFieldError.
// typeName is the name of the type / class, used to determine where the error was thrown.
// fieldName name of the field that was set.
func NewSetFieldError(typeName, fieldName string) *SetFieldError {
	return &SetFieldError{
		Type:  typeName,
		Field: fieldName,
	}
}
