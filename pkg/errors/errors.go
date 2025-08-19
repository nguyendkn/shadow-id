package errors

import (
	"fmt"
)

// ErrorType represents the type of error
type ErrorType string

const (
	ErrorTypeValidation ErrorType = "validation"
	ErrorTypeNotFound   ErrorType = "not_found"
	ErrorTypeConflict   ErrorType = "conflict"
	ErrorTypeInternal   ErrorType = "internal"
	ErrorTypeExternal   ErrorType = "external"
)

// AppError represents an application error with additional context
type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    string    `json:"code,omitempty"`
	Cause   error     `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Cause
}

// New creates a new application error
func New(errorType ErrorType, message string) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
	}
}

// Wrap wraps an existing error with additional context
func Wrap(err error, message string) *AppError {
	if err == nil {
		return nil
	}
	
	// If it's already an AppError, wrap it
	if appErr, ok := err.(*AppError); ok {
		return &AppError{
			Type:    appErr.Type,
			Message: message,
			Cause:   appErr,
		}
	}
	
	return &AppError{
		Type:    ErrorTypeInternal,
		Message: message,
		Cause:   err,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeNotFound,
		Message: message,
	}
}

// NewConflictError creates a new conflict error
func NewConflictError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeConflict,
		Message: message,
	}
}

// NewInternalError creates a new internal error
func NewInternalError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeInternal,
		Message: message,
	}
}

// IsType checks if the error is of a specific type
func IsType(err error, errorType ErrorType) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == errorType
	}
	return false
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	return IsType(err, ErrorTypeValidation)
}

// IsNotFoundError checks if the error is a not found error
func IsNotFoundError(err error) bool {
	return IsType(err, ErrorTypeNotFound)
}

// IsConflictError checks if the error is a conflict error
func IsConflictError(err error) bool {
	return IsType(err, ErrorTypeConflict)
}

// IsInternalError checks if the error is an internal error
func IsInternalError(err error) bool {
	return IsType(err, ErrorTypeInternal)
}
