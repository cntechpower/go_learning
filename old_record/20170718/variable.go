package main

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

//g,h:=123,"hello"这种声明方式只能出现在函数体内

func variable() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
