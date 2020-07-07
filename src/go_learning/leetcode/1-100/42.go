package main

import (
	"fmt"
)

func trap1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[0] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	res := 0
	leftMax, rightMax := 0, 0
	for left <= right {
		if leftMax < rightMax {
			res += max(0, leftMax-height[left])
			leftMax = max(leftMax, height[left])
			left++
		} else {
			res += max(0, rightMax-height[right])
			rightMax = max(rightMax, height[right])
			right--
		}
	}
	return res
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
