package main

import (
	"github.com/cntechpower/misc/algorithms4th/sort/testUtil"
)

func selectionSort(list []int, isAsc bool) {
	//totalSwapCount := 0
	for i := 0; i < len(list); i++ {
		candidateIdx := i
		for j := i + 1; j < len(list); j++ {
			if (list[j] < list[candidateIdx] && isAsc) || (list[j] > list[candidateIdx] && !isAsc) {
				candidateIdx = j
			}
		}
		//swap a[i] and a[candidateIdx]
		if i != candidateIdx {
			tmp := list[i]
			list[i] = list[candidateIdx]
			list[candidateIdx] = tmp
			//totalSwapCount++

		}
	}
	//fmt.Printf("total swap count: %v\n", totalSwapCount)
}

func main() {
	testUtil.Wrapper(selectionSort)
}
