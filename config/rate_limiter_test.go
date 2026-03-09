package config

import (
	"context"
	"sync"
	"testing"
	"time"
)

// TestNewRateLimiter tests rate limiter creation
func TestNewRateLimiter(t *testing.T) {
	tests := []struct {
		name           string
		config         RateLimitConfig
		logger         Logger
		expectEnabled  bool
		expectTokens   float64
		expectRefiller bool
	}{
		{
			name: "enabled rate limiter",
			config: RateLimitConfig{
				Enabled:           true,
				RequestsPerSecond: 10,
				BurstCapacity:     20,
				WaitForToken:      true,
			},
			logger:         NewNoOpLogger(),
			expectEnabled:  true,
			expectTokens:   20.0,
			expectRefiller: true,
		},
		{
			name: "disabled rate limiter",
			config: RateLimitConfig{
				Enabled:           false,
				RequestsPerSecond: 10,
				BurstCapacity:     20,
				WaitForToken:      true,
			},
			logger:         NewNoOpLogger(),
			expectEnabled:  false,
			expectTokens:   20.0,
			expectRefiller: false,
		},
		{
			name: "nil logger uses no-op",
			config: RateLimitConfig{
				Enabled:           true,
				RequestsPerSecond: 5,
				BurstCapacity:     10,
				WaitForToken:      false,
			},
			logger:         nil,
			expectEnabled:  true,
			expectTokens:   10.0,
			expectRefiller: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := NewRateLimiter(tt.config, tt.logger)
			defer rl.Stop()

			if rl.config.Enabled != tt.expectEnabled {
				t.Errorf("Expected enabled=%v, got %v", tt.expectEnabled, rl.config.Enabled)
			}

			if rl.tokens != tt.expectTokens {
				t.Errorf("Expected tokens=%v, got %v", tt.expectTokens, rl.tokens)
			}

			if tt.expectRefiller && rl.refillTicker == nil {
				t.Error("Expected refill ticker to be started")
			}

			if !tt.expectRefiller && rl.refillTicker != nil {
				t.Error("Expected refill ticker to not be started")
			}

			if rl.logger == nil {
				t.Error("Logger should never be nil")
			}
		})
	}
}

// TestRateLimiter_Wait_Disabled tests that disabled rate limiter doesn't block
func TestRateLimiter_Wait_Disabled(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           false,
		RequestsPerSecond: 1,
		BurstCapacity:     1,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()

	// Should not block even with many requests
	for i := 0; i < 100; i++ {
		if err := rl.Wait(ctx); err != nil {
			t.Errorf("Wait() should not error when disabled: %v", err)
		}
	}
}

// TestRateLimiter_Wait_ImmediateSuccess tests immediate token acquisition
func TestRateLimiter_Wait_ImmediateSuccess(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()

	// Should succeed immediately for burst capacity
	for i := 0; i < 5; i++ {
		if err := rl.Wait(ctx); err != nil {
			t.Errorf("Wait() failed on request %d: %v", i, err)
		}
	}
}

// TestRateLimiter_Wait_NoWait tests rate limiter with WaitForToken=false
func TestRateLimiter_Wait_NoWait(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 1,
		BurstCapacity:     2,
		WaitForToken:      false,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()

	// First 2 should succeed (burst capacity)
	for i := 0; i < 2; i++ {
		if err := rl.Wait(ctx); err != nil {
			t.Errorf("Wait() failed on request %d: %v", i, err)
		}
	}

	// Third should fail immediately
	err := rl.Wait(ctx)
	if err == nil {
		t.Error("Expected rate limit error, got nil")
	}

	// Check if it's an InstanaError with RateLimit type
	if instanaErr, ok := err.(*InstanaError); ok {
		if instanaErr.Type != ErrorTypeRateLimit {
			t.Errorf("Expected ErrorTypeRateLimit, got %v", instanaErr.Type)
		}
	} else {
		t.Errorf("Expected *InstanaError, got %T: %v", err, err)
	}
}

// TestRateLimiter_Wait_WithWait tests rate limiter with WaitForToken=true
func TestRateLimiter_Wait_WithWait(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10, // 10 requests per second = 100ms per token
		BurstCapacity:     2,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()

	start := time.Now()

	// First 2 should succeed immediately (burst capacity)
	for i := 0; i < 2; i++ {
		if err := rl.Wait(ctx); err != nil {
			t.Errorf("Wait() failed on request %d: %v", i, err)
		}
	}

	// Third should wait for token refill
	if err := rl.Wait(ctx); err != nil {
		t.Errorf("Wait() failed on request 3: %v", err)
	}

	elapsed := time.Since(start)

	// Should have waited at least 80ms (allowing some margin)
	if elapsed < 80*time.Millisecond {
		t.Errorf("Expected to wait at least 80ms, waited %v", elapsed)
	}

	// Should not have waited more than 200ms
	if elapsed > 200*time.Millisecond {
		t.Errorf("Expected to wait less than 200ms, waited %v", elapsed)
	}
}

// TestRateLimiter_Wait_ContextCancellation tests context cancellation
func TestRateLimiter_Wait_ContextCancellation(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 1, // Very slow to ensure we hit cancellation
		BurstCapacity:     1,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	// Exhaust burst capacity
	ctx := context.Background()
	if err := rl.Wait(ctx); err != nil {
		t.Fatalf("Initial Wait() failed: %v", err)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// Should fail with context error
	err := rl.Wait(ctx)
	if err == nil {
		t.Error("Expected context error, got nil")
	}

	if err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

// TestRateLimiter_TryAcquire tests token acquisition
func TestRateLimiter_TryAcquire(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     3,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	// Should succeed for burst capacity
	for i := 0; i < 3; i++ {
		if !rl.tryAcquire() {
			t.Errorf("tryAcquire() failed on attempt %d", i)
		}
	}

	// Should fail when exhausted
	if rl.tryAcquire() {
		t.Error("tryAcquire() should fail when tokens exhausted")
	}
}

// TestRateLimiter_RefillTokens tests token refill mechanism
func TestRateLimiter_RefillTokens(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10, // 10 tokens per second
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	// Exhaust all tokens
	for i := 0; i < 5; i++ {
		rl.tryAcquire()
	}

	// Check tokens are near 0 (allowing for small refill from background goroutine)
	rl.mu.Lock()
	initialTokens := rl.tokens
	rl.mu.Unlock()

	if initialTokens > 0.2 {
		t.Errorf("Expected tokens near 0, got %v", initialTokens)
	}

	// Wait for refill (200ms should give us 2 tokens at 10/sec, with margin for CI)
	time.Sleep(250 * time.Millisecond)

	rl.mu.Lock()
	rl.refillTokens()
	tokens := rl.tokens
	rl.mu.Unlock()

	// Should have refilled at least 0.8 tokens (allowing margin for timing variance)
	if tokens < 0.8 {
		t.Errorf("Expected at least 0.8 tokens after 250ms, got %v", tokens)
	}

	// Should not exceed burst capacity
	if tokens > float64(config.BurstCapacity) {
		t.Errorf("Tokens exceeded burst capacity: %v > %d", tokens, config.BurstCapacity)
	}
}

// TestRateLimiter_CalculateWaitTime tests wait time calculation
func TestRateLimiter_CalculateWaitTime(t *testing.T) {
	tests := []struct {
		name              string
		requestsPerSecond int
		currentTokens     float64
		expectedMin       time.Duration
		expectedMax       time.Duration
	}{
		{
			name:              "no wait needed",
			requestsPerSecond: 10,
			currentTokens:     5.0,
			expectedMin:       0,
			expectedMax:       0,
		},
		{
			name:              "need 1 token at 10/sec",
			requestsPerSecond: 10,
			currentTokens:     0.0,
			expectedMin:       90 * time.Millisecond,
			expectedMax:       110 * time.Millisecond,
		},
		{
			name:              "need 0.5 tokens at 10/sec",
			requestsPerSecond: 10,
			currentTokens:     0.5,
			expectedMin:       40 * time.Millisecond,
			expectedMax:       60 * time.Millisecond,
		},
		{
			name:              "need 1 token at 1/sec",
			requestsPerSecond: 1,
			currentTokens:     0.0,
			expectedMin:       900 * time.Millisecond,
			expectedMax:       1100 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := RateLimitConfig{
				Enabled:           true,
				RequestsPerSecond: tt.requestsPerSecond,
				BurstCapacity:     10,
				WaitForToken:      true,
			}

			rl := NewRateLimiter(config, NewNoOpLogger())
			defer rl.Stop()

			rl.mu.Lock()
			rl.tokens = tt.currentTokens
			rl.mu.Unlock()

			waitTime := rl.calculateWaitTime()

			if waitTime < tt.expectedMin || waitTime > tt.expectedMax {
				t.Errorf("Expected wait time between %v and %v, got %v",
					tt.expectedMin, tt.expectedMax, waitTime)
			}
		})
	}
}

// TestRateLimiter_GetAvailableTokens tests token count retrieval
func TestRateLimiter_GetAvailableTokens(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	// Initial tokens should equal burst capacity
	tokens := rl.GetAvailableTokens()
	if tokens < 4.9 || tokens > 5.1 {
		t.Errorf("Expected ~5 tokens initially, got %v", tokens)
	}

	// Acquire some tokens
	rl.tryAcquire()
	rl.tryAcquire()

	tokens = rl.GetAvailableTokens()
	// Allow for small refill from background goroutine
	if tokens < 2.9 || tokens > 3.1 {
		t.Errorf("Expected ~3 tokens after 2 acquisitions, got %v", tokens)
	}
}

// TestRateLimiter_Reset tests rate limiter reset
func TestRateLimiter_Reset(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	// Exhaust tokens
	for i := 0; i < 5; i++ {
		rl.tryAcquire()
	}

	// Check tokens are near 0 (allowing for small refill)
	if rl.tokens > 0.1 {
		t.Errorf("Expected tokens near 0, got %v", rl.tokens)
	}

	// Reset
	rl.Reset()

	// Check tokens are reset to burst capacity (allowing small margin)
	if rl.tokens < 4.9 || rl.tokens > 5.1 {
		t.Errorf("Expected ~5 tokens after reset, got %v", rl.tokens)
	}
}

// TestRateLimiter_UpdateConfig tests configuration updates
func TestRateLimiter_UpdateConfig(t *testing.T) {
	initialConfig := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(initialConfig, NewNoOpLogger())
	defer rl.Stop()

	// Verify initial state
	if rl.refillTicker == nil {
		t.Error("Expected refill ticker to be running")
	}

	// Update to disabled
	disabledConfig := RateLimitConfig{
		Enabled:           false,
		RequestsPerSecond: 10,
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl.UpdateConfig(disabledConfig)

	if rl.config.Enabled {
		t.Error("Expected rate limiter to be disabled")
	}

	// Update to enabled with different capacity
	enabledConfig := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 20,
		BurstCapacity:     3,
		WaitForToken:      false,
	}

	rl.UpdateConfig(enabledConfig)

	if !rl.config.Enabled {
		t.Error("Expected rate limiter to be enabled")
	}

	if rl.config.RequestsPerSecond != 20 {
		t.Errorf("Expected 20 requests/sec, got %d", rl.config.RequestsPerSecond)
	}

	// Tokens should be capped at new burst capacity
	if rl.tokens > 3.0 {
		t.Errorf("Expected tokens <= 3, got %v", rl.tokens)
	}
}

// TestRateLimiter_Stop tests cleanup
func TestRateLimiter_Stop(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     5,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())

	if rl.refillTicker == nil {
		t.Error("Expected refill ticker to be running")
	}

	rl.Stop()

	// Give goroutine time to stop
	time.Sleep(50 * time.Millisecond)

	// Channel should be closed
	select {
	case <-rl.stopChan:
		// Expected - channel is closed
	default:
		t.Error("Expected stop channel to be closed")
	}
}

// TestRateLimiter_ConcurrentAccess tests thread safety
func TestRateLimiter_ConcurrentAccess(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 100,
		BurstCapacity:     50,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()
	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	// Launch 100 concurrent requests
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := rl.Wait(ctx); err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// All requests should eventually succeed
	if successCount != 100 {
		t.Errorf("Expected 100 successful requests, got %d", successCount)
	}
}

// TestRateLimiter_BurstThenSteady tests burst followed by steady rate
func TestRateLimiter_BurstThenSteady(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10, // 10 requests per second
		BurstCapacity:     5,  // Allow burst of 5
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()
	start := time.Now()

	// Make 10 requests (5 burst + 5 steady)
	for i := 0; i < 10; i++ {
		if err := rl.Wait(ctx); err != nil {
			t.Errorf("Request %d failed: %v", i, err)
		}
	}

	elapsed := time.Since(start)

	// First 5 should be immediate (burst)
	// Next 5 should take ~500ms at 10/sec
	// Total should be around 500ms (allowing margin)
	if elapsed < 400*time.Millisecond {
		t.Errorf("Completed too quickly: %v", elapsed)
	}

	if elapsed > 700*time.Millisecond {
		t.Errorf("Took too long: %v", elapsed)
	}
}

// TestRateLimiter_TokenRefillAccuracy tests refill accuracy
func TestRateLimiter_TokenRefillAccuracy(t *testing.T) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10, // 10 tokens per second
		BurstCapacity:     10,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	// Exhaust all tokens
	for i := 0; i < 10; i++ {
		rl.tryAcquire()
	}

	// Wait for 500ms (should refill 5 tokens at 10/sec)
	time.Sleep(500 * time.Millisecond)

	tokens := rl.GetAvailableTokens()

	// Should have approximately 5 tokens (allowing 1 token margin)
	if tokens < 4.0 || tokens > 6.0 {
		t.Errorf("Expected ~5 tokens after 500ms, got %v", tokens)
	}
}

// Benchmark tests

// BenchmarkRateLimiter_Wait benchmarks Wait() performance
func BenchmarkRateLimiter_Wait(b *testing.B) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 1000000, // Very high to avoid blocking
		BurstCapacity:     1000000,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rl.Wait(ctx)
	}
}

// BenchmarkRateLimiter_TryAcquire benchmarks tryAcquire() performance
func BenchmarkRateLimiter_TryAcquire(b *testing.B) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 1000000,
		BurstCapacity:     1000000,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rl.tryAcquire()
	}
}

// BenchmarkRateLimiter_GetAvailableTokens benchmarks token retrieval
func BenchmarkRateLimiter_GetAvailableTokens(b *testing.B) {
	config := RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     10,
		WaitForToken:      true,
	}

	rl := NewRateLimiter(config, NewNoOpLogger())
	defer rl.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rl.GetAvailableTokens()
	}
}
