package main

import "fmt"

func main() {
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781
	fmt.Println(a, b, c)
	fmt.Println(a == b, a == c) //always true, because value is truncated
	fmt.Println(a, b, c)
}
