package main

import (
	"fmt"
	"go_learning/algorithms4th/sort/testUtil"
	"math/rand"
	"time"
)

func NewRandomList(length int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	}
	return res
}
func IsSorted(list []int, isAsc bool) bool {
	for i := 1; i < len(list); i++ {
		if (isAsc && list[i-1] > list[i]) || (!isAsc && list[i-1] < list[i]) {
			return false
		}
	}
	return true
}

func print(list []int) {
	fmt.Println("-------------------------------------")
	for _, p := range list {
		fmt.Println(p)
	}
	fmt.Println("-------------------------------------")
}

func partition(list []int, lo, hi int, isAsc bool) int {
	i := lo
	j := hi
	v := list[lo]
	for {
		for (isAsc && list[i] < v) || (!isAsc && list[i] > v) {
			i++
			if i == hi {
				break
			}
		}
		for (isAsc && list[j] >= v) || (!isAsc && list[j] <= v) {
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		fmt.Printf("will swap i=%v, j=%v\n", i, j)
		tmp := list[i]
		list[i] = list[j]
		list[j] = tmp
	}
	//TODO: Why need this?  P184
	//fmt.Printf("will swap i=%v, j=%v\n", lo, j)
	//tmp := list[lo]
	//list[lo] = list[j]
	//list[j] = tmp
	return j
}

func sort(list []int, lo, hi int, isAsc bool) {
	if hi <= lo {
		return
	}
	j := partition(list, lo, hi, isAsc)
	sort(list, lo, j-1, isAsc)
	sort(list, j+1, hi, isAsc)
}

func quickSort(list []int, isAsc bool) {
	sort(list, 0, len(list)-1, isAsc)
}

func main() {
	testUtil.Wrapper(quickSort)
}
