package time_test

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clearing...")
	}()
	fmt.Println("Starting...")
	panic("err")
}
