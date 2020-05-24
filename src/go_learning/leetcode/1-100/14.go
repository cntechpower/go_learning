package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := 0
	for _, str := range strs {
		if len(str) > minLen {
			minLen = len(str)
		}
	}
	prefix := strings.Builder{}
	for i := 0; i < minLen; i++ {
		var c uint8
		for _, str := range strs {
			if len(str) <= i {
				return prefix.String()
			}
			if c == 0 {
				c = str[i]
			}
			if str[i] != c {
				return prefix.String()
			}
		}
		prefix.WriteRune(rune(strs[0][i]))
	}
	return prefix.String()
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	fmt.Println(longestCommonPrefix([]string{"", "b"}))
}
