package main

import "fmt"

func Sort(list []int, isAsc bool) []int {
	totalSwapCount := 0
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
				totalSwapCount++
			}
		}
		h = h / 3
	}
	fmt.Printf("total swap count: %v\n", totalSwapCount)
	return list
}

func main() {
	list := []int{1, 3, 4, 5, 2, 6, 7, 3, 9}
	listSorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Sort(list, true))
	fmt.Println(Sort(listSorted, true))
	fmt.Println(Sort(list, false))
	fmt.Println(Sort(listSorted, false))
}
