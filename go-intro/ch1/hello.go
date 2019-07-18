package main //包,表明代码所在的模块

import (
	"fmt" //引入代码依赖
	"os"
)

//功能实现
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		//os.Exit(2)
		fmt.Println("Hello World: ", os.Args[1])
	}
	fmt.Println("hello world")
	os.Exit(1)
}
