package main

type person struct {
	human
	Sex  int
	Name string
}

func (p *person) setSex(i int) {
	p.Sex = i
}
