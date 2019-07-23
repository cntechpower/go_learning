package main

import "fmt"

func main1() {
	s := "172.10.1.1"
	fmt.Println(s[0]) // 49-->1
	fmt.Println(int(s[0]))
	fmt.Println(int(s[0] - '0'))
}
