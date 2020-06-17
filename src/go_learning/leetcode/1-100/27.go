package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	l := []int{3, 2, 2, 3}
	fmt.Println(removeElement(l, 3))
	fmt.Println(l)
	l1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(l1, 2))
	fmt.Println(l1)
}
