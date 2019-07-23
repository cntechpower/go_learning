package main

import "fmt"

var name = "OOuter"

func main() {
	fmt.Printf("Name is : %v\n", name)
	name := "Outer"
	{
		name := "Inner"
		fmt.Printf("Name is : %v\n", name)
	}
	fmt.Printf("Name is : %v\n", name)
}
