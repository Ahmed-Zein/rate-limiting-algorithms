package memory

import (
	"sync"
	"time"
)

type SlidingWindowCounter struct {
	count       int
	limit       int
	weight      float32
	windowSize  time.Duration
	windowStart time.Time
	mu          sync.Mutex
}

func NewSlidingWindowCounter(windowSize time.Duration, limit int) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		weight:      0.75,
		limit:       limit,
		windowSize:  windowSize,
		windowStart: time.Now(),
		count:       0,
	}
}

func (sc *SlidingWindowCounter) Allow(id string) (bool, error) {
	return sc.AllowN(id, 1)

}
func (sc *SlidingWindowCounter) AllowN(id string, n int) (bool, error) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	now := time.Now()

	// Check if current window is **Expired**
	// The sliding window counter(S.W.C) differs from fixed window in count intialization in the new window
	// the S.W.C gives a weight to the last count in order to make a smoother transition between windows
	if now.Sub(sc.windowStart) > sc.windowSize {
		sc.count = int(float32(sc.count) * sc.weight)
		sc.windowStart = time.Now()
	}

	if sc.count+n <= sc.limit {
		sc.count += n
		return true, nil
	}

	return false, nil
}
