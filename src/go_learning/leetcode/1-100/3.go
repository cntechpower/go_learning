package main

import "fmt"

//wrong
//fmt.Println(lengthOfLongestSubstring("dvdf")) should be 3, but this is 2.
func lengthOfLongestSubstringWrong(s string) int {
	byteMap := make(map[byte]struct{}, 0)
	maxLen := 0
	for _, b := range []byte(s) {
		if _, ok := byteMap[b]; !ok {
			byteMap[b] = struct{}{}
		} else {
			byteMap = map[byte]struct{}{b: {}}
		}
		if len(byteMap) > maxLen {
			maxLen = len(byteMap)
		}
	}
	return maxLen
}

//滑动窗口
func lengthOfLongestSubstringCommon(s string) int {
	byteMap := make(map[byte]struct{}, 0)
	maxLen := 0
	slice := []byte(s)
	left := 0
	right := 0
	for left < len(slice) && right < len(slice) {
		if _, ok := byteMap[slice[right]]; !ok {
			byteMap[slice[right]] = struct{}{}
			right++
			if right-left > maxLen {
				maxLen = right - left
			}
		} else {
			delete(byteMap, slice[left])
			left++
		}
	}
	return maxLen
}

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	byteMap := make(map[byte]struct{}, 0)
	maxLen := 0
	slice := []byte(s)
	left := 0
	right := 0
	for left < len(slice) && right < len(slice) {
		if _, ok := byteMap[slice[right]]; !ok {
			byteMap[slice[right]] = struct{}{}
			right++
			if right-left > maxLen {
				maxLen = right - left
			}
		} else {
			delete(byteMap, slice[left])
			left++
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("asjrgapa"))
}
