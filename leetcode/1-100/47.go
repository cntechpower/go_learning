package main

import (
	"fmt"
	"sort"
)

func dfs(nums []int, size, depth int, path []int, used []bool, res *[][]int) {
	if size == depth {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		if used[i] == false {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, size, depth+1, path, used, res)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums, len(nums), 0, make([]int, 0, len(nums)), used, &res)
	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}
