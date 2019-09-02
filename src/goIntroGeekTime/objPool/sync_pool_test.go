package objpool

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC()
	v1 := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPool2(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 88
		},
	}
	pool.Put(99)
	pool.Put(99)
	pool.Put(99)
	var wg sync.WaitGroup
	//var mu sync.Mutex
	//var execCount = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			// mu.Lock()
			// execCount++
			// mu.Unlock()
			v := pool.Get().(int)
			fmt.Printf("Values is %d\n", v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
