package config

import (
	"fmt"
	"net/http"
)

// ErrorType represents the type of error
type ErrorType int

const (
	// ErrorTypeUnknown represents an unknown error
	ErrorTypeUnknown ErrorType = iota
	// ErrorTypeNetwork represents a network connectivity error
	ErrorTypeNetwork
	// ErrorTypeAPI represents an API response error
	ErrorTypeAPI
	// ErrorTypeValidation represents a validation error
	ErrorTypeValidation
	// ErrorTypeAuthentication represents an authentication error
	ErrorTypeAuthentication
	// ErrorTypeRateLimit represents a rate limit error
	ErrorTypeRateLimit
	// ErrorTypeTimeout represents a timeout error
	ErrorTypeTimeout
	// ErrorTypeSerialization represents a serialization/deserialization error
	ErrorTypeSerialization
)

// String returns the string representation of the error type
func (t ErrorType) String() string {
	switch t {
	case ErrorTypeNetwork:
		return "NetworkError"
	case ErrorTypeAPI:
		return "APIError"
	case ErrorTypeValidation:
		return "ValidationError"
	case ErrorTypeAuthentication:
		return "AuthenticationError"
	case ErrorTypeRateLimit:
		return "RateLimitError"
	case ErrorTypeTimeout:
		return "TimeoutError"
	case ErrorTypeSerialization:
		return "SerializationError"
	default:
		return "UnknownError"
	}
}

// InstanaError is the base error type for all Instana client errors
type InstanaError struct {
	Type       ErrorType
	Message    string
	StatusCode int
	Err        error
	Retryable  bool
	Temporary  bool
}

// Error implements the error interface
func (e *InstanaError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("[%s] HTTP %d: %s", e.Type.String(), e.StatusCode, e.Message)
	}
	return fmt.Sprintf("[%s] %s", e.Type.String(), e.Message)
}

// Unwrap returns the wrapped error
func (e *InstanaError) Unwrap() error {
	return e.Err
}

// IsRetryable returns true if the error is retryable
func (e *InstanaError) IsRetryable() bool {
	return e.Retryable
}

// IsTemporary returns true if the error is temporary
func (e *InstanaError) IsTemporary() bool {
	return e.Temporary
}

// NetworkError creates a new network error
func NetworkError(message string, err error) *InstanaError {
	return &InstanaError{
		Type:      ErrorTypeNetwork,
		Message:   message,
		Err:       err,
		Retryable: true,
		Temporary: true,
	}
}

// APIError creates a new API error
func APIError(statusCode int, message string, err error) *InstanaError {
	retryable := isRetryableStatusCode(statusCode)
	return &InstanaError{
		Type:       ErrorTypeAPI,
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
		Retryable:  retryable,
		Temporary:  retryable,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string, err error) *InstanaError {
	return &InstanaError{
		Type:      ErrorTypeValidation,
		Message:   message,
		Err:       err,
		Retryable: false,
		Temporary: false,
	}
}

// AuthenticationError creates a new authentication error
func AuthenticationError(message string, err error) *InstanaError {
	return &InstanaError{
		Type:       ErrorTypeAuthentication,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Err:        err,
		Retryable:  false,
		Temporary:  false,
	}
}

// RateLimitError creates a new rate limit error
func RateLimitError(message string, retryAfter int) *InstanaError {
	return &InstanaError{
		Type:       ErrorTypeRateLimit,
		Message:    fmt.Sprintf("%s (retry after %d seconds)", message, retryAfter),
		StatusCode: http.StatusTooManyRequests,
		Retryable:  true,
		Temporary:  true,
	}
}

// TimeoutError creates a new timeout error
func TimeoutError(message string, err error) *InstanaError {
	return &InstanaError{
		Type:      ErrorTypeTimeout,
		Message:   message,
		Err:       err,
		Retryable: true,
		Temporary: true,
	}
}

// SerializationError creates a new serialization error
func SerializationError(message string, err error) *InstanaError {
	return &InstanaError{
		Type:      ErrorTypeSerialization,
		Message:   message,
		Err:       err,
		Retryable: false,
		Temporary: false,
	}
}

// isRetryableStatusCode determines if an HTTP status code is retryable
func isRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusRequestTimeout, // 408
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}

// IsRetryableError checks if an error is retryable
func IsRetryableError(err error) bool {
	if err == nil {
		return false
	}

	if instanaErr, ok := err.(*InstanaError); ok {
		return instanaErr.IsRetryable()
	}

	return false
}

// IsTemporaryError checks if an error is temporary
func IsTemporaryError(err error) bool {
	if err == nil {
		return false
	}

	if instanaErr, ok := err.(*InstanaError); ok {
		return instanaErr.IsTemporary()
	}

	// Check if error implements Temporary() method (net.Error interface)
	if temp, ok := err.(interface{ Temporary() bool }); ok {
		return temp.Temporary()
	}

	return false
}

// ExtractStatusCode extracts the HTTP status code from an error
func ExtractStatusCode(err error) int {
	if err == nil {
		return 0
	}

	if instanaErr, ok := err.(*InstanaError); ok {
		return instanaErr.StatusCode
	}

	return 0
}

// WrapError wraps an error with additional context
func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}

	if instanaErr, ok := err.(*InstanaError); ok {
		instanaErr.Message = fmt.Sprintf("%s: %s", message, instanaErr.Message)
		return instanaErr
	}

	return fmt.Errorf("%s: %w", message, err)
}

// Made with Bob
