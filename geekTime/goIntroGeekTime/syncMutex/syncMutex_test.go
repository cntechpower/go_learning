package syncMutex

import (
	"fmt"
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

//func TestFirstResponse(t *testing.T) {
//	t.Log("Before:", runtime.NumGoroutine())
//	//t.Log(FirstResponse())
//	t.Log(AllResponse())
//	time.Sleep(time.Second)
//	t.Log("Afer:", runtime.NumGoroutine()) //会有协程泄露
//}

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

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	fmt.Println(counter)
}

func TestCounterWithMutex(t *testing.T) {
	var muCounter sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			muCounter.Lock()
			defer func() { muCounter.Unlock() }()
			counter++
		}()
	}
	fmt.Println(counter)
}

func TestCounterWithMutexAndWaitGroup(t *testing.T) {
	var muCounter sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			muCounter.Lock()
			defer func() { muCounter.Unlock() }()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
