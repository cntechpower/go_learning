package main

import (
	"fmt"
)

type a struct {
	name string
	age  int
}

func main() {
	a1 := a{
		name: "dujinyang",
		age:  23,
	}
	fmt.Printf("%v\n", a1)
	fmt.Printf("%+v\n", a1)
	fmt.Printf("%#v\n", a1)
}
