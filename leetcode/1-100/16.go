package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return res
			}
		}
	}
	return res
}
func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, 4}, 1))
}
