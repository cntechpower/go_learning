package loop_test

import "testing"

func TestLoop(t *testing.T) {
	n := 0
	for n < 10 {
		t.Log(n)
		n++
	}
}
