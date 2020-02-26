package select_timeout_test

import (
	"testing"
	"time"
)

func NotifyAferSeconds(i int, retCh chan string) {
	time.Sleep(time.Duration(i) * time.Second)
	retCh <- "MESSAGE"

}

func TestSelectTimeout(t *testing.T) {
	retCh := make(chan string, 1)
	go NotifyAferSeconds(1, retCh)
	select {
	case ret := <-retCh:
		t.Logf("Receive Message: %v", ret)
	case <-time.After(time.Second * 10):
		t.Logf("Receive Timeout")
		// default:
		// 	t.Logf("No Messgae")
	}
}
