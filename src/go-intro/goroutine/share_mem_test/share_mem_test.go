package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounterWithoutProtect(t *testing.T) {
	counter := 0
	for i := 0; i <= 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 50)
	t.Logf("After goroutine, counter is %d", counter)
}

func TestCounterWithProtect(t *testing.T) {
	var mu sync.Mutex
	counter := 0
	for i := 0; i <= 5000; i++ {
		go func() {
			defer func() {
				mu.Unlock()
			}()
			mu.Lock()
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 50)
	t.Logf("After goroutine, counter is %d", counter)
}

func TestCounterWithProtectWaitGroup(t *testing.T) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i <= 5000; i++ {
		go func() {
			wg.Add(1)
			defer func() {
				mu.Unlock()
			}()
			mu.Lock()
			counter++
			wg.Done()
		}()
	}
	// time.Sleep(time.Millisecond * 50)
	wg.Wait()
	t.Logf("After goroutine, counter is %d", counter)
}
