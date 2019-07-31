package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) //ALL receiver's ok will be false
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			//ok==ture正常, ok==false表示通道被关闭
			if data, ok := <-ch; ok {
				fmt.Println("Receive data: ", data)
			} else {
				fmt.Println("Channel Closed")
			}
		}
		wg.Done()
	}()
}

func TestProRec(t *testing.T) {
	//wg := sync.WaitGroup{}
	var wg sync.WaitGroup
	intChan := make(chan int, 10)
	wg.Add(1)
	dataProducer(intChan, &wg)
	wg.Add(1)
	dataReceiver(intChan, &wg)
	wg.Add(1)
	dataReceiver(intChan, &wg)
	wg.Wait()
}
