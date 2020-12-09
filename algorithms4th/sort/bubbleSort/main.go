package main

import (
	"github.com/cntechpower/misc/algorithms4th/sort/testUtil"
)

func bubbleSort(a []int, isAsc bool) {
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a)-j-1; i++ { //最后j个元素已经有序, 无需再次比较
			if (isAsc && a[i] > a[i+1]) || (!isAsc && a[i] < a[i+1]) {
				a[i+1], a[i] = a[i], a[i+1]
			}
		}
	}
}

func main() {
	testUtil.Wrapper(bubbleSort)
}
