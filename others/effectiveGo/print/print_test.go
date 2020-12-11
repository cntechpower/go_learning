package main

import (
	"fmt"
	"testing"
)

type a struct {
	name string
	age  int
}

func TestPrint(t *testing.T) {
	a1 := a{
		name: "dujinyang",
		age:  23,
	}
	fmt.Printf("%v\n", a1)
	fmt.Printf("%+v\n", a1)
	fmt.Printf("%#v\n", a1)
}
