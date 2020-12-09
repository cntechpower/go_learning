package main

import (
	"fmt"
	"math"
)

func reverse_old(x int) int {
	lengthInt := 0
	for i := x; i != 0; i = i / 10 {
		lengthInt++
	}
	res := 0
	for i := x; i != 0; i = i / 10 {
		lengthInt--
		res = res + ((i % 10) * int(math.Pow10(lengthInt)))

	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

func reverse(x int) int {
	res := 0
	for i := x; i != 0; {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		res = (res * 10) + (i % 10)
		i = i / 10
	}
	return res
}

func main() {
	fmt.Println(reverse(112345))
	fmt.Println(reverse(120))
	fmt.Println(reverse(-2147483648))
	fmt.Println(reverse(math.MaxInt32))
}
