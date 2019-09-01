package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i) //传值,输出内容不会冲突
		// go func() {
		// 	fmt.Println(i)
		// }() //共享变量,输出内容会出现冲突
	}
	time.Sleep(time.Millisecond * 50)

}
