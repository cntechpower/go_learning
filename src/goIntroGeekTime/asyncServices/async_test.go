package asyncServices

import (
	"fmt"
	"testing"
	"time"
)

func asyncService() chan string {
	retCh := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		retCh <- "return something"
	}()

	return retCh
}

func TestAsyncServices(t *testing.T) {
	asyncRetCh := asyncService()
	timeoutTimer := time.After(5 * time.Second)
MAIN:
	for {
		select {
		case ret := <-asyncRetCh:
			fmt.Println(ret)
			break MAIN
		//case <-time.After(5 * time.Second):
		case <-timeoutTimer:
			fmt.Println("Operation Timeout")
			break MAIN
		default:
			time.Sleep(1 * time.Second) //time.After不能和default连用,不然每次都会命中default, 从而导致每次进入select都会重新创建timer
		}
	}

}
