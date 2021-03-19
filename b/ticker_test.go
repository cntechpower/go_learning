package b

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	for range time.NewTicker(5 * time.Second).C {
		t.Logf("start at %v", time.Now().Format(time.RFC3339))
		time.Sleep(time.Second * 3)
		t.Logf("end at %v", time.Now().Format(time.RFC3339))
	}
}
