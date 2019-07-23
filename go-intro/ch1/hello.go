package main //包,表明代码所在的模块

import (
	"fmt" //引入代码依赖
	"net"
	"os"
	"strconv"
	"strings"
)

//功能实现
func main() {
	if len(os.Args) == 2 {
		ip := net.ParseIP(os.Args[1])
		if ip == nil {
			fmt.Printf("Parse IP Error: %v\n", os.Args[1])
		} else {
			fmt.Printf("IP: %v", ip)
		}
		s := os.Args[1]
		fmt.Println(s)
		for k, v := range s {
			fmt.Printf("Key %v: %v\n", k, v)
		}
		for i := 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
			fmt.Println(int(s[i] - '0'))
			//fmt.Println(int(''))
		}
		fmt.Println(s[0])
		fmt.Println(int(s[0]))
		fmt.Println(int(s[0] - '0'))

		ipSplit := strings.Split(os.Args[1], ".")
		if len(ipSplit) != 4 {
			fmt.Printf("IP Format Wrong\n")
			os.Exit(1)
		}
		for _, v := range ipSplit {
			if ipValue, err := strconv.Atoi(v); err != nil || ipValue > 255 {
				fmt.Printf("%v is Not Number Or Bigger Than 255\n", v)
				os.Exit(1)
			}
		}
		ipHyphen := strings.Join(ipSplit, "-")
		fmt.Printf("OS Hostname is :%v\n", ipHyphen)
	} else {
		fmt.Printf("Args Count Must=1\n")
	}
	os.Exit(0)
}
