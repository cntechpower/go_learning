package main

import (
	"math/rand"
)

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	randIdx := rand.Intn(len(a))
	lower := make([]int, 0)
	bigger := make([]int, 0)
	for k, v := range a {
		if randIdx == k {
			continue
		}
		if v >= a[randIdx] {
			bigger = append(bigger, v)

		} else {
			lower = append(lower, v)
		}
	}
	lowerMid := append(quickSort(lower), a[randIdx])
	return append(lowerMid, quickSort(bigger)...)
}

func main() {
	in := make([]int, 0)
	for i := 1000000; i >= 1; i-- {
		in = append(in, i)
	}
	// fmt.Println(in)
	// fmt.Println(quickSort(in))
	in = quickSort(in)

}
