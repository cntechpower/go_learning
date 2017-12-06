package main

import (
	"fmt"
	"unsafe"
)

const (
	j = "abc"
	k = len(j)
	l = unsafe.Sizeof(j)
)

func staticvari() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
}
func staticvar() {
	println(j, k, l)
}
