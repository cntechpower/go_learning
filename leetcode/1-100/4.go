package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n { //to ensure m <=n
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		temp := m
		m = n
		n = temp
	}
	//now len(nums2) >=len(nums1)
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] { // i is to small
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] { // i is to big
			iMax = i - 1
		} else {
			maxLeft := float64(0)
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				return maxLeft
			}
			minRight := float64(0)
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			//fmt.Printf("got i=%v, j=%v\n", i, j)
			//fmt.Printf("got maxLeft=%v ,minRight=%v\n", maxLeft, minRight)
			return (maxLeft + minRight) / 2
		}
	}
	return 0.0
}

func main() {
	a := []int{1, 3}
	b := []int{2}
	fmt.Println(findMedianSortedArrays(a, b))
	c := []int{1, 2}
	d := []int{1, 1}
	fmt.Println(findMedianSortedArrays(c, d))
}
