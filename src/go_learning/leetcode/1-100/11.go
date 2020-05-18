package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left := 0
	areaMax := 0
	right := len(height) - 1
	minInt := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	for left < right {
		area := minInt(height[left], height[right]) * (right - left)
		if area > areaMax {
			areaMax = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return areaMax
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
