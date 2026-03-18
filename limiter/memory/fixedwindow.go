package memory

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	count      int
	limit      int
	windowSize time.Duration
	start      time.Time
	mu         sync.Mutex
}

func NewFixedWindowCounter(windowSize time.Duration, limit int) (*FixedWindowCounter, error) {
	return &FixedWindowCounter{
		limit:      limit,
		windowSize: windowSize,
		start:      time.Now(),
		count:      0,
	}, nil
}

func (fw *FixedWindowCounter) Allow(id string) bool {
	return fw.AllowN(id, 1)
}
func (fw *FixedWindowCounter) AllowN(id string, n int) bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()

	// Check if current window is **Expired**
	if now.Sub(fw.start) > fw.windowSize {
		fw.count = 0
		fw.start = time.Now()
	}

	if fw.count < fw.limit {
		fw.count += 1
		return true
	}

	return false

}
