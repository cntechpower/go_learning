package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (s[0] == p[0] || rune(p[0]) == '.')
	if len(p) >= 2 && rune(p[1]) == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

func main() {
	fmt.Println(isMatch("a", "."))
	fmt.Println(isMatch("a", "a"))
	fmt.Println(isMatch("", ""))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
}
