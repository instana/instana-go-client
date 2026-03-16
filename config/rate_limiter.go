package config

import (
	"context"
	"sync"
	"time"
)

// RateLimiter implements a token bucket rate limiter
type RateLimiter struct {
	config       RateLimitConfig
	tokens       float64
	lastRefill   time.Time
	mu           sync.Mutex
	logger       Logger
	refillTicker *time.Ticker
	stopChan     chan struct{}
}

// NewRateLimiter creates a new rate limiter with the given configuration
func NewRateLimiter(config RateLimitConfig, logger Logger) *RateLimiter {
	if logger == nil {
		logger = NewNoOpLogger()
	}

	rl := &RateLimiter{
		config:     config,
		tokens:     float64(config.BurstCapacity),
		lastRefill: time.Now(),
		logger:     logger,
		stopChan:   make(chan struct{}),
	}

	// Start background refill goroutine
	if config.Enabled {
		rl.mu.Lock()
		rl.startRefill()
		rl.mu.Unlock()
	}

	return rl
}

// Wait waits for a token to become available
func (rl *RateLimiter) Wait(ctx context.Context) error {
	if !rl.config.Enabled {
		return nil
	}

	for {
		// Try to acquire a token
		if rl.tryAcquire() {
			return nil
		}

		// If configured to not wait, return error immediately
		if !rl.config.WaitForToken {
			return RateLimitError("rate limit exceeded", 1)
		}

		// Calculate wait time
		waitTime := rl.calculateWaitTime()
		rl.logger.Debug("Rate limit reached, waiting",
			"wait_time", waitTime.String())

		// Wait for token or context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitTime):
			// Continue loop to try again
		}
	}
}

// tryAcquire attempts to acquire a token without blocking
func (rl *RateLimiter) tryAcquire() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Refill tokens based on time elapsed
	rl.refillTokens()

	// Check if we have tokens available
	if rl.tokens >= 1.0 {
		rl.tokens -= 1.0
		return true
	}

	return false
}

// refillTokens adds tokens based on time elapsed since last refill
func (rl *RateLimiter) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)

	// Calculate tokens to add based on elapsed time
	tokensToAdd := elapsed.Seconds() * float64(rl.config.RequestsPerSecond)

	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		rl.lastRefill = now

		// Cap at burst capacity
		if rl.tokens > float64(rl.config.BurstCapacity) {
			rl.tokens = float64(rl.config.BurstCapacity)
		}
	}
}

// calculateWaitTime calculates how long to wait for the next token
func (rl *RateLimiter) calculateWaitTime() time.Duration {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Calculate time needed to get one token
	tokensNeeded := 1.0 - rl.tokens
	if tokensNeeded <= 0 {
		return 0
	}

	secondsNeeded := tokensNeeded / float64(rl.config.RequestsPerSecond)
	return time.Duration(secondsNeeded * float64(time.Second))
}

// startRefill starts the background token refill goroutine
// Note: This method should be called while holding rl.mu lock
func (rl *RateLimiter) startRefill() {
	// Refill every 100ms for smooth rate limiting
	ticker := time.NewTicker(100 * time.Millisecond)
	rl.refillTicker = ticker

	go func() {
		for {
			select {
			case <-ticker.C:
				rl.mu.Lock()
				rl.refillTokens()
				rl.mu.Unlock()
			case <-rl.stopChan:
				return
			}
		}
	}()
}

// Stop stops the rate limiter and cleans up resources
func (rl *RateLimiter) Stop() {
	if rl.refillTicker != nil {
		rl.refillTicker.Stop()
	}
	close(rl.stopChan)
}

// GetAvailableTokens returns the current number of available tokens
func (rl *RateLimiter) GetAvailableTokens() float64 {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refillTokens()
	return rl.tokens
}

// Reset resets the rate limiter to full capacity
func (rl *RateLimiter) Reset() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.tokens = float64(rl.config.BurstCapacity)
	rl.lastRefill = time.Now()
}

// UpdateConfig updates the rate limiter configuration
func (rl *RateLimiter) UpdateConfig(config RateLimitConfig) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Stop existing ticker if disabling
	if !config.Enabled && rl.refillTicker != nil {
		rl.refillTicker.Stop()
		rl.refillTicker = nil
	}

	// Start ticker if enabling
	if config.Enabled && !rl.config.Enabled {
		rl.startRefill()
	}

	rl.config = config

	// Adjust tokens if burst capacity changed
	if rl.tokens > float64(config.BurstCapacity) {
		rl.tokens = float64(config.BurstCapacity)
	}
}

// Made with Bob
