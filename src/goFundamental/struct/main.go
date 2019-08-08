package main

import "fmt"

func main() {
	var p1, s1, h1 men
	p1 = &person{
		human: human{1},
		Sex:   0,
		Name:  "person",
	}
	fmt.Println(p1)
	s1 = &student{
		person: person{
			human: human{1},
			Sex:   0,
			Name:  "human",
		},
		sid: 1,
	}
	h1 = &human{Sex: 1}
	fmt.Println(s1)
	menSlice := make([]men, 3)
	menSlice[0], menSlice[1], menSlice[2] = p1, s1, h1
	for _, v := range menSlice {
		v.sayHi()
	}
	//s1.printName() //不能调用,因为接口中没有定义这个函数
	s2 := &student{
		person: person{
			human: human{1},
			Sex:   0,
			Name:  "human",
		},
		sid: 1,
	}
	s2.printName()
	fmt.Println("---------------------------------------------")
	fmt.Println("Comma-ok")
	for _, v := range menSlice {
		if val, ok := v.(*person); ok {
			fmt.Printf("%v : val is %v, type is person\n", v, val)
		} else if val, ok := v.(*student); ok {
			fmt.Printf("%v : val is %v, type is student\n", v, val)
		} else if val, ok := v.(*human); ok {
			fmt.Printf("%v : val is %v, type is human\n", v, val)
		}
	}
	fmt.Println("---------------------------------------------")
	fmt.Println("switch")
	for _, v := range menSlice {
		switch v.(type) {
		case *person:
			fmt.Printf("%v : type is person\n", v)
		case *student:
			fmt.Printf("%v : type is student\n", v)
		case *human:
			fmt.Printf("%v : type is human\n", v)

		}
	}

}
