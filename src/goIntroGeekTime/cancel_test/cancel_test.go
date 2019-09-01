package cancel_test

import (
	"fmt"
	"testing"
	"time"
)

func cancel_1(cancelChan chan int) {
	close(cancelChan)
}

func cancel_2(cancelChan chan int) {
	cancelChan <- 1
}

func isCancelled(cancelChan chan int) bool {
	select {
	case v, ok := <-cancelChan:
		fmt.Printf("isCancelled is True. Values is %d,status is %v\n", v, ok)
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan int, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan int) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			t.Logf("%d goroutine is stopped", i)
		}(i, cancelChan)
	}
	cancel_1(cancelChan)
	//cancel_2(cancelChan)
	time.Sleep(1 * time.Second)
}
