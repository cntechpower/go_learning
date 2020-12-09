package main

import (
	"github.com/cntechpower/misc/algorithms4th/sort/testUtil"
)

func partition(list []int, start, end int, isAsc bool) int {
	i := start
	pivot := list[end]
	for j := start; j < end; j++ {
		if (isAsc && list[j] < pivot) || (!isAsc && list[j] > pivot) {
			if i != j {
				list[i], list[j] = list[j], list[i]
			}
			i++
		}
	}
	list[i], list[end] = list[end], list[i]
	return i
}

func sort(list []int, start, end int, isAsc bool) {
	if end <= start {
		return
	}
	j := partition(list, start, end, isAsc)
	sort(list, start, j-1, isAsc)
	sort(list, j+1, end, isAsc)
}

func quickSort(list []int, isAsc bool) {
	sort(list, 0, len(list)-1, isAsc)
}

func main() {
	testUtil.Wrapper(quickSort)
}
