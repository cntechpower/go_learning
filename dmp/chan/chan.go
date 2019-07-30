package main

import (
	"fmt"
	"time"
)

func closeTest(exitChan chan struct{}) {
	time.Sleep(time.Millisecond * 5000)
	close(exitChan)
}

func main() {
	exitChan := make(chan struct{}, 0)
	go closeTest(exitChan)
	for {

		select {
		case <-exitChan:
			fmt.Println("exitChan Closed")
		default:
			fmt.Println("Nothing Happend")
		}
		time.Sleep(500 * time.Millisecond)
	}

}
