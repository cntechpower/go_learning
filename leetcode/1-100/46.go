package main

import "fmt"

func dfs(nums []int, size, depth int, path []int, used []bool, res *[][]int) {
	if size == depth {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < size; i++ {
		if used[i] == false {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, size, depth+1, path, used, res)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, len(nums), 0, make([]int, 0, len(nums)), used, &res)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 1, 2}))
}
