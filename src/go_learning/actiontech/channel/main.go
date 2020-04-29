package main

import (
	"fmt"
	"time"
)

func do() {
	for {
		fmt.Println(time.Now().Format(time.RFC3339))
		time.Sleep(1 * time.Second)
	}

}

func test(closeChan chan struct{}) {
	go do()
	select {
	case <-closeChan:
		return
	}
}

func main() {
	fmt.Println("starting test...")
	closeChan := make(chan struct{}, 0)
	go test(closeChan)
	time.Sleep(10 * time.Second)
	fmt.Println("stopping test...")
	close(closeChan)
	time.Sleep(10 * time.Second)
}
