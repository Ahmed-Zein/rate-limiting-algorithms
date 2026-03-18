package memory

import (
	"testing"
	"time"

	"github.com/ahmed-zein/go_rate_limiting/config"
)

var cfg *config.BucketBasedConfig = &config.BucketBasedConfig{
	Capacity: 10,
	Rate:     1.,
}

func TestFlowRate2(t *testing.T) {
	bucket, err := NewTokenBucket(cfg)
	if err != nil {
		t.Error(err)
	}
	count := 0
	expected := 2
	incrCount := func() { count++ }

	// Drain the bucket
	for range 10 {
		if bucket.IsAllowed() {
		}
	}
	time.Sleep(2 * time.Second)

	for range 10 {
		if bucket.IsAllowed() {
			incrCount()
		}
	}
	if count != expected {
		t.Errorf("Expected count: %d, Got: %d", expected, count)
	}

}

func TestFlowRate(t *testing.T) {
	bucket, err := NewTokenBucket(cfg)
	if err != nil {
		t.Error(err)
	}
	count := 0
	incrCount := func() { count++ }
	for range cfg.Capacity {
		if bucket.IsAllowed() {
			incrCount()
		}
	}
	if count != cfg.Capacity {
		t.Errorf("Expected count: %d, Got: %d", 5, count)
	}

}

func TestNewBucket(t *testing.T) {
	wandtedCap := cfg.Capacity
	wantedFillRate := cfg.Rate
	bucket, err := NewTokenBucket(cfg)
	if err != nil {
		t.Error(err)
	}
	if bucket.capacity != wandtedCap {
		t.Errorf("Wanted a bucket with capacity: %d got %d", wandtedCap, bucket.capacity)
	}
	if bucket.tokens != wandtedCap {
		t.Errorf("Wanted a bucket with initial size of: %d got %d", wandtedCap, bucket.tokens)
	}
	if bucket.rate != wantedFillRate {
		t.Errorf("Wanted a bucket with fill rate of %d got %f", wandtedCap, bucket.rate)
	}
}
