package main

import (
	"fmt"
	"time"
)

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() { fmt.Println("defer_closure i=", i) }()
		fs[i] = func() { fmt.Println("closure i=", i) }
	}
	for _, f := range fs {
		f()
	}
	time.Sleep(5 * time.Second)
}
