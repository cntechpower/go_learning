package main

import (
	"strconv"

	"fmt"
)

type s struct {
	a string
}

func main() {
	ss := make([]s, 0)
	for i := 1; i < 10; i++ {
		ss = append(ss, s{a: strconv.Itoa(i)})
	}
	for idx, s := range ss {
		if idx == 0 {
			s.a = "test"
			fmt.Println("Updated with variable!")
		}
		if idx == 1 {
			ss[idx].a = "test"
			fmt.Println("Updated with idx!")
		}

	}
	for idx, s := range ss {
		fmt.Printf("%v -- %v\n", idx, s.a)
	}
}
