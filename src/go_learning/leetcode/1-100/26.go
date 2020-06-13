package main

import "fmt"

//160ms,4.6MB
func removeDuplicates2(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > len(nums) {
			break
		}
		for j := i + 1; j < length; j++ {
			if j >= len(nums) {
				break
			}
			for j < len(nums) && nums[i] == nums[j] {
				//fmt.Printf("removing nums[%v]: %v\n", i, nums[i])
				nums = append(nums[:i], nums[i+1:]...)
			}
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1}))
}
