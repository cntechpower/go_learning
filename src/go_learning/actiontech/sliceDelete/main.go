package main

import "fmt"

type a struct {
	name string
}

func del1(as []a, name string) []a {
	for idx := len(as) - 1; idx >= 0; idx-- {
		if as[idx].name == name {
			as = append(as[:idx], as[idx+1:]...)
		}
	}
	return as
}
func del2(as []a, name string) []a {
	for idx, value := range as {
		if value.name == name {
			as = append(as[:idx], as[idx+1:]...)
		}
	}
	return as
}

func main() {
	as := []a{{name: "t1"}, {name: "t2"}, {name: "t3"}, {name: "t2"}, {name: "t4"}, {name: "t5"}}
	fmt.Println(as)
	fmt.Println(del1(as, "t2"))
	fmt.Println(del2(as, "t2"))
}
