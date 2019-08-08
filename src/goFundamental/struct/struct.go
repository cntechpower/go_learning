package main

import "fmt"

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

type str string

func (s str) printStd() {
	s = str("new string")
	fmt.Printf("This is printStd: %v\n", s)
}

func main() {
	//a := student{
	//	human: human{1},
	//	Name:  "student1",
	//	Age:   18,
	//}
	//b := teacher{
	//	human: human{0},
	//	Name:  "teacher1",
	//	Age:   22,
	//}
	//fmt.Println(a)
	//a.Age = 13
	//a.human.Sex = 0
	//fmt.Println(b)
	p := person{
		human: human{1},
		Sex:   2,
		Name:  "person",
	}
	p.Sex = 3
	p.human.Sex = 4
	fmt.Println(p)
	p.setSex(5)
	fmt.Println(p)
	p.human.setSex(6)
	fmt.Println(p)
	s1 := str("testStr")
	s1.printStd()
}
