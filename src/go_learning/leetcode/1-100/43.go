package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		iInt := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			jInt := int(num2[j] - '0')
			sum := tmp[i+j+1] + iInt*jInt
			tmp[i+j+1] = sum % 10
			tmp[i+j] = tmp[i+j] + sum/10
		}
	}
	stringB := strings.Builder{}
	for i, v := range tmp {
		if i == 0 && v == 0 {
			continue
		}
		stringB.WriteString(strconv.Itoa(v))
	}
	return stringB.String()
}

func main() {
	fmt.Println(multiply("123", "4"))
	fmt.Println(multiply("123", "456"))
}
