package main

import (
	"fmt"
)

var rows []int
var dale []int
var hills []int
var res [][]string
var count int

func isValidQueen(row, col, size int) bool {
	return (rows[col] + hills[row+col] + dale[row-col+size]) == 0
}

func placeQueen(board [][]rune, row, col, size int) {
	rows[col] = 1
	hills[row+col] = 1
	dale[row-col+size] = 1
	board[row][col] = 'Q'
}

func deleteQueen(board [][]rune, row, col, size int) {
	rows[col] = 0
	hills[row+col] = 0
	dale[row-col+size] = 0
	board[row][col] = '.'
}

func getEmptyBoard(size int) [][]rune {
	res := make([][]rune, 0, size)
	for i := 0; i < size; i++ {
		tmp := make([]rune, 0, size)
		for j := 0; j < size; j++ {
			tmp = append(tmp, '.')
		}
		res = append(res, tmp)
	}
	return res
}

func addSolution() {
	count++
}

func backtrack(board [][]rune, row, size int) {
	for col := 0; col < size; col++ {
		if isValidQueen(row, col, size) {
			placeQueen(board, row, col, size)
			if row == size-1 {
				addSolution()
			} else {
				backtrack(board, row+1, size)
			}
			deleteQueen(board, row, col, size)
		}
	}
}

func totalNQueens(n int) int {
	rows = make([]int, n)
	res = make([][]string, 0)
	hills = make([]int, 2*n)
	dale = make([]int, 2*n)
	board := getEmptyBoard(n)
	count = 0
	backtrack(board, 0, n)
	return count
}

func main() {
	fmt.Println(totalNQueens(1))
	fmt.Println(totalNQueens(4))
}
