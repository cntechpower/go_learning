package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
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
	if target == 0 {
		*res = append(*res, resI)
		fmt.Printf("append %v to res\n", resI)
		return
	}
	for i := left; i < size; i++ {
		newTarget := target - candidates[i]
		if newTarget < 0 {
			continue
		}
		resI = append(resI, candidates[i])
		fmt.Printf("appending %v to resI\n", candidates[i])
		innerCombinationSum(res, append(resI, candidates[i]), i, size, candidates, newTarget)
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
