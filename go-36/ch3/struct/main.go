package main

import "fmt"

type pAge struct {
	sex string
	age int
}

type person struct {
	name string
	pAge
}

func (p pAge) String() string {
	return fmt.Sprintf("sex is %s,Age is %d\n", p.sex, p.age)
}
func (p person) printAge() string {
	return fmt.Sprintf("name is %s,sex is %s,Age is %d\n", p.name, p.pAge.sex, p.pAge.age)
}
func (p person) String() string {
	return fmt.Sprintf("name is %s,%s", p.name, p.pAge) //struct和String做了逐层嵌套
}
func main() {
	a := pAge{sex: "male", age: 22}
	fmt.Printf("%s\n", a)
	p := person{name: "dujinyang", pAge: a}
	fmt.Printf("%s\n", p.printAge())
	fmt.Printf("%s\n", p)

}
