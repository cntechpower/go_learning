package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	bytesSlice := []byte(s)
	tmp := make([]strings.Builder, numRows)
	currRow := 0
	goingDown := false
	for _, c := range bytesSlice {
		tmp[currRow].WriteByte(c)
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			currRow = currRow + 1
		} else {
			currRow = currRow - 1
		}
	}
	res := strings.Builder{}
	for _, t := range tmp {
		res.WriteString(t.String())
	}

	return res.String()
}
func main() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
	fmt.Println(convert("LEETCODEISHIRING", 4))
}
