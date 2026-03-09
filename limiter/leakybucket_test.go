package limiter

import (
	"testing"
	"time"
)

func TestLeakyBucket2(t *testing.T) {
	wantedCap := 5
	wantedFLowRate := 0.5
	lb := NewLeakyBucket(wantedCap, wantedFLowRate)
	for range wantedCap {
		if lb.Take() == false {
			t.Error("Bucket should be available at first")
		}
	}
	// Now its empty
	if lb.Take() == true {
		t.Error("Bucket should be drained by now")
	}
	time.Sleep(2 * time.Second)

	if lb.Take() == false {
		t.Error("Bucket should have enough water")
	}

	if lb.Take() == true {
		t.Error("Bucket should be drained by now")
	}

}

func TestLeakyBucket(t *testing.T) {
	wantedCap := 10
	wantedWater := 0
	wantedFLowRate := 0.5
	lb := NewLeakyBucket(wantedCap, wantedFLowRate)

	if lb.capacity != wantedCap {
		t.Errorf("Wanted cap: %d, Got: %d", wantedCap, lb.capacity)
	}
	if lb.flowRate != lb.flowRate {
		t.Errorf("Wanted cap: %f, Got: %f", wantedFLowRate, lb.flowRate)
	}
	if lb.water != 0 {
		t.Errorf("Wanted cap: %d, Got: %d", wantedWater, lb.water)
	}
}
