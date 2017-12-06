package main

import "fmt"

func Fackorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Fackorial(x-1)
	}
	return
}
func main() {
	var i int = 15
	fmt.Printf("%d的阶乘是%d\n", i, Fackorial(i))
}
