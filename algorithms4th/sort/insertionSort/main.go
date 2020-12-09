package main

import (
	"github.com/cntechpower/misc/algorithms4th/sort/testUtil"
)

func insertionSort(list []int, isAsc bool) {
	totalSwapCount := 0
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && ((list[j] < list[j-1] && isAsc) || (list[j] > list[j-1] && !isAsc)); i-- {
			tmp := list[i]
			list[i] = list[j]
			list[j] = tmp
			totalSwapCount++
		}
	}
	//fmt.Printf("total swap count: %v\n", totalSwapCount)
}

func main() {
	testUtil.Wrapper(insertionSort)
}
