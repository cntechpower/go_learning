package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	leftValue := 0
	maxValue := nums[0]
	for i := 0; i < len(nums); i++ {
		leftValue = max(leftValue+nums[i], nums[i])
		maxValue = max(leftValue, maxValue)
	}
	return maxValue
}

func maxSubArrayWithOutput(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	leftValue := 0
	currLeft := 0
	currRight := 0
	left := 0
	right := 0
	maxValue := nums[0]
	for i := 0; i < len(nums); i++ {
		if leftValue+nums[i] > nums[i] {
			leftValue = leftValue + nums[i]
			currRight = i
		} else {
			leftValue = nums[i]
			currLeft = i
		}
		if leftValue > maxValue {
			maxValue = leftValue
			left = currLeft
			right = currRight
		}
	}
	if left == 0 {
		fmt.Printf("left: %v, right: %v, nums: %v\n", left, right, nums[:right])
	} else {
		fmt.Printf("left: %v, right: %v, nums: %v\n", left, right, nums[left:right+1])
	}
	return maxValue
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArrayWithOutput([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	t := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(t[3:6])
}
