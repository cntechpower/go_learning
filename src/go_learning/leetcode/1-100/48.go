package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	fmt.Println(" before rotate: ")
	printMatrix(matrix)
	reverse := func(s []int) {
		n := len(s)
		for i := 0; i < n/2; i++ {
			s[i], s[n-i-1] = s[n-i-1], s[i]
		}
	}
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	fmt.Println("after transpose: ")
	printMatrix(matrix)
	for i := 0; i < n; i++ {
		reverse(matrix[i])
	}
	fmt.Println("after reverse: ")
	printMatrix(matrix)
}

func printMatrix(res [][]int) {
	fmt.Println("-----------------------------")
	for _, s := range res {
		fmt.Println(s)
	}
	fmt.Println("-----------------------------")
}

func main() {
	//res1 := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	//res2 := [][]int{
	//	{5, 1, 9, 11},
	//	{2, 4, 8, 10},
	//	{13, 3, 6, 7},
	//	{15, 14, 12, 16},
	//}

	res3 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	//rotate(res1)
	//rotate(res2)
	rotate(res3)
}
