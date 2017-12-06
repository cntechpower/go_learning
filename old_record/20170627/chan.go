// chan
package main

import (
	"fmt"
	"time"
)

func child() {
	fmt.Println("Hello World!")
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Printf("I am child")
	}()
	time.Sleep(1 * time.Second)
}

func deadlock() {
	chanrw := make(chan int)
	chanrw <- 1
	go func() {
		<-chanrw
		fmt.Printf("go func")
	}()
	fmt.Printf("main func")
}

func buffer_chan() {
	chanrw := make(chan int, 1)
	chanrw <- 88
	go func() {
		chan_value := <-chanrw
		fmt.Printf("go func,chan value is: %v\n", chan_value)
	}()
	time.Sleep(1 * time.Second)
	fmt.Printf("main func\n")
}

func waitchildbyblock() {
	chanrw := make(chan int)
	go func() {
		chan_value := <-chanrw
		fmt.Printf("chan value is %v\n", chan_value)
	}()
	chanrw <- 1
	fmt.Printf("main func\n")
}
func producer(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Printf("Produce:%v\n", i)

	}
}
func customer(p <-chan int) {
	for i := 0; i < 10; i++ {
		j := <-p
		fmt.Printf("Customer:%v\n", j)
	}
}

func pcnobuffer() {
	chanrw := make(chan int)
	go producer(chanrw)
	go customer(chanrw)
	time.Sleep(1 * time.Second)
}

func pcbuffer() {
	chanrw := make(chan int, 10)
	go producer(chanrw)
	go customer(chanrw)
	time.Sleep(10 * time.Second)
}
