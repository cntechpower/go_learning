package main

import (
	"fmt"
)

//执行用时：992 ms
//内存消耗：2.3 MB
func longestPalindromeForce(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	byteSlice := []byte(s)
	isPalindrome := func(s string) bool {
		length := len(s)
		if s == "" || length == 1 {
			return true
		}
		for i := 0; i < length/2; i++ {
			if s[i] != s[length-i-1] {
				return false
			}
		}
		return true
	}
	longestNum := 0
	longestPalindrome := ""
	for i := 0; i <= len(byteSlice); i++ {
		for j := i; j <= len(byteSlice); j++ {
			subString := s[i:j]
			if i == j {
				subString = s[:i]
			}
			length := len(subString)
			if isPalindrome(subString) && length > longestNum {
				longestNum = length
				longestPalindrome = subString
			}
		}
	}
	return longestPalindrome
}

func longestPalindromeBuffer(s string) string {
	palindromeMap := make(map[string]struct{}, 0)
	if s == "" || len(s) == 1 {
		return s
	}
	byteSlice := []byte(s)
	isPalindrome := func(s string) bool {
		if s == "" || len(s) == 1 {
			return true
		}
		byteSliceInner := []byte(s)
		length := len(byteSliceInner)
		_, ok := palindromeMap[string(byteSliceInner[1:length-1])]
		//fmt.Printf("checking %v in map to ensure %v\n", string(byteSliceInner[1:length-1]), s)
		if byteSliceInner[0] == byteSliceInner[length-1] && ok {
			palindromeMap[s] = struct{}{}
			return true
		}
		for i := 0; i < length/2; i++ {
			if byteSliceInner[i] != byteSliceInner[length-i-1] {
				return false
			}
		}
		//fmt.Printf("add %v to map\n", s)
		palindromeMap[s] = struct{}{}
		return true
	}
	longestNum := 0
	longestPalindrome := ""
	for i := 0; i <= len(byteSlice); i++ {
		for j := i; j <= len(byteSlice); j++ {
			subString := s[i:j]
			if i == j {
				subString = s[:i]
			}
			length := len(subString)
			if isPalindrome(subString) && length > longestNum {
				longestNum = length
				longestPalindrome = subString
			}
		}
	}
	return longestPalindrome
}

func longestPalindromeExpand(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	runes := []rune(s)
	expand := func(left, right int) int {
		for left >= 0 && right < len(runes) && runes[left] == runes[right] {
			left--
			right++
		}
		return right - left - 1
	}
	start := 0
	end := 0
	for i := 0; i < len(runes); i++ {
		len1 := expand(i, i)
		len2 := expand(i, i+1)
		max := len1
		if len2 > len1 {
			max = len2
		}
		if max > end-start {
			start = i - ((max - 1) / 2)
			end = i + ((max) / 2)
		}

	}
	return string(runes[start : end+1])
}

func main() {
	//fmt.Println(longestPalindromeForce("abba"))
	//fmt.Println(longestPalindromeForce("babad"))
	//fmt.Println(longestPalindromeForce("a"))
	//fmt.Println(longestPalindromeForce("bb"))
	//fmt.Println(longestPalindromeForce("aaabaaaa"))
	//fmt.Println(longestPalindromeBuffer("aaabaaaa"))
	//fmt.Println(longestPalindromeForce("abacab"))
	//fmt.Println(longestPalindromeBuffer("abacab"))
	fmt.Println(longestPalindromeExpand("aaabaaaa"))
}
