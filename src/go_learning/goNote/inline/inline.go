package main

import "fmt"

func test() *int {
	a := 0x100
	return &a
}

func main() {
	var a *int = test()
	fmt.Println(a)

}
