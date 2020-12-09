package main

import "fmt"

func nextPermutationFalse(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			return
		}
	}
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func nextPermutation(nums []int) {
	reverse := func(start int) {
		left, right := start, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(i + 1)
}

func nextPermutation1(nums []int) []int {
	nextPermutation(nums)
	return nums

}

func main() {
	fmt.Println(nextPermutation1([]int{1, 2, 3}))
	fmt.Println(nextPermutation1([]int{1, 3, 2}))
	fmt.Println(nextPermutation1([]int{3, 2, 1}))
	fmt.Println(nextPermutation1([]int{1, 1, 5}))
}
