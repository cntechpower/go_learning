package b

import (
	"testing"
	"time"
)

func TestUniqueId(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v", time.Now().UnixNano()/1e6-1600000000000)
		t.Logf("%v", time.Now().Unix())
		time.Sleep(1 * time.Millisecond)
	}
}
