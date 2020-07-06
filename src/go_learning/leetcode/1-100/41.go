package main

import (
	"fmt"
	"math"
	"sort"
)

func firstMissingPositiveUsingMap(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}
	for i := 1; i < math.MaxInt32; i++ {
		if _, ok := numsMap[i]; !ok {
			return i
		}
	}
	return 0
}

func firstMissingPositiveUsingSort(nums []int) int {
	sort.Ints(nums)
	res := 1
	for _, num := range nums {
		if num == res {
			res++
		}
	}
	return res
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{2, 1}))
}
