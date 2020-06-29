package main

import (
	"strconv"
)

func intToBytes(n int) byte {
	return byte(n)

}
func boxIndex(row, col int) int {
	return (row/3)*3 + col/3
}
func isNumInSlice(num int, slice []int) bool {
	return slice[num] == 0
}
func checkSlice(rows, cols, boxs [][]int, row, col int) {
	if rows[row] == nil {
		rows[row] = make([]int, 10)
	}
	if cols[col] == nil {
		cols[col] = make([]int, 10)
	}
	boxIdx := boxIndex(row, col)
	if boxs[boxIdx] == nil {
		boxs[boxIdx] = make([]int, 10)
	}
}
func couldPlace(rows, cols, boxs [][]int, num, row, col int) bool {

	return !isNumInSlice(num, cols[col]) &&
		!isNumInSlice(num, rows[row]) &&
		!isNumInSlice(num, boxs[boxIndex(row, col)])
}
func placeNum(board [][]byte, rows, cols, boxs [][]int, num, row, col int) {
	checkSlice(rows, cols, boxs, row, col)
	rows[row][num] += 1
	cols[col][num] += 1
	boxs[boxIndex(row, col)][num] += 1
	board[row][col] = intToBytes(num)
}
func removeNum(board [][]byte, rows, cols, boxs [][]int, num, row, col int) {
	checkSlice(rows, cols, boxs, row, col)
	rows[row][num] = 0
	cols[col][num] = 0
	boxs[boxIndex(row, col)][num] = 0
	board[row][col] = '.'
}

func placeNextNum(board [][]byte, rows, cols, boxs [][]int, row, col int, solved *bool) {
	if row == len(board)-1 && col == len(board) {
		*solved = true
	} else {
		if col == len(board)-1 {
			backTrack(board, rows, cols, boxs, row+1, 0, solved)
		} else {
			backTrack(board, rows, cols, boxs, row, col+1, solved)
		}
	}
}

func solveSudoku(board [][]byte) {
	if board == nil || len(board) != 9 {
		return
	}
	rows := make([][]int, 10)
	cols := make([][]int, 10)
	boxs := make([][]int, 10)
	solved := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				numInt, _ := strconv.Atoi(string(board[i][j]))
				placeNum(board, rows, cols, boxs, numInt, i, j)
			}
		}
	}

	backTrack(board, rows, cols, boxs, 0, 0, &solved)

}

func backTrack(board [][]byte, rows, cols, boxs [][]int, row, col int, solved *bool) {
	if board[row][col] == '.' {
		for i := 1; i < 9; i++ {
			if couldPlace(rows, cols, boxs, i, row, col) {
				placeNum(board, rows, cols, boxs, i, row, col)
				placeNextNum(board, rows, cols, boxs, row, col, solved)
				if !*solved {
					removeNum(board, rows, cols, boxs, i, row, col)
				}
			}
		}
	} else {
		placeNextNum(board, rows, cols, boxs, row, col, solved)
	}
}
