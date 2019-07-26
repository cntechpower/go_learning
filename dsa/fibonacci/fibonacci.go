package main

import (
	"flag"
	"fmt"
)

var num int

func init() {
	flag.IntVar(&num, "num", 10, "Numbers Of Fibonacci")
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	fmt.Printf("Called finbonacci(%d) + fibonacci(%d)\n", num-1, num-2)
	return fibonacci(num-1) + fibonacci(num-2)
}

func fibonacciLinear(num int) []int {
	if num <= 1 {
		return []int{0, 1}
	}
	lastSlice := fibonacciLinear((num - 1))
	fmt.Printf("Called finbonacci(%d)\n", num-1)
	return []int{lastSlice[1], lastSlice[0] + lastSlice[1]}

}

func main() {
	flag.Parse()
	fmt.Println("*********************************************************")
	fmt.Println("finbonacci")
	fmt.Println(fibonacci(num)) //时间复杂度:O(n^2)
	fmt.Println("*********************************************************")
	fmt.Println("finbonacciLinear")
	fmt.Println(fibonacciLinear(num)[1]) //时间复杂度: O(n)
}
