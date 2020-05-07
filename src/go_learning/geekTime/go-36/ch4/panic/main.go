package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter Main Function")
	defer func() {
		fmt.Println("Enter defer function")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
	}()
	panic(errors.New("Something wrong"))
	//fmt.Println("Fine, I'm OK") //unreachable code
}
