package main

import "fmt"

func Sort(list []int, isAsc bool) []int {
	totalSwapCount := 0
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && ((list[j] < list[j-1] && isAsc) || (list[j] > list[j-1] && !isAsc)); i-- {
			tmp := list[i]
			list[i] = list[j]
			list[j] = tmp
			totalSwapCount++
		}
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
