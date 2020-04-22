package main

import "fmt"

//https://leetcode-cn.com/problems/two-sum/solution/liang-shu-zhi-he-by-leetcode-2/

//执行用时：44 ms
//内存消耗：2.9 MB
func twoSum(nums []int, target int) []int {
	for start := 0; start <= (len(nums) - 1); start++ {
		for end := start + 1; end <= (len(nums) - 1); end++ {
			if nums[start]+nums[end] == target {
				return []int{start, end}
			}
		}
	}
	return []int{-1}
}

//执行用时：4 ms
//内存消耗：3.8 MB
func twoSum2(nums []int, target int) []int {
	valueIndexMap := make(map[int]int, 0)
	for idx, value := range nums {
		valueIndexMap[value] = idx
	}
	for startIdx, value := range nums {
		endIdx, ok := valueIndexMap[target-value]
		if ok && startIdx != endIdx {
			return []int{startIdx, endIdx}
		}
	}
	return []int{-1}
}

//执行用时：4 ms
//内存消耗：3.8 MB
func twoSum3(nums []int, target int) []int {
	valueIndexMap := make(map[int]int, 0)
	for startIdx, value := range nums {
		if endIdx, ok := valueIndexMap[target-value]; ok {
			return []int{startIdx, endIdx}
		}
		valueIndexMap[value] = startIdx
	}
	return []int{-1}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 8))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 10))

	fmt.Println(twoSum2([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(twoSum2([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(twoSum2([]int{1, 2, 3, 4, 5}, 8))
	fmt.Println(twoSum2([]int{1, 2, 3, 4, 5}, 10))

	fmt.Println(twoSum3([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(twoSum3([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(twoSum3([]int{1, 2, 3, 4, 5}, 8))
	fmt.Println(twoSum3([]int{1, 2, 3, 4, 5}, 10))
}
