package main

// #include <stdlib.h>
import "C"
import "fmt"

func Random() int{
	return int(C.Random())
}

func Seed(i int){
	C.srandom(C.uint(i))
}
func main(){
	auto a=C.Random()
	fmt.Println(a)
}
