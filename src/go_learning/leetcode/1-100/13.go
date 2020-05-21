package main

import (
	"fmt"
)

func romanToInt(s string) int {
	res := 0
	intValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanValues := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(intValues); i++ {
		//for strings.HasPrefix(s, romanValues[i]) {
		for len(s) >= len(romanValues[i]) && s[0:len(romanValues[i])] == romanValues[i] {
			res = res + intValues[i]
			s = s[len(romanValues[i]):]
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("MCMXCIV"))
}
