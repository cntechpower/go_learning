package main

import (
	"fmt"
)

var list = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func groupAnagrams(strs []string) [][]string {
	res := make(map[int][]string, 0)
	if len(strs) == 0 {
		return nil
	}
	stringHash := func(a string) int {
		ret := 1
		for i := 0; i < len(a); i++ {
			ret *= list[int(a[i]-'a')]
		}
		return ret
	}
	for i := 0; i < len(strs); i++ {
		hashValue := stringHash(strs[i])
		res[hashValue] = append(res[hashValue], strs[i])
	}
	slices := make([][]string, 0, len(res))
	for _, s := range res {
		slices = append(slices, s)

	}
	return slices
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
