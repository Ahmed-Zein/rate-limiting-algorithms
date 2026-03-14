package memory

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	flowRate float64 // per_second
	water    int
	lastTime time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, flowRate float64) *LeakyBucket {
	return &LeakyBucket{
		water:    0,
		capacity: capacity,
		flowRate: flowRate,
		lastTime: time.Now(),
	}

}

func (lb *LeakyBucket) IsAllowed() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	elapsed := time.Since(lb.lastTime).Seconds()
	leakedWater := elapsed * lb.flowRate

	if leakedWater > 0 {
		lb.water -= int(leakedWater)
		if lb.water < 0 {
			lb.water = 0
		}
		lb.lastTime = time.Now()
	}

	if lb.water < lb.capacity {
		lb.water++
		return true
	}

	return false
}
