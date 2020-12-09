package main

import (
	"fmt"
	"sort"
)

func combinationSumWrong(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return nil
	}
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
		if candidates[i] <= target {
			innerCombinationSum(res, append(resI, candidates[i]), i, size, candidates, target-candidates[i])
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	ret := [][]int{}
	for i, d := range candidates {
		if target-d < 0 {
			break
		} else if target-d == 0 {
			ret = append(ret, []int{d})
			continue
		}
		for _, v := range dfs(candidates[i:], target-d) {
			ret = append(ret, append([]int{d}, v...))
		}
	}
	return ret
}

func main() {
	fmt.Println(combinationSumWrong([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSumWrong([]int{2, 3, 5}, 8))
	fmt.Println(combinationSumWrong([]int{7, 3, 2}, 18))
}
