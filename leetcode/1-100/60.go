package main

import (
	"fmt"
	"strings"
)

var used []bool
var gn int
var gk int
var sb strings.Builder

var factorial = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func dfs(index int) {
	if index == gn {
		return
	}
	for i := 1; i <= gn; i++ {
		if used[i] {
			continue
		}
		currentCount := factorial[gn-index-1]
		if currentCount < gk {
			gk = gk - currentCount
			continue
		}
		sb.WriteString(fmt.Sprintf("%v", i))
		used[i] = true
		dfs(index + 1)
	}
}

func getPermutation(n int, k int) string {
	sb = strings.Builder{}
	used = make([]bool, n+1)
	gn = n
	gk = k
	dfs(0)
	return sb.String()
}

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
}
