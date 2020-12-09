package cancel_test

import (
	"fmt"
	"testing"
	"time"
)

func closeChan(cancelChan chan int) {
	close(cancelChan)
}

func sendExitChan(cancelChan chan int) {
	cancelChan <- 1
}

func isCancelled(cancelChan chan int) bool {
	select {
	case v, ok := <-cancelChan:
		fmt.Printf("isCancelled is True. Values is %d,status is %v\n", v, ok)
		return true
	default:
		fmt.Printf("isCancelled is False\n")
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
	time.Sleep(2 * time.Second)
	closeChan(cancelChan)
	//sendExitChan(cancelChan)
	time.Sleep(1 * time.Second)
}
