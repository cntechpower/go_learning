package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2 // avoid overflow for (l+r)/2
		if mid >= n {
			break
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
