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

// NewRequiredError created a new RequiredFieldError.
// typeName is the name of the type / class, used to determine where the error was thrown.
// fieldName name of the field that was not set.
func NewRequiredError(typeName, fieldName string) *RequiredFieldError {
	return &RequiredFieldError{
		Type:  typeName,
		Field: fieldName,
	}
}
