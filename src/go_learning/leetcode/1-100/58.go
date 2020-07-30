package main

import "fmt"

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return end - start
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World "))
}
