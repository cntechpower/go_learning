package main

import "fmt"

func uniquePaths(m int, n int) int {
	cur := make([]int, n)
	for idx := range cur {
		cur[idx] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = cur[j-1] + cur[j]
		}
	}
	return cur[n-1]
}

func main() {
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 2))
}
