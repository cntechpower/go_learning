package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstacne *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Called Create Obj")
		singleInstacne = new(Singleton)
	})
	return singleInstacne
}

func TestOnce(t *testing.T) { //只运行一次
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}(&wg)
	}
	wg.Add(10)
	wg.Wait()
}
