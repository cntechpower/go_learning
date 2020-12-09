package main

import (
	"fmt"
	"regexp"
)

func main() {
	var mysqlVersion = regexp.MustCompile("^(\\d+\\.\\d+)\\.")

	fmt.Println(mysqlVersion.FindStringSubmatch("5.7.28"))
	fmt.Println(mysqlVersion.FindString("5.7.28"))

	var mysqlVersionWithoutDot = regexp.MustCompile("^(\\d+\\.\\d+)")

	//为啥不用下面这个..
	fmt.Println(mysqlVersionWithoutDot.FindString("5.7.28"))
}
