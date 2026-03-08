package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestFlowRate2(t *testing.T) {
	bucket := NewTokenBucket(5, 1)
	count := 0
	expected := 2
	incrCount := func() { count++ }

	// Drain the bucket
	for range 10 {
		if bucket.Take() {
		}
	}
	time.Sleep(2 * time.Second)

	for range 10 {
		if bucket.Take() {
			incrCount()
		}
	}
	if count != expected {
		t.Errorf("Expected count: %d, Got: %d", expected, count)
	}

}

func TestFlowRate(t *testing.T) {
	bucket := NewTokenBucket(5, 1)
	count := 0
	incrCount := func() { count++ }
	for range 10 {
		if bucket.Take() {
			incrCount()
		}
		fmt.Printf("%+v\n", bucket)
	}
	if count != 5 {
		t.Errorf("Expected count: %d, Got: %d", 5, count)
	}

}

func TestNewBucket(t *testing.T) {
	wandtedCap := 10
	wantedFillRate := 1
	bucket := NewTokenBucket(wandtedCap, wantedFillRate)
	if bucket.capacity != wandtedCap {
		t.Errorf("Wanted a bucket with capacity: %d got %d", wandtedCap, bucket.capacity)
	}
	if bucket.tokens != wandtedCap {
		t.Errorf("Wanted a bucket with initial size of: %d got %d", wandtedCap, bucket.tokens)
	}
	if bucket.fillRate != wantedFillRate {
		t.Errorf("Wanted a bucket with fill rate of %d got %d", wandtedCap, bucket.fillRate)
	}
}
