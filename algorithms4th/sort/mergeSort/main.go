package main

import (
	"fmt"
	"github.com/cntechpower/misc/algorithms4th/sort/testUtil"
	"math"
)

var tmp []int

func merge(list []int, lo, mid, hi int, isAsc bool) {
	i := lo
	j := mid + 1
	//tmp := make([]int, 0, hi-lo)
	for k := lo; k <= hi; k++ {
		tmp[k] = list[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid {
			// [lo:mid] is empty, just append [mid:hi] to result
			list[k] = tmp[j]
			j++
		} else if j > hi {
			// [mid:hi] is empty, just append [lo:mid] to result
			list[k] = tmp[i]
			i++

		} else if (isAsc && tmp[i] < tmp[j]) || (!isAsc && tmp[i] > tmp[j]) {
			// both not empty, append bigger object to result
			list[k] = tmp[i]
			i++
		} else {
			list[k] = tmp[j]
			j++
		}
	}
}

func sort(list []int, lo, hi int, isAsc bool) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	//fmt.Printf("sorting lo=%v,mid=%v,hi=%v\n", lo, mid, hi)
	sort(list, lo, mid, isAsc)
	sort(list, mid+1, hi, isAsc)
	merge(list, lo, mid, hi, isAsc)
}
func mergeSort(list []int, isAsc bool) {
	tmp = make([]int, len(list))
	sort(list, 0, len(list)-1, isAsc)
}

func mergeSortBottomUp(list []int, isAsc bool) {
	n := len(list)
	tmp = make([]int, len(list))
	for size := 1; size < n; size = size * 2 {
		for lo := 0; lo < n-size; lo += size * 2 {
			fmt.Printf("merging lo=%v,mid=%v,hi=%v\n", lo, lo+size-1, int(math.Min(float64(lo+size+size-1), float64(n-1))))
			merge(list, lo, lo+size-1, int(math.Min(float64(lo+size+size-1), float64(n-1))), isAsc)
		}
	}

}

func main() {
	testUtil.Wrapper(mergeSort)
}
