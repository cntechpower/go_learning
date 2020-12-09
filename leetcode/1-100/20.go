package main

import (
	"fmt"
)

func isValidOld(string) bool {
	left := make([]rune, 0)
	removeFromLeft := func() {
		if len(left) == 1 {
			left = make([]rune, 0)
		} else {
			left = left[:len(left)-1]
		}
	}
	isLastRuneMatch := func(r rune) bool {
		lastIdx := len(left) - 1
		if lastIdx < 0 {
			return false
		}
		return left[lastIdx] == r
	}
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			left = append(left, r)
		case ')':
			if isLastRuneMatch('(') {
				removeFromLeft()
			} else {
				return false
			}
		case '}':
			if isLastRuneMatch('{') {
				removeFromLeft()
			} else {
				return false
			}
		case ']':
			if isLastRuneMatch('[') {
				removeFromLeft()
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(left) != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	left := make([]rune, len(s))
	leftLength := 0
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			left[leftLength] = r
			leftLength++
		} else {
			if leftLength == 0 {
				return false
			}
			if (r == ']' && left[leftLength-1] == '[') ||
				(r == '}' && left[leftLength-1] == '{') ||
				(r == ')' && left[leftLength-1] == '(') {
				leftLength--
			} else {
				return false
			}
		}
	}

	return leftLength == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)])"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{[()]}"))
}
