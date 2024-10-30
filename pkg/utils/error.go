package utils

import (
	"reflect"
)

// PgError represents a custom error type for PostgreSQL errors.
type PgError struct {
	Message  string
	Code     string
	Detail   string
	Hint     string
	Position string
}

// Error implements the error interface for PgError.
func (e *PgError) Error() string {
	return e.Message
}

// NewPgError creates a new PgError with the given details.
func NewPgError(message, detail, hint, position, code string) *PgError {
	return &PgError{
		Message:  message,
		Code:     code,
		Detail:   detail,
		Hint:     hint,
		Position: position,
	}
}

// IsSuperset checks if obj has all the fields of the baseStruct type.
func IsSuperset(obj interface{}, baseStruct interface{}) bool {
	baseValue := reflect.ValueOf(baseStruct)
	objValue := reflect.ValueOf(obj)

	// Check if obj has all the fields that baseStruct has
	for i := 0; i < baseValue.NumField(); i++ {
		fieldName := baseValue.Type().Field(i).Name
		if _, ok := objValue.Type().FieldByName(fieldName); !ok {
			return false
		}
	}

	return true
}
