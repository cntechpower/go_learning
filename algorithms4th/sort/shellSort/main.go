package main

import (
	"github.com/cntechpower/misc/algorithms4th/sort/testUtil"
)

func shellSort(list []int, isAsc bool) {
	//totalSwapCount := 0
	n := len(list)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && ((list[j] < list[j-h] && isAsc) || (list[j] > list[j-h] && !isAsc)); j -= h {
				tmp := list[j-h]
				list[j-h] = list[j]
				list[j] = tmp
				//totalSwapCount++
			}
		}
		h = h / 3
	}
	//fmt.Printf("total swap count: %v\n", totalSwapCount)
}

func main() {
	testUtil.Wrapper(shellSort)
}
