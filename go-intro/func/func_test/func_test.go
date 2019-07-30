package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	//返回两个随机数,一个不大于10,一个不大于20
	return rand.Intn(10), rand.Intn(20)
}

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
		t.Log("Defer Called 1")
	}()
	defer func() {
		t.Log("Defer Called 2")
	}()
	defer func() {
		t.Log("Defer Called 3")
	}()
	t.Log("Started")
	panic("Fatal error")
}

func slowFunc(op int) int {
	time.Sleep(1 * time.Second)
	return op * op
}

func sumMulti(ops ...int) int {
	s := 0
	for _, op := range ops {
		s = s + op
	}
	return s
}

func TestSumMulti(t *testing.T) {
	s := sumMulti(1, 2, 3, 4, 5)
	t.Log(s)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	// tsFun := timeSpent(slowFunc)
	// tsFun(5)
}
