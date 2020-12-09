package main

import (
	"fmt"
	"time"
)

func doWithStop(closeChan chan struct{}) {
	for {
		select {
		case <-closeChan:
			return
		default:
			fmt.Println("doWithStop: " + time.Now().Format(time.RFC3339))
			time.Sleep(1 * time.Second)
		}
	}
}

func testWithStop(closeChan chan struct{}) {
	innerCloseChan := make(chan struct{}, 0)
	defer close(innerCloseChan)
	go doWithStop(innerCloseChan)
	select {
	case <-closeChan:
		return
	}
}

func doWithoutStop() {
	for {
		fmt.Println("doWithoutStop: " + time.Now().Format(time.RFC3339))
		time.Sleep(1 * time.Second)
	}
}

func testWithoutStop(closeChan chan struct{}) {
	go doWithoutStop()
	select {
	case <-closeChan:
		return
	}
}

func main() {
	fmt.Println("starting testWithStop...")
	closeChan := make(chan struct{}, 0)
	go testWithStop(closeChan)
	time.Sleep(10 * time.Second)
	fmt.Println("stopping testWithStop...")
	close(closeChan) //inner goroutine will stop

	fmt.Println("starting testWithoutStop...")
	closeChan2 := make(chan struct{}, 0)
	go testWithoutStop(closeChan2)
	time.Sleep(10 * time.Second)
	fmt.Println("stopping testWithoutStop...")
	close(closeChan2) //inner goroutine won't stop
	time.Sleep(10 * time.Second)
}
