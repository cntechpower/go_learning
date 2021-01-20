package b

import (
	"math/rand"
	"testing"
	"time"
)

func getRandomValueFromSlice(s []int64) int64 {
	idx := int(rand.NewSource(time.Now().UnixNano()).Int63()) % len(s)
	return s[idx]
}

func TestRand(t *testing.T) {
	s := []int64{5, 6, 7, 8, 9}
	for i := 0; i < 100; i++ {
		t.Logf("got value: %v", getRandomValueFromSlice(s))
	}
}
