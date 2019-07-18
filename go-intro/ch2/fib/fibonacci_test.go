package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	t.Log(a, "Test Output\n")
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)
