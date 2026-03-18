package config

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestNetworkError(t *testing.T) {
	originalErr := errors.New("connection refused")
	err := NetworkError("failed to connect", originalErr)

	if err == nil {
		t.Fatal("NetworkError() returned nil")
	}

	if err.Type != ErrorTypeNetwork {
		t.Errorf("Expected error type %v, got %v", ErrorTypeNetwork, err.Type)
	}

	if err.Message != "failed to connect" {
		t.Errorf("Expected message 'failed to connect', got '%s'", err.Message)
	}

	if err.Err != originalErr {
		t.Error("Original error not preserved")
	}

	if !err.IsRetryable() {
		t.Error("Network errors should be retryable")
	}

	if !err.IsTemporary() {
		t.Error("Network errors should be temporary")
	}
}

func TestAPIError(t *testing.T) {
	err := APIError(500, "Internal Server Error", nil)

	if err == nil {
		t.Fatal("APIError() returned nil")
	}

	if err.Type != ErrorTypeAPI {
		t.Errorf("Expected error type %v, got %v", ErrorTypeAPI, err.Type)
	}

	if err.StatusCode != 500 {
		t.Errorf("Expected status code 500, got %d", err.StatusCode)
	}

	if err.Message != "Internal Server Error" {
		t.Errorf("Expected message 'Internal Server Error', got '%s'", err.Message)
	}

	if !err.IsRetryable() {
		t.Error("500 errors should be retryable")
	}
}

func TestAPIErrorNonRetryable(t *testing.T) {
	err := APIError(400, "Bad request", nil)

	if err.IsRetryable() {
		t.Error("400 errors should not be retryable")
	}

	if err.IsTemporary() {
		t.Error("400 errors should not be temporary")
	}
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("invalid configuration", nil)

	if err == nil {
		t.Fatal("NewValidationError() returned nil")
	}

	if err.Type != ErrorTypeValidation {
		t.Errorf("Expected error type %v, got %v", ErrorTypeValidation, err.Type)
	}

	if err.IsRetryable() {
		t.Error("Validation errors should not be retryable")
	}

	if err.IsTemporary() {
		t.Error("Validation errors should not be temporary")
	}
}

func TestAuthenticationError(t *testing.T) {
	err := AuthenticationError("invalid API token", nil)

	if err == nil {
		t.Fatal("AuthenticationError() returned nil")
	}

	if err.Type != ErrorTypeAuthentication {
		t.Errorf("Expected error type %v, got %v", ErrorTypeAuthentication, err.Type)
	}

	if err.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, err.StatusCode)
	}

	if err.IsRetryable() {
		t.Error("Authentication errors should not be retryable")
	}
}

func TestRateLimitError(t *testing.T) {
	err := RateLimitError("rate limit exceeded", 60)

	if err == nil {
		t.Fatal("RateLimitError() returned nil")
	}

	if err.Type != ErrorTypeRateLimit {
		t.Errorf("Expected error type %v, got %v", ErrorTypeRateLimit, err.Type)
	}

	if err.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status code %d, got %d", http.StatusTooManyRequests, err.StatusCode)
	}

	if !strings.Contains(err.Message, "60 seconds") {
		t.Errorf("Expected message to contain retry after duration, got: %s", err.Message)
	}

	if !err.IsRetryable() {
		t.Error("Rate limit errors should be retryable")
	}
}

func TestTimeoutError(t *testing.T) {
	err := TimeoutError("request timeout", nil)

	if err == nil {
		t.Fatal("TimeoutError() returned nil")
	}

	if err.Type != ErrorTypeTimeout {
		t.Errorf("Expected error type %v, got %v", ErrorTypeTimeout, err.Type)
	}

	if !err.IsRetryable() {
		t.Error("Timeout errors should be retryable")
	}

	if !err.IsTemporary() {
		t.Error("Timeout errors should be temporary")
	}
}

func TestSerializationError(t *testing.T) {
	originalErr := errors.New("json unmarshal error")
	err := SerializationError("failed to parse response", originalErr)

	if err == nil {
		t.Fatal("SerializationError() returned nil")
	}

	if err.Type != ErrorTypeSerialization {
		t.Errorf("Expected error type %v, got %v", ErrorTypeSerialization, err.Type)
	}

	if err.IsRetryable() {
		t.Error("Serialization errors should not be retryable")
	}
}

func TestInstanaErrorError(t *testing.T) {
	err := NetworkError("test error", nil)
	errMsg := err.Error()

	if !strings.Contains(errMsg, "NetworkError") {
		t.Errorf("Error message should contain error type, got: %s", errMsg)
	}

	if !strings.Contains(errMsg, "test error") {
		t.Errorf("Error message should contain message, got: %s", errMsg)
	}
}

func TestInstanaErrorErrorWithStatusCode(t *testing.T) {
	err := APIError(404, "Not Found", nil)
	errMsg := err.Error()

	if !strings.Contains(errMsg, "404") {
		t.Errorf("Error message should contain status code, got: %s", errMsg)
	}

	if !strings.Contains(errMsg, "Not Found") {
		t.Errorf("Error message should contain message, got: %s", errMsg)
	}
}

func TestInstanaErrorUnwrap(t *testing.T) {
	originalErr := errors.New("original error")
	err := NetworkError("test error", originalErr)

	unwrapped := errors.Unwrap(err)
	if unwrapped != originalErr {
		t.Error("Unwrap() should return the original error")
	}
}

func TestInstanaErrorUnwrapNil(t *testing.T) {
	err := NetworkError("test error", nil)

	unwrapped := errors.Unwrap(err)
	if unwrapped != nil {
		t.Error("Unwrap() should return nil when no wrapped error")
	}
}

func TestIsRetryableForDifferentStatusCodes(t *testing.T) {
	tests := []struct {
		statusCode int
		retryable  bool
	}{
		{408, true},  // Request Timeout
		{429, true},  // Too Many Requests
		{500, true},  // Internal Server Error
		{502, true},  // Bad Gateway
		{503, true},  // Service Unavailable
		{504, true},  // Gateway Timeout
		{400, false}, // Bad Request
		{401, false}, // Unauthorized
		{403, false}, // Forbidden
		{404, false}, // Not Found
		{200, false}, // OK (shouldn't be an error)
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			err := APIError(tt.statusCode, "test", nil)
			if err.IsRetryable() != tt.retryable {
				t.Errorf("Status code %d: expected retryable=%v, got %v",
					tt.statusCode, tt.retryable, err.IsRetryable())
			}
		})
	}
}

func TestErrorTypeString(t *testing.T) {
	tests := []struct {
		errorType ErrorType
		expected  string
	}{
		{ErrorTypeNetwork, "NetworkError"},
		{ErrorTypeAPI, "APIError"},
		{ErrorTypeValidation, "ValidationError"},
		{ErrorTypeAuthentication, "AuthenticationError"},
		{ErrorTypeRateLimit, "RateLimitError"},
		{ErrorTypeTimeout, "TimeoutError"},
		{ErrorTypeSerialization, "SerializationError"},
		{ErrorTypeUnknown, "UnknownError"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.errorType.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.errorType.String())
			}
		})
	}
}

func TestErrorChaining(t *testing.T) {
	// Test error chaining with errors.Is
	originalErr := errors.New("original")
	wrappedErr := NetworkError("wrapped", originalErr)

	if !errors.Is(wrappedErr, originalErr) {
		t.Error("errors.Is should work with wrapped errors")
	}
}

func TestErrorAs(t *testing.T) {
	// Test error type assertion with errors.As
	err := APIError(500, "test", nil)

	var instanaErr *InstanaError
	if !errors.As(err, &instanaErr) {
		t.Error("errors.As should work with InstanaError")
	}

	if instanaErr.Type != ErrorTypeAPI {
		t.Error("Error type not preserved through errors.As")
	}
}

func TestMultipleErrorWrapping(t *testing.T) {
	err1 := errors.New("base error")
	err2 := NetworkError("network issue", err1)
	err3 := TimeoutError("timeout occurred", err2)

	// Should be able to unwrap to original
	if !errors.Is(err3, err1) {
		t.Error("Should be able to find base error through multiple wraps")
	}

	// Should be able to find intermediate error
	var networkErr *InstanaError
	if !errors.As(err3, &networkErr) {
		t.Error("Should be able to find InstanaError in chain")
	}
}

func TestIsRetryableError(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		retryable bool
	}{
		{"nil error", nil, false},
		{"network error", NetworkError("test", nil), true},
		{"timeout error", TimeoutError("test", nil), true},
		{"rate limit error", RateLimitError("test", 60), true},
		{"API 500 error", APIError(500, "test", nil), true},
		{"API 400 error", APIError(400, "test", nil), false},
		{"validation error", NewValidationError("test", nil), false},
		{"authentication error", AuthenticationError("test", nil), false},
		{"standard error", errors.New("test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsRetryableError(tt.err) != tt.retryable {
				t.Errorf("Expected retryable=%v, got %v", tt.retryable, IsRetryableError(tt.err))
			}
		})
	}
}

func TestIsTemporaryError(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		temporary bool
	}{
		{"nil error", nil, false},
		{"network error", NetworkError("test", nil), true},
		{"timeout error", TimeoutError("test", nil), true},
		{"rate limit error", RateLimitError("test", 60), true},
		{"API 500 error", APIError(500, "test", nil), true},
		{"API 400 error", APIError(400, "test", nil), false},
		{"validation error", NewValidationError("test", nil), false},
		{"standard error", errors.New("test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsTemporaryError(tt.err) != tt.temporary {
				t.Errorf("Expected temporary=%v, got %v", tt.temporary, IsTemporaryError(tt.err))
			}
		})
	}
}

func TestExtractStatusCode(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		statusCode int
	}{
		{"nil error", nil, 0},
		{"API error 404", APIError(404, "test", nil), 404},
		{"API error 500", APIError(500, "test", nil), 500},
		{"authentication error", AuthenticationError("test", nil), http.StatusUnauthorized},
		{"rate limit error", RateLimitError("test", 60), http.StatusTooManyRequests},
		{"network error", NetworkError("test", nil), 0},
		{"standard error", errors.New("test"), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ExtractStatusCode(tt.err) != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, ExtractStatusCode(tt.err))
			}
		})
	}
}

func TestNilErrorHandling(t *testing.T) {
	// Test that constructors handle nil wrapped errors gracefully
	err := NetworkError("test", nil)
	if err.Err != nil {
		t.Error("Nil wrapped error should remain nil")
	}

	errMsg := err.Error()
	if strings.Contains(errMsg, "<nil>") {
		t.Error("Error message should not contain '<nil>'")
	}
}

func TestEmptyMessageHandling(t *testing.T) {
	err := NetworkError("", nil)
	errMsg := err.Error()

	if errMsg == "" {
		t.Error("Error message should not be empty even with empty input")
	}

	if !strings.Contains(errMsg, "NetworkError") {
		t.Error("Error message should at least contain error type")
	}
}

func BenchmarkNetworkError(b *testing.B) {
	originalErr := errors.New("test error")
	for i := 0; i < b.N; i++ {
		_ = NetworkError("network error", originalErr)
	}
}

func BenchmarkAPIError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = APIError(500, "Internal Server Error", nil)
	}
}

func BenchmarkErrorError(b *testing.B) {
	err := NetworkError("test error", errors.New("original"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func BenchmarkIsRetryable(b *testing.B) {
	err := APIError(500, "test", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.IsRetryable()
	}
}

func BenchmarkIsRetryableError(b *testing.B) {
	err := APIError(500, "test", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsRetryableError(err)
	}
}

func BenchmarkExtractStatusCode(b *testing.B) {
	err := APIError(404, "test", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ExtractStatusCode(err)
	}
}
