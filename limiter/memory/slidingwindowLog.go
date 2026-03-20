package memory

import (
	"sync"
	"time"
)

type SlidingWindowLog struct {
	limit      int
	windowSize time.Duration
	log        []time.Time
	mu         sync.Mutex
}

func NewSlidingWindowLog(windowSize time.Duration, limit int) *SlidingWindowLog {
	return &SlidingWindowLog{
		limit:      limit,
		windowSize: windowSize,
		log:        make([]time.Time, 0, limit),
	}

}
func (sw *SlidingWindowLog) Allow(id string) (bool, error) {
	return sw.AllowN(id, 1)
}

func (sw *SlidingWindowLog) AllowN(id string, n int) (bool, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	border := now.Add(-sw.windowSize)

	validIndex := 0
	for i, t := range sw.log {
		if t.After(border) {
			validIndex = i
			break
		}
		validIndex++

	}
	if validIndex > 0 {
		sw.log = sw.log[validIndex:]
	}

	if len(sw.log)+n <= sw.limit {
		for range n {
			now := time.Now()
			sw.log = append(sw.log, now)
		}
		return true, nil
	}
	return false, nil

}
