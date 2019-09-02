package wait_complate_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var muId sync.Mutex

func runTask(t int) string {
	time.Sleep(time.Millisecond)
	return fmt.Sprintf("No %d Call Returned", t)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret

		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	//t.Log(FirstResponse())
	t.Log(AllResponse())
	time.Sleep(time.Second)
	t.Log("Afer:", runtime.NumGoroutine()) //会有协程泄露
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret

		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet = finalRet + <-ch + "\n"
	}
	return finalRet
}
