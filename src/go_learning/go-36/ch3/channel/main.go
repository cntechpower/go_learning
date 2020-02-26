package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// channels := [3]chan int{
	// 	make(chan int, 1),
	// 	make(chan int, 1),
	// 	make(chan int, 1),
	// }
	// index := rand.Intn(3)
	// fmt.Printf("Select Index: %v\n", index)
	// //channels[index] <- index
	// select {
	// case <-channels[0]:
	// 	fmt.Println("channel 0 selected")
	// case <-channels[1]:
	// 	fmt.Println("channel 1 selected")
	// case <-channels[2]:
	// 	fmt.Println("channel 2 selected")
	// default:
	// 	fmt.Println("no channel selected")
	// }

	intChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() { close(intChan) }) //不会阻塞在这里,会继续进行执行
	intChan <- 2
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				fmt.Printf("channel closed,value is %v\n", v)
				//break
				os.Exit(1)
			} else {
				fmt.Printf("channel open,value is %v\n", v)
			}
		}
	}

}
