package memory

import (
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	id_test := "fake-id"
	seconds := 1 * time.Second
	wantedWindowSize := time.Duration(seconds)
	wantedLimit := 10
	sw := NewSlidingWindowLog(wantedWindowSize, wantedLimit)
	// fill out the window
	for range wantedLimit {
		if ok, _ := sw.Allow(id_test); !ok {
			t.Errorf("the limiter should be able to take requests %+v", sw)
		}
	}

	if ok, _ := sw.Allow(id_test); ok {
		t.Errorf("the limiter should not be able to take requests %+v", sw)
	}

	time.Sleep(seconds)

	if ok, _ := sw.Allow(id_test); !ok {
		t.Errorf("the limiter should be able to take requests %+v", sw)
	}

}

func TestNewSlidingWindow(t *testing.T) {
	wantedWindowSize := time.Duration(1 * time.Second)
	wantedLimit := 10
	sw := NewSlidingWindowLog(wantedWindowSize, wantedLimit)
	if sw.windowSize != wantedWindowSize {
		t.Errorf("Wandted window size")
	}

	if sw.limit != wantedLimit {
		t.Errorf("Wandted max Request per Window: %d, Got: %d", wantedLimit, sw.limit)
	}

}
