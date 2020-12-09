package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || (n == 1 && nums[0] != target) {
		return []int{-1, -1}
	}
	if n == 1 && nums[0] == target {
		return []int{0, 0}
	}
	l, r := 0, n-1
	leftMatch, rightMath := -1, -1
	for l <= r {
		mid := l + (r-l)/2 // avoid overflow for (l+r)/2
		if mid >= n {
			break
		}
		//fmt.Printf("left: %v, right: %v, mid: %v\n", l, r, mid)
		if nums[mid] == target {
			leftMatch = mid
			rightMath = mid
			for i := mid; i >= 1; i-- {
				if nums[i] == nums[i-1] {
					leftMatch--
				} else {
					break
				}
			}
			for i := mid; i < n-1; i++ {
				if nums[i] == nums[i+1] {
					rightMath++
				} else {
					break
				}
			}
			break
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{leftMatch, rightMath}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{1, 3}, 1))
	fmt.Println(searchRange([]int{2, 2}, 3))
	fmt.Println(searchRange([]int{3, 3, 3}, 3))
}
