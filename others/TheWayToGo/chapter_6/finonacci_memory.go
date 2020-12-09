package main

import "fmt"
import "time"

var fibs [41]int

func main() {
	start := time.Now()
	result := 0
	for i := 0; i < 41; i++ {
		result = fibonacci(i)
		fmt.Printf("No %d is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Exec Time is: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
