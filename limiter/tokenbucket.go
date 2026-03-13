package limiter

import (
	"errors"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity int
	tokens   int
	rate     float64 // persecond
	lastTime time.Time
	mu       sync.Mutex
}

func NewTokenBucket(capacity int, rate float64) (*TokenBucket, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than 0")
	}
	if rate <= 0 {
		return nil, errors.New("rate must be greater than 0")
	}

	return &TokenBucket{
		capacity: capacity,
		tokens:   capacity,
		rate:     rate,
		lastTime: time.Now(),
	}, nil
}

func (tb *TokenBucket) IsAllowed() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()

	elapsed := now.Sub(tb.lastTime).Seconds()
	newTokens := int(elapsed * tb.rate)

	if newTokens > 0 {
		tb.tokens += newTokens
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}

		durationUsed := time.Duration((float64(newTokens) / tb.rate) * float64(time.Second)) // calibrating the time spent generating the new tokens since last time
		tb.lastTime = tb.lastTime.Add(durationUsed)
	}

	if tb.tokens > 0 {
		tb.tokens -= 1
		return true
	}

	return false
}
