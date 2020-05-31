package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		min := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if min > target {
			break
		}
		max := nums[i] + nums[length-1] + nums[length-2] + nums[length-3]
		if max < target {
			continue
		}
		for k := i + 1; k < length-2; k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}
			left := k + 1
			right := length - 1
			min := nums[i] + nums[k] + nums[left] + nums[left+1]
			if min > target {
				continue
			}
			max := nums[i] + nums[k] + nums[right] + nums[right-1]
			if max < target {
				continue
			}
			for left < right {
				curr := nums[i] + nums[k] + nums[left] + nums[right]
				if curr == target {
					res = append(res, []int{nums[i], nums[k], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if curr > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
