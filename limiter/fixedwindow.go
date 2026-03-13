package limiter

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	count       int
	limit       int
	windowSize  time.Duration
	windowStart time.Time
	mu          sync.Mutex
}

func NewFixedWindowCounter(windowSize time.Duration, limit int) *FixedWindowCounter {
	return &FixedWindowCounter{
		limit:       limit,
		windowSize:  windowSize,
		windowStart: time.Now(),
		count:       0,
	}
}

func (fw *FixedWindowCounter) IsAllowed() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()

	// Check if current window is **Expired**
	if now.Sub(fw.windowStart) > fw.windowSize {
		fw.count = 0
		fw.windowStart = time.Now()
	}

	if fw.count < fw.limit {
		fw.count += 1
		return true
	}

	return false
}
