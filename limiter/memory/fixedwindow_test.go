package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestFixedWindow2(t *testing.T) {
	wantedWindowSize := time.Duration(1 * time.Second)
	limit := 10
	fw, _ := NewFixedWindowCounter(wantedWindowSize, limit)
	for range limit {
		if ok, _ := fw.Allow("uid1"); !ok {
			t.Errorf("Should be able to make requests: %+v", fw)
		}
	}
	fmt.Printf("%+v\n", fw)
	if ok, _ := fw.Allow("uid1"); ok {
		t.Errorf("Should not be able to make requests yet: %+v", fw)
	}
	time.Sleep(time.Duration(wantedWindowSize))
	if ok, _ := fw.Allow("uid1"); !ok {
		t.Errorf("Should be able to make requests as a new window should have been opened: %+v", fw)
	}
}

func TestFixedWindow(t *testing.T) {
	wantedWindowSize := time.Duration(1 * time.Second)
	limit := 10
	fw, _ := NewFixedWindowCounter(wantedWindowSize, limit)
	if fw.windowSize != wantedWindowSize {
		t.Errorf("Wandted window size")
	}

	if fw.limit != limit {
		t.Errorf("Wandted max Request per Window: %d, Got: %d", limit, fw.limit)
	}

}
