package main

import (
	"fmt"
	"math"
)

func isPalindrome1(x int) bool {
	// 如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	lenInt := 0
	for i := x; i != 0; i = i / 10 {
		lenInt++
	}
	reverse := 0
	for i := x; i != 0; {
		reverse = reverse + (i%10)*int(math.Pow10(lenInt-1))
		lenInt--
		i = i / 10

	}
	fmt.Printf("check %v, reverse is %v\n", x, reverse)
	return x == reverse
}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverse := 0
	for i := x; i != 0; {
		reverse = reverse*10 + i%10
		i = i / 10
	}
	fmt.Printf("check %v, reverse is %v\n", x, reverse)
	return x == reverse
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x = x / 10
	}
	fmt.Printf("check %v, reverse is %v\n", x, reverse)
	return x == reverse || x == reverse/10
}
func main() {
	fmt.Println(isPalindrome(12345))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(0))
}
