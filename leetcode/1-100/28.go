package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" || len(haystack) < len(needle) {
		return -1
	}
	if haystack[:len(needle)] == needle {
		return 0
	}
	for start := 1; start < len(haystack)-len(needle)+1; start++ {
		if haystack[start:start+len(needle)] == needle {
			return start
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("aaaab", "b"))
	stringSlice("abc")
}

func stringSlice(s string) {
	//左闭右开
	fmt.Printf("s[0:2]=%v\n", s[0:2])
	fmt.Printf("s[0:3]=%v\n", s[0:3])
	fmt.Printf("s[1:2]=%v\n", s[1:2])
}
