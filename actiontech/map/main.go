package main

import "fmt"

func main() {
	m := make(map[string]string, 0)
	m["a"] = "a"
	fmt.Println(m["b"])
	fmt.Println(m["b"] == "") //true
}
