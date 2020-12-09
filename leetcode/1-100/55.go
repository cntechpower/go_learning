package main

import "fmt"

var numsGlobal []int
var n int

func canJumpSlow(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	numsGlobal = nums
	n = len(numsGlobal)
	return backtrack(0)
}

func backtrack(currentPos int) bool {
	if currentPos == n-1 {
		// we arrive
		return true
	}
	if numsGlobal[currentPos] >= (n - 1 - currentPos) {
		//current value is bigger than left length
		//this is always can jump
		//fmt.Printf("quick early in pos: %v\n", currentPos)
		return true
	}
	if currentPos+numsGlobal[currentPos] >= n-1 {
		return true
	}
	if numsGlobal[currentPos] == 0 {
		// we not arrive, but we can not move any more
		return false
	}

	for step := numsGlobal[currentPos]; step >= 1; {
		if backtrack(currentPos + step) {
			return true
		}
		step--
	}
	return false
}

func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	rightMost := 0
	for k, v := range nums {
		if k <= rightMost {
			rightMost = max(rightMost, k+v)
			if rightMost >= n-1 {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 0, 0}))
	nums := make([]int, 0, 26000)
	for v := 25000; v >= 0; v-- {
		nums = append(nums, v)
	}
	fmt.Println(canJump(nums))
}
