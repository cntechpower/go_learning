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


		select {
		case ret := <-asyncRetCh:
			//
		case <-time.After(5 * time.Second):
			fmt.Println("Operation Timeout")

		}
	}

}
