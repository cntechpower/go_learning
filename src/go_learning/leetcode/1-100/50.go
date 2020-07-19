package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}
func quickMul(x float64, n int) float64 {
	res := float64(1)
	current := x
	for n > 0 {
		if n%2 == 1 {
			res = res * current
		}
		current *= current
		n = n / 2
	}
	return res
}
func main() {
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}
