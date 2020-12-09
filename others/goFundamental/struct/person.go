package main

import "fmt"

type person struct {
	human
	Sex  int
	Name string
}

func (p *person) setSex(i int) {
	p.Sex = i
}

func (p *person) sayHi() {
	fmt.Printf("Hi I'm %v,my Sex is %v\n", p.Name, p.Sex)
}

func (p *person) sing(s string) {
	fmt.Printf("Singing %v\n", s)
}
