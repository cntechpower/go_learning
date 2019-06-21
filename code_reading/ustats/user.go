package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {
	//https://golang.org/pkg/os/user/
	// 将字符串转换为十进制整数，即：ParseInt(s, 10, 0) 的简写）
	//func Atoi(s string) (int, error)
	username, _ := user.Current()
	fmt.Println(username.Name)
	if rootUser, err := user.Lookup("root"); err == nil {
		uid, _ := strconv.Atoi(rootUser.Uid)
		fmt.Printf("root uid is %v", uid)
	} else {
		fmt.Printf("Lookup User Root fail, reason is %v\n", err)
	}
	uid := os.Getuid()
	fmt.Println(uid)
	hostname, _ := os.Hostname()
	fmt.Println(hostname)
	fmt.Printf(os.Args[0], os.Args[1:])

}
