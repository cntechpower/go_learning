package main

import "fmt"

func isStringMatch(s string, words []string) bool {
	//bitmap
	bitmap := make([]bool, len(words))
	matchCount := 0
	for len(words) != matchCount {
		isMatch := false
		for i := 0; i < len(words); i++ {
			if bitmap[i] == true {
				continue
			}
			wLen := len(words[i])
			if len(s) < wLen {
				return false
			}
			//fmt.Printf("compare %v to %v\n", s[:wLen], words[i])
			if s[:wLen] == words[i] {
				s = s[wLen:]
				isMatch = true
				bitmap[i] = true
				matchCount++
			}
		}
		if !isMatch {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) []int {
	wordLen := 0
	res := make([]int, 0)
	for _, w := range words {
		wordLen += len(w)
	}
	if len(s) == 0 || wordLen == 0 {
		return res
	}
	if len(s) < wordLen {
		return res
	}
	for i := 0; i < len(s)-wordLen+1; i++ {
		//fmt.Printf("check isStringMatch for %v\n", s[i:])
		if isStringMatch(s[i:], words) {
			res = append(res, i)
		}
		//fmt.Printf("current words is %v\n", words)
	}
	return res
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
}
