package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	i := 1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			res[top][column] = i
			i++
		}
		for row := top + 1; row <= bottom; row++ {
			res[row][right] = i
			i++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				res[bottom][column] = i
				i++
			}
			for row := bottom; row > top; row-- {
				res[row][left] = i
				i++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
