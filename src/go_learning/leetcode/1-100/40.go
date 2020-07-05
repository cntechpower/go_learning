package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return nil
	}
	sort.Ints(candidates)
	res := make([][]int, 0)
	innerCombinationSum(&res, make([]int, 0), 0, n, candidates, target)
	return res
}

func innerCombinationSum(res *[][]int, resI []int, left, size int, candidates []int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, resI...))
		return
	}
	for i := left; i < size; i++ {
		if i > left && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] <= target {
			innerCombinationSum(res, append(resI, candidates[i]), i+1, size, candidates, target-candidates[i])
		} else {
			break
		}
	}
}
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
