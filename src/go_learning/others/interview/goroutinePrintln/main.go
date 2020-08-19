package main

import (
	"fmt"
	"sync"
)

var s = []string{"a", "b", "c", "d", "e"}
var ch1 chan struct{}
var ch2 chan struct{}
var wg sync.WaitGroup

func g1() {
	for i := 0; i < len(s); i = i + 2 {
		<-ch1
		fmt.Printf("g1: %v\n", s[i])
		ch2 <- struct{}{}
		wg.Done()
	}
}

func g2() {
	for i := 1; i < len(s); i = i + 2 {
		<-ch2
		fmt.Printf("g2: %v\n", s[i])
		ch1 <- struct{}{}
		wg.Done()
	}
}

func main() {
	ch1 = make(chan struct{}, 1)
	ch2 = make(chan struct{}, 1)
	wg.Add(len(s))
	ch1 <- struct{}{}
	go g1()
	go g2()
	wg.Wait()

}
