package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	//num := 10
	//chanNum := make(chan struct{}, num)
	for i := uint32(0); i < 10; i++ {
		// go func() {
		// 	fmt.Println(i)
		// 	chanNum <- struct{}{}

		// }()

		// go func(i int) {
		// 	fmt.Println(i)
		// 	chanNum <- struct{}{}

		// }(i)

		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)

		}(i)
	}
	//time.Sleep(time.Millisecond * 500)

	// for i := 0; i < num; i++ {
	// 	<-chanNum

	// }
	trigger(10, func() {})
}
