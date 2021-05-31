package agent

import (
	"testing"
	"time"
)

func TestAgent_sleep(t *testing.T) {
	d := time.Second
	ag := &agent{stopSleepCh: make(chan struct{})}

	t0 := time.Now()
	ag.sleep(d)
	if time.Since(t0) < d {
		t.Fatalf("agent sleep few than %v", d)
	}

	close(ag.stopSleepCh)
	t1 := time.Now()
	ag.sleep(d)
	if time.Since(t1) >= d {
		t.Fatalf("agent sleep more than %v", d)
	}
}
