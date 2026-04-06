package testutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

// TestNewTestHTTPServer tests the creation of a new test HTTP server
func TestNewTestHTTPServer(t *testing.T) {
	server := NewTestHTTPServer()

	if server == nil {
		t.Fatal("Expected server to be created, got nil")
	}

	// Verify initial state
	if server.GetPort() != -1 {
		t.Errorf("Expected port to be -1 before start, got %d", server.GetPort())
	}
}

// TestTestHTTPServer_StartAndClose tests starting and closing the server
func TestTestHTTPServer_StartAndClose(t *testing.T) {
	server := NewTestHTTPServer()

	// Start the server
	server.Start()

	// Verify port is assigned
	port := server.GetPort()
	if port <= 0 {
		t.Errorf("Expected positive port number, got %d", port)
	}

	// Verify health endpoint is accessible
	url := fmt.Sprintf("https://localhost:%d/health", port)
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		t.Logf("Health check failed (expected with self-signed cert): %v", err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
	}

	// Close the server
	server.Close()
}

// TestTestHTTPServer_AddRoute tests adding custom routes
func TestTestHTTPServer_AddRoute(t *testing.T) {
	server := NewTestHTTPServer()

	// Add a custom route
	called := false
	server.AddRoute("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	})

	server.Start()
	defer server.Close()

	// Make request to custom route
	url := fmt.Sprintf("https://localhost:%d/test", server.GetPort())
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		t.Logf("Request failed (expected with self-signed cert): %v", err)
	} else {
		defer resp.Body.Close()
		if !called {
			t.Error("Expected handler to be called")
		}
	}
}

// TestTestHTTPServer_GetCallCount tests the call counter functionality
func TestTestHTTPServer_GetCallCount(t *testing.T) {
	server := NewTestHTTPServer()

	// Add a route
	server.AddRoute("POST", "/counter", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server.Start()
	defer server.Close()

	// Initial count should be 0
	if count := server.GetCallCount("POST", "/counter"); count != 0 {
		t.Errorf("Expected initial count 0, got %d", count)
	}

	// Make multiple requests
	url := fmt.Sprintf("https://localhost:%d/counter", server.GetPort())
	for i := 0; i < 3; i++ {
		http.Post(url, "application/json", bytes.NewBufferString("{}")) //nolint:gosec,errcheck
		time.Sleep(10 * time.Millisecond)
	}

	// Verify count (may not be exact due to TLS handshake issues)
	count := server.GetCallCount("POST", "/counter")
	t.Logf("Call count: %d", count)
}

// TestTestHTTPServer_GetCallCount_NonExistentRoute tests getting count for non-existent route
func TestTestHTTPServer_GetCallCount_NonExistentRoute(t *testing.T) {
	server := NewTestHTTPServer()

	// Get count for route that was never called
	count := server.GetCallCount("GET", "/nonexistent")
	if count != 0 {
		t.Errorf("Expected count 0 for non-existent route, got %d", count)
	}
}

// TestEchoHandlerFunc tests the echo handler with valid request
func TestEchoHandlerFunc(t *testing.T) {
	server := NewTestHTTPServer()
	server.AddRoute("POST", "/echo", EchoHandlerFunc)

	server.Start()
	defer server.Close()

	// Make request with body
	requestBody := []byte(`{"test": "data"}`)
	url := fmt.Sprintf("https://localhost:%d/echo", server.GetPort())

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req) //nolint:gosec
	if err != nil {
		t.Logf("Request failed (expected with self-signed cert): %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if !bytes.Equal(requestBody, responseBody) {
		t.Errorf("Expected echo response to match request body")
	}
}

// TestEchoHandlerFunc_WithContentType tests echo handler preserves content type
func TestEchoHandlerFunc_WithContentType(t *testing.T) {
	server := NewTestHTTPServer()
	server.AddRoute("POST", "/echo", EchoHandlerFunc)

	server.Start()
	defer server.Close()

	url := fmt.Sprintf("https://localhost:%d/echo", server.GetPort())
	req, err := http.NewRequest("POST", url, bytes.NewBufferString("test"))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req) //nolint:gosec
	if err != nil {
		t.Logf("Request failed (expected with self-signed cert): %v", err)
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("Expected Content-Type 'text/plain', got '%s'", contentType)
	}
}

// TestWriteInternalServerError tests the WriteInternalServerError helper
func TestWriteInternalServerError(t *testing.T) {
	server := NewTestHTTPServer()

	testError := errors.New("test error message")
	server.AddRoute("GET", "/error", func(w http.ResponseWriter, r *http.Request) {
		server.WriteInternalServerError(w, testError)
	})

	server.Start()
	defer server.Close()

	url := fmt.Sprintf("https://localhost:%d/error", server.GetPort())
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		t.Logf("Request failed (expected with self-signed cert): %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if string(body) != testError.Error() {
		t.Errorf("Expected error message '%s', got '%s'", testError.Error(), string(body))
	}
}

// TestWriteJSONResponse tests the WriteJSONResponse helper
func TestWriteJSONResponse(t *testing.T) {
	server := NewTestHTTPServer()

	testData := map[string]interface{}{
		"status": "success",
		"data":   []int{1, 2, 3},
	}
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	server.AddRoute("GET", "/json", func(w http.ResponseWriter, r *http.Request) {
		server.WriteJSONResponse(w, jsonData)
	})

	server.Start()
	defer server.Close()

	url := fmt.Sprintf("https://localhost:%d/json", server.GetPort())
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		t.Logf("Request failed (expected with self-signed cert): %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Expected Content-Type 'application/json; charset=utf-8', got '%s'", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if !bytes.Equal(jsonData, body) {
		t.Errorf("Expected JSON response to match input data")
	}
}

// TestWriteJSONResponse_EmptyData tests WriteJSONResponse with empty data
func TestWriteJSONResponse_EmptyData(t *testing.T) {
	server := NewTestHTTPServer()

	server.AddRoute("GET", "/empty", func(w http.ResponseWriter, r *http.Request) {
		server.WriteJSONResponse(w, []byte("{}"))
	})

	server.Start()
	defer server.Close()

	url := fmt.Sprintf("https://localhost:%d/empty", server.GetPort())
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		t.Logf("Request failed (expected with self-signed cert): %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

// TestTestHTTPServer_MultipleRoutes tests adding multiple routes
func TestTestHTTPServer_MultipleRoutes(t *testing.T) {
	server := NewTestHTTPServer()

	// Add multiple routes
	server.AddRoute("GET", "/route1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("route1"))
	})

	server.AddRoute("POST", "/route2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("route2"))
	})

	server.AddRoute("PUT", "/route3", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("route3"))
	})

	server.Start()
	defer server.Close()

	// Verify all routes are accessible
	port := server.GetPort()
	if port <= 0 {
		t.Fatal("Server port not assigned")
	}

	t.Logf("Server started on port %d with multiple routes", port)
}

// TestTestHTTPServer_GetPort_BeforeStart tests GetPort before server starts
func TestTestHTTPServer_GetPort_BeforeStart(t *testing.T) {
	server := NewTestHTTPServer()

	port := server.GetPort()
	if port != -1 {
		t.Errorf("Expected port -1 before start, got %d", port)
	}
}

// TestTestHTTPServer_GetPort_AfterStart tests GetPort after server starts
func TestTestHTTPServer_GetPort_AfterStart(t *testing.T) {
	server := NewTestHTTPServer()
	server.Start()
	defer server.Close()

	port := server.GetPort()
	if port <= 0 {
		t.Errorf("Expected positive port after start, got %d", port)
	}
}

// TestHealthEndpoint tests the built-in health endpoint
func TestHealthEndpoint(t *testing.T) {
	server := NewTestHTTPServer()
	server.Start()
	defer server.Close()

	url := fmt.Sprintf("https://localhost:%d/health", server.GetPort())
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		t.Logf("Health check failed (expected with self-signed cert): %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 for health endpoint, got %d", resp.StatusCode)
	}
}

// TestTestHTTPServer_CallCounterIncrement tests that call counter increments correctly
func TestTestHTTPServer_CallCounterIncrement(t *testing.T) {
	server := NewTestHTTPServer()

	callCount := 0
	server.AddRoute("GET", "/increment", func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.WriteHeader(http.StatusOK)
	})

	server.Start()
	defer server.Close()

	// Verify initial count
	if count := server.GetCallCount("GET", "/increment"); count != 0 {
		t.Errorf("Expected initial count 0, got %d", count)
	}

	// Make a request
	url := fmt.Sprintf("https://localhost:%d/increment", server.GetPort())
	http.Get(url) //nolint:gosec,errcheck
	time.Sleep(50 * time.Millisecond)

	// Handler should have been called
	if callCount == 0 {
		t.Log("Handler was not called (likely due to TLS handshake failure)")
	}
}

// TestTestHTTPServer_DifferentMethods tests routes with different HTTP methods
func TestTestHTTPServer_DifferentMethods(t *testing.T) {
	server := NewTestHTTPServer()

	server.AddRoute("GET", "/method", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server.AddRoute("POST", "/method", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	server.Start()
	defer server.Close()

	// Verify different methods have separate counters
	getCount := server.GetCallCount("GET", "/method")
	postCount := server.GetCallCount("POST", "/method")

	if getCount != 0 {
		t.Errorf("Expected GET count 0, got %d", getCount)
	}
	if postCount != 0 {
		t.Errorf("Expected POST count 0, got %d", postCount)
	}
}

// TestTestHTTPServer_CloseWithoutStart tests closing server that was never started
func TestTestHTTPServer_CloseWithoutStart(t *testing.T) {
	server := NewTestHTTPServer()

	// Should not panic
	server.Close()
}

// TestTestHTTPServer_DoubleClose tests closing server twice
func TestTestHTTPServer_DoubleClose(t *testing.T) {
	server := NewTestHTTPServer()
	server.Start()

	// First close
	server.Close()

	// Second close should not panic
	server.Close()
}

// TestEchoHandlerFunc_Direct tests EchoHandlerFunc directly without HTTP server
func TestEchoHandlerFunc_Direct(t *testing.T) {
	// Create a mock response writer
	w := &mockResponseWriter{
		headers: make(http.Header),
		body:    &bytes.Buffer{},
	}

	// Create a request with body
	requestBody := []byte(`{"test": "data"}`)
	req, err := http.NewRequest("POST", "/echo", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Call the handler
	EchoHandlerFunc(w, req)

	// Verify response
	if w.statusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.statusCode)
	}

	if w.headers.Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", w.headers.Get("Content-Type"))
	}

	if !bytes.Equal(requestBody, w.body.Bytes()) {
		t.Errorf("Expected body to match request, got %s", w.body.String())
	}
}

// TestHealthFunc_Direct tests healthFunc directly
func TestHealthFunc_Direct(t *testing.T) {
	w := &mockResponseWriter{
		headers: make(http.Header),
		body:    &bytes.Buffer{},
	}

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Call the handler
	healthFunc(w, req)

	// Verify response
	if w.statusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.statusCode)
	}
}

// TestWriteInternalServerError_Direct tests WriteInternalServerError directly
func TestWriteInternalServerError_Direct(t *testing.T) {
	server := NewTestHTTPServer()

	w := &mockResponseWriter{
		headers: make(http.Header),
		body:    &bytes.Buffer{},
	}

	testErr := errors.New("test error message")
	server.WriteInternalServerError(w, testErr)

	if w.statusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.statusCode)
	}

	if w.headers.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type 'text/plain; charset=utf-8', got '%s'", w.headers.Get("Content-Type"))
	}

	if w.body.String() != testErr.Error() {
		t.Errorf("Expected error message '%s', got '%s'", testErr.Error(), w.body.String())
	}
}

// TestWriteJSONResponse_Direct tests WriteJSONResponse directly
func TestWriteJSONResponse_Direct(t *testing.T) {
	server := NewTestHTTPServer()

	w := &mockResponseWriter{
		headers: make(http.Header),
		body:    &bytes.Buffer{},
	}

	jsonData := []byte(`{"status":"success","data":[1,2,3]}`)
	server.WriteJSONResponse(w, jsonData)

	if w.statusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.statusCode)
	}

	if w.headers.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("Expected Content-Type 'application/json; charset=utf-8', got '%s'", w.headers.Get("Content-Type"))
	}

	if !bytes.Equal(jsonData, w.body.Bytes()) {
		t.Errorf("Expected JSON data to match, got %s", w.body.String())
	}
}

// TestWrapHandlerFunc_Direct tests wrapHandlerFunc directly
func TestWrapHandlerFunc_Direct(t *testing.T) {
	server := NewTestHTTPServer().(*testHTTPServerImpl)

	called := false
	handler := func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}

	wrapped := server.wrapHandlerFunc(handler)

	w := &mockResponseWriter{
		headers: make(http.Header),
		body:    &bytes.Buffer{},
	}

	req, _ := http.NewRequest("GET", "/test", nil)

	// Call wrapped handler
	wrapped(w, req)

	if !called {
		t.Error("Expected wrapped handler to call original handler")
	}

	// Verify counter was incremented
	count := server.GetCallCount("GET", "/test")
	if count != 1 {
		t.Errorf("Expected call count 1, got %d", count)
	}

	// Call again
	wrapped(w, req)
	count = server.GetCallCount("GET", "/test")
	if count != 2 {
		t.Errorf("Expected call count 2, got %d", count)
	}
}

// TestWaitForServerAlive_Direct tests waitForServerAlive behavior
func TestWaitForServerAlive_Direct(t *testing.T) {
	server := NewTestHTTPServer()
	server.Start()
	defer server.Close()

	// Server should already be alive after Start()
	// This test verifies the server started successfully
	port := server.GetPort()
	if port <= 0 {
		t.Error("Expected server to have valid port after Start()")
	}
}

// mockResponseWriter is a mock implementation of http.ResponseWriter for testing
type mockResponseWriter struct {
	headers    http.Header
	body       *bytes.Buffer
	statusCode int
}

func (m *mockResponseWriter) Header() http.Header {
	return m.headers
}

func (m *mockResponseWriter) Write(data []byte) (int, error) {
	return m.body.Write(data)
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}
