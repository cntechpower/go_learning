package main

import "fmt"

type student struct {
	person
	sid int
}

func (s *student) sayHi() {
	fmt.Printf("Hi I'm student\n")
}

func (s *student) printName() {
	fmt.Printf("student only\n")
}
