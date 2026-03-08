package limiter

import (
	"sync"
	"time"
)

// we need a bucket-capacity; bucket-size; time2generate; lastgenTime;
type TokenBucket struct {
	capacity int
	tokens   int
	fillRate int // persecond
	lastTime time.Time
	mu       sync.Mutex
}

func NewTokenBucket(capacity int, fillRate int) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		tokens:   capacity,
		fillRate: fillRate,
		lastTime: time.Now(),
	}
}

func (t *TokenBucket) Empty() bool {
	return t.tokens < 1
}

func (t *TokenBucket) Take() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	now := time.Now()

	elapsed := now.Sub(t.lastTime).Seconds()
	newTokens := int(elapsed) * t.fillRate
	if newTokens > 0 {
		t.tokens = min(t.tokens+newTokens, t.capacity)
		t.lastTime = time.Now()
	}
	if t.tokens > 0 {
		t.tokens -= 1
		return true
	}

	return false
}
