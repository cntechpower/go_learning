package main

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rows := make([][]int, 9)
	cols := make([][]int, 9)
	boxs := make([][]int, 9)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			num := board[i][j]
			if num != '.' {
				numInt, _ := strconv.Atoi(string(num))
				boxIdx := (i/3)*3 + j/3
				if rows[i] == nil {
					rows[i] = make([]int, 10)
				}
				rows[i][numInt] = rows[i][numInt] + 1
				if cols[j] == nil {
					cols[j] = make([]int, 10)
				}
				cols[j][numInt] = cols[j][numInt] + 1
				if boxs[boxIdx] == nil {
					boxs[boxIdx] = make([]int, 10)
				}
				boxs[boxIdx][numInt] = boxs[boxIdx][numInt] + 1
				if rows[i][numInt] > 1 || cols[j][numInt] > 1 || boxs[boxIdx][numInt] > 1 {
					return false
				}
			}
		}
	}
	return true
}
