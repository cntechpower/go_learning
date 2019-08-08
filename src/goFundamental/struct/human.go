package main

import "fmt"

type human struct {
	Sex int
}

func (h *human) setSex(i int) {
	h.Sex = i
}

func (h *human) sayHi() {
	fmt.Printf("Hi my Sex is %v\n", h.Sex)
}

func (h *human) sing(s string) {
	fmt.Printf("Human Singing %v\n", s)
}
