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
}
