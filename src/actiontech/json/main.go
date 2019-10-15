package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int
	Name string
}

func main() {
	p := &Person{
		Age:  24,
		Name: "DuJinYang",
	}
	pStr, _ := json.Marshal(p)
	fmt.Println(pStr)
	pNew := new(Person)
	if err := json.Unmarshal(pStr, pNew); err == nil {
		fmt.Println(pNew)
	}
}
