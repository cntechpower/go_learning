package main

import "fmt"
import "time"

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	//fmt.Println(<-ch1)
	go pop(ch1)
	time.Sleep(1e9)
}
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func pop(ch chan int) {
	var input int
	for {
		input = <-ch
		fmt.Printf("%d\n", input)
	}
}
