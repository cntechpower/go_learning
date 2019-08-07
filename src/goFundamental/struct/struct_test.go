package main

import (
	"fmt"
	"testing"
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

func TestStruct(t *testing.T) {
	a := student{
		human: human{1},
		Name:  "student1",
		Age:   18,
	}
	b := teacher{
		human: human{0},
		Name:  "teacher1",
		Age:   22,
	}
	fmt.Println(a)
	a.Age = 13
	a.human.Sex =
		fmt.Println(b)
}
