package main

import "fmt"

func main() {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers3 := numbers2
	maxIndex2 := len(numbers2) - 1
	fmt.Println(maxIndex2)
	for i, e := range numbers3 {
		fmt.Println(i, e)
		if i == maxIndex2 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)

}
