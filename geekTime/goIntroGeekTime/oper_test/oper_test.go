package oper_test

import (
	"strconv"
	"testing"
)

func TestCompare(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) //长度不同的数组无法比较
	t.Log(a == d)
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	sInt, _ := strconv.Atoi("10")
	t.Log(sInt)
}
