package main

import "fmt"

func Sort(list []int, isAsc bool) []int {
	totalSwapCount := 0
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
