package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	np := len(p)
	if np == 0 {
		return false
	}
	lastPIsAny := false
	for i := 0; i < len(s); i++ {
		if i >= np {
			if lastPIsAny {
				return true
			}
			return false
		}
		if s[i] == p[i] || p[i] == '?' {
			lastPIsAny = false
			continue
		}
		if p[i] == '*' {
			lastPIsAny = true
			continue
		}
		if lastPIsAny {
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("abedeb", "*a*b"))
	fmt.Println(isMatch("ac", "a*c"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
}
