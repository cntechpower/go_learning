// struct
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type student struct {
	person
	sid int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}
func structtest() {
	var p1 person
	var p2 person
	p1.name = "dujinyang"
	p1.age = 18
	p2.name = "zhaoyijing"
	p2.age = 16
	b, b_age := Older(p1, p2)
	fmt.Printf("%v Is %v Older\n", b.name, b_age)
	var p3 student
	p3.age = 18
	p3.name = "xuesheng"
	p3.sid = 22
	fmt.Printf("name is %v,age is %v,sid is %v", p3.name, p3.age, p3.sid)
}
